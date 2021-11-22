package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
)

func htmlGenerator(cssify Cssify) {
	fmt.Println("HTML GENERATOR")
	template, err := template.ParseFiles("./template/cssify.html")
	if err != nil {
		fmt.Println("Error: File could not be opened")
		log.Fatal(err)
	}

	file := createFile()

	writer := bufio.NewWriter(file)

	if err := template.Execute(writer, cssify); err != nil {
		fmt.Println("Error: Template could not be executed")
		log.Fatal(err)
	}

	writer.Flush()
}

func createFile() *os.File {
	prettyPrint("HTML File Created", true)

	err := os.MkdirAll("./out/", 0700)
	if err != nil {
		fmt.Println("Error: Directory could not be created")
		log.Fatal(err)
	}

	file, err := os.Create("./out/index.html")

	if err != nil {
		fmt.Println("Error: index.html could not be created")
		log.Fatal(err)
	}
	return file
}

func fileOpen(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		log.Fatal(err)
	}
	return file
}
