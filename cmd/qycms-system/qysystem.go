package main

import (
	"gitee.com/windcoder/qingyucms/internal/qysystem"
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
	qysystem.NewApp("qycms-system").Run()
}
