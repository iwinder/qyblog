package main

import (
	qycmssystem "gitee.com/windcoder/qingyucms/internal/qycms-system"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	qycmssystem.NewApp("qycms-system").Run()
}
