package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("TO DO Command Line")

	// inisialisasi i/o buat nerima input
	consoleReader := bufio.NewReader(os.Stdin)

	// inisialisasi array untuk menampung to-do list
	fmt.Print("Masukkan Task : ")
	list, _ := consoleReader.ReadString('\n')

	fmt.Println("Task should to do : ")
	fmt.Println(list)
}
