package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//func main() {
//	c := exec.Command("cat", "text.txt")
//	c.Stdin = os.Stdin
//	c.Stdout = os.Stdout
//	c.Stderr = os.Stderr
//	c.Run()
//}

func main() {
	appendToFile()
}

func appendToFile() {
	f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(" Konstantin"); err != nil {
		panic(err)
	}
}

func writeToFile() {
	data := []byte("My name is Konstantin")
	err := ioutil.WriteFile("text.txt", data, 0600)
	if err != nil {
		fmt.Println(err)
	}
}

func readFile() {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
