package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adelberteng/currency_converter/routers"
)

func main() {
	cmd, err := exec.Command("python3", "rate_crawler/main.py").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(cmd))
	routers.Init()
}
