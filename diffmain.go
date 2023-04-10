package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	go func() {
		fileName := "./example/file_create.go"
		cmd := exec.Command("go", "run", fileName)
		cmd.CombinedOutput()
	}()

	go func() {
		fileName := "main.go"
		cmd := exec.Command("go", "run", fileName)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(output)
	}()

	time.Sleep(3 * time.Second)

}
