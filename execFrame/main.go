package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	localFile = "/home/c019/Develop/logFrame/"
)

func main() {

	i := 0

	for {
		local := localFile + "logFrame_" + strconv.Itoa(i) + ".txt"
		i++
		cmdOut, err := exec.Command("cat", local).Output()

		if err != nil {
			panic(err.Error())
		}

		os.Stdout.WriteString(string(cmdOut))
		os.Stdout.Sync()
		time.Sleep(50 * time.Millisecond)
	}

}
