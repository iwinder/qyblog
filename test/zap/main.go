package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://exampl.windcoder.com/api"
	logger.Warn("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	sugar := logger.Sugar()
	sugar.Infow("sugar fail to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second)

	sugar.Infof("sugar INNFO Failed to fetch URL %s", url)
}
