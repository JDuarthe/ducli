package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	switch os.Args[1] {
	case "lol":
		err := exec.Command("C:/Riot Games/Riot Client/RiotClientServices.exe").Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func printUsage() {
	fmt.Println("Uso: duCli [lol]")
}
