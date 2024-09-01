package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.GetString("channel")
			if token == "" {
				fmt.Println("Please set channel ID")
				return
			}

			q := url.Values{}
			q.Set("channel_id", token)

			u := url.URL{
				Scheme:   "https",
				Host:     "www.youtube.com",
				Path:     "/feeds/videos.xml",
				RawQuery: q.Encode(),
			}
			fmt.Printf("Polling %s\n", u.String())

			timer := time.NewTicker(30 * time.Second)
			defer timer.Stop()

			ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
			defer done()

			var lastFetch time.Time

			fmt.Printf("Start polling\nCtrl+C to stop\n")
			feed, err := fetchFeed(u)
			if err != nil {
				panic(err)
			}

			lastFetch = time.Now()
			for _, item := range feed.Items {
				printFeed(item)
			}

			for {
				select {
				case <-timer.C:
					feed, err := fetchFeed(u)
					if err != nil {
						panic(err)
					}

					var count int
					for _, item := range feed.Items {
						if item.PublishedParsed.After(lastFetch) {
							printFeed(item)
							count++
						}
					}
					if count == 0 {
						fmt.Println("No new videos")
					}
					lastFetch = time.Now()
				case <-ctx.Done():
					return
				}
			}
		},
	}
)

func fetchFeed(u url.URL) (*gofeed.Feed, error) {
	feed, err := gofeed.NewParser().ParseURL(u.String())
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func printFeed(item *gofeed.Item) {
	fmt.Printf(`-----
Title : %s
Link : %s
Video ID : %s
Channel : %s
-----`, item.Title, item.Link, item.GUID, item.Author.Name)
	fmt.Println()
}

func init() {
	rootCmd.PersistentFlags().String("channel", "", "チャンネルID")

	if err := viper.BindPFlag("channel", rootCmd.PersistentFlags().Lookup("channel")); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
