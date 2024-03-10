package main

import (
	qyblog "github.com/iwinder/qyblog/internal/qycms-system"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	qyblog.NewApp("iam-apiserver")
}
