package logging

import (
	"fmt"
	"time"
)

func Log(msg string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339Nano), msg)
}
