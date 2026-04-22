package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/aqyuki/sandbox/postgres-replication/internal/config"
	"github.com/aqyuki/sandbox/postgres-replication/internal/db"
	"github.com/aqyuki/sandbox/postgres-replication/internal/logging"
	"github.com/aqyuki/sandbox/postgres-replication/internal/lox"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/sync/errgroup"
)

var log = logging.Log

const (
	ModeInsert = "insert"
	ModeClean  = "clean"
	ModeDemo   = "demo"
)

const recordCount = 1_000_000
const concurrentAccessCount = 1_000

var (
	flagMode  = flag.String("mode", "", "動作モードを指定 (insert / clean / demo)")
	flagDebug = flag.Bool("debug", false, "Debug モード")
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID        uuid.UUID  `bun:"id,type:uuid"`
	Username  string     `bun:"username"`
	CreatedAt *time.Time `bun:"created_at,type:timestamptz"`
}

func init() { flag.Parse() }

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		done()
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.LoadEnv()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	db, err := db.CreateBunConn(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database connection: %w", err)
	}
	defer db.Close()

	mode := strings.ToLower(lox.FromPtr(flagMode))

	if lox.FromPtr(flagDebug) {
		fmt.Printf("[Mode] : %s\n", mode)
	}

	switch mode {
	case ModeInsert:
		if err := insert(ctx, db); err != nil {
			return fmt.Errorf("failed to exec insert command: %w", err)
		}
		return nil
	case ModeClean:
		if err := clean(ctx, db); err != nil {
			return fmt.Errorf("failed to exec clean command: %w", err)
		}
		return nil
	case ModeDemo:
		if err := demo(ctx, db); err != nil {
			return fmt.Errorf("failed to exec demo command: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("invalid mode specified: %s", mode)
	}
}

func insert(ctx context.Context, db *bun.DB) error {

	accounts := make([]Account, 0, recordCount)

	var b strings.Builder

	log("data creating.")
	for i := range recordCount {
		// エラー返されるけど、キャッチしたところでどうしょうもないので握りつぶす
		_, _ = b.WriteString("user")
		_, _ = b.WriteString(strconv.Itoa(i + 1))

		accounts = append(accounts, Account{
			ID:       uuid.New(),
			Username: b.String(),
		})

		// Builder をリセット
		b.Reset()
	}
	log("data created.")

	log("data inserting.")
	if _, err := db.NewInsert().Model(&accounts).Exec(ctx); err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}
	log("data inserted.")

	// trigger GC
	// 確保してるデータ量が大きいので後続処理に影響が及ばないようにGCをトリガー
	log("GC started.")
	accounts = nil
	runtime.GC()
	log("GC completed")

	return nil
}

func clean(ctx context.Context, db *bun.DB) error {
	log("data truncating.")
	if _, err := db.NewTruncateTable().Model((*Account)(nil)).Exec(ctx); err != nil {
		return fmt.Errorf("failed to truncate: %w", err)
	}
	log("data truncated.")
	return nil
}

func demo(ctx context.Context, db *bun.DB) error {
	var g errgroup.Group

	for i := range concurrentAccessCount {
		g.Go(func() error {
			var b strings.Builder
			b.WriteString("user")
			b.WriteString(strconv.Itoa(i + 1))

			var account Account
			if err := db.NewSelect().Model(&account).Where("username = ?", b.String()).Scan(ctx); err != nil {
				return fmt.Errorf("failed to scan data: %w", err)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("returned some errors: %w", err)
	}
	return nil
}
