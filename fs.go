package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func getTodoFiles() ([]fs.DirEntry, error) {
	files, err := os.ReadDir(todosDir)
	if err != nil {
		panic(err)
	}
	return files, err
}

func writeTodo(n string, d string) {
	os.WriteFile(todosDir+"/"+n+".todo", []byte(d), 0644)
	fmt.Println("Todo added successfully")
}

func listFiles(files []fs.DirEntry) {
	for i, file := range files {
		fmt.Printf("%-5v %v \n", strconv.Itoa(i)+":", file.Name())
		content, err := os.ReadFile(todosDir + "/" + file.Name())
		if err != nil {
			fmt.Printf("%-5v %v\n", "", "Error reading todo content")
			continue
		}
		fmt.Printf("%-5v %v\n", "", string(content))
	}
}
