package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	start()
	prisonCell()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What do you want to do? ")
		text, _ := reader.ReadString('\n')

		if text == "end\n" {
			fmt.Println("Goodbye for now.")
			break
		}

		fmt.Printf("I don't know how to %s", text)
	}
}

func start() {
	content, err := ioutil.ReadFile("introduction")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n%s\n", content)
}

func prisonCell() {
	content, err := ioutil.ReadFile("prison-cell")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n%s\n", content)
}
