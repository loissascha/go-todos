package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var todosDir = "todos"

func main() {
	// flags
	add := flag.Bool("a", false, "Add a new todo")
	addLong := flag.Bool("add", false, "Add a new todo")

	ls := flag.Bool("ls", false, "List all todos")
	lsLong := flag.Bool("list", false, "List all todos")

	rm := flag.Bool("rm", false, "Remove a todo")
	rmLong := flag.Bool("remove", false, "Remove a todo")

	flag.Parse()

	if *add || *addLong {
		addOperation()
	} else if *ls || *lsLong {
		listOperation()
	} else if *rm || *rmLong {
		removeOperation()
	} else {
		listOperation()
	}
}

func removeOperation() {
	reader := bufio.NewReader(os.Stdin)

	files, _ := getTodoFiles()
	listFiles(files)

	i, _ := readInput("Number of todo to remove: ", reader)
	index, err := strconv.ParseInt(i, 0, 64)
	if err != nil {
		fmt.Println("This is not a valid number!")
		removeOperation()
		return
	}

	fmt.Printf("Deleting file from index: %v \n", index)
	fn := files[index].Name()
	os.Remove(todosDir + "/" + fn)
}

func listOperation() {
	files, _ := getTodoFiles()
	listFiles(files)
}

func addOperation() {
	reader := bufio.NewReader(os.Stdin)
	todo, _ := readInput("Enter a new todo: ", reader)
	desc, _ := readInput("Enter a description: ", reader)

	_, err := os.Stat(todosDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(todosDir, 0755)
		if err != nil {
			fmt.Println("Error creating todos directory")
			return
		}
	}

	writeTodo(todo, desc)
}

func readInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	text, err := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSpace(text)
	return text, err
}
