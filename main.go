package main

import (
	"errors"
	"fmt"
)

func main() {
	//logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any
	//sugar := logger.Sugar()
	//url := "www.baidu.com"

	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
	fmt.Println("test01")
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
