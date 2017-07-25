package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main1() {
	exmpl_2()
}

func exmpl_1() {
	cmd := exec.Command("tput", "-S")
	cmd.Stdin = bytes.NewBufferString("clear\ncup 5 10")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Hello")
}

func exmpl_2() {
	var cmd *exec.Cmd
	for i := 1; i < 100; i++ {
		cmd = exec.Command("tput", "-S")
		cmd.Stdin = bytes.NewBufferString("clear\ncup 3 10")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Hello [" + strings.Repeat("#", i) + "]")
		time.Sleep(80 * time.Millisecond)
	}

}
