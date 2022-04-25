package main

import (
	"gitee.com/windcoder/qingyucms/internal/qycms"
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
	qycms.NewApp("qycms-system").Run()
}
