package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {
	log.Println("Sample code to test exec package")
	log.Println("Executing command - date")

	dateCmd := exec.Command("date")
	
	dateOut,err := dateCmd.Output()
	if err!=nil {
		panic(err)
		return
	}
	fmt.Printf("Command output is - %s",dateOut)
}
