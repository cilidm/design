package main

import (
	"fmt"
	"time"
)

//go:generate rsrc -ico resource/favicon32.ico -manifest resource/goversioninfo.exe.manifest -o main.syso

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i,time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
