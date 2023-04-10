package main

import (
	"fmt"
	"os/exec"
)

func main() {

	fileName := "./example/file_create.go"
	cmd := exec.Command("go", "run", fileName)
	cmd.CombinedOutput()

	fileName = "main.go"
	cmd = exec.Command("go", "run", fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))

}
