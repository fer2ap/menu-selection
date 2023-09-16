package main

import (
	"fer2ap/menu-selection/operations"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	EXIT_OPTION = -1
)

func main() {
	menuOptions := []string{"Create new book", "Get book by ID", "Get books", "Update a book", "Delete a book"}
	var selected int
	for selected >= 0 {
		selected = 0
		selected = printMenu(&menuOptions)
		switch selected {
		case 0:
			var name string
			fmt.Println("Creating book. What's the name of the book?")
			_, err := fmt.Scan(&name)
			checkError(err)
			operations.CreateBook(name)
		case 1:
			bookId := printGetIdMenu()
			operations.GetBookById(bookId)
		case 2:
			var name string
			var page int
			var size int
			fmt.Println("What are you looking for?")
			fmt.Scan(&name)
			printGetIntInput("Select a page: ", &page)
			printGetIntInput("And select the number of entries: ", &size)
			operations.GetBooks(page, size, name)
		case 3:
			bookId := printGetIdMenu()
			operations.UpdateBook(bookId)
		case 4:
			bookId := printGetIdMenu()
			operations.DeleteBookById(bookId)
		}
		fmt.Println("Do you want to quit? [y]es/[n]o")
		var quit string
		fmt.Scanln(&quit)
		if quit == "y" {
			selected = EXIT_OPTION
		}
	}
}

func printGetIntInput(message string, input *int) {
	fmt.Println(message)
	fmt.Scanln(input)
}

func printGetIdMenu() int {
	fmt.Println("Enter the book id your are looking for:")
	var input int
	fmt.Scanln(&input)
	return input
}

func printMenu(menuOptions *[]string) int {
	exit := 1
	selected := 0
	lastOption := len(*menuOptions) - 1
	for exit != 0 {
		fmt.Println("Chose your option and press ENTER to confirm:")
		for i, v := range *menuOptions {
			if selected == i {
				fmt.Println(SELECTED, i, BLANK, v)
			} else {
				fmt.Println(BLANK, i, BLANK, v)
			}
		}
		var newSelection int
		var newOption string
		fmt.Scanln(&newOption)
		newSelection, err := strconv.Atoi(newOption)
		if err != nil {
			if strings.ToLower(newOption) == "ok" || strings.ToLower(newOption) == "" {
				return selected
			}
			newSelection = lastOption
		}
		if strings.ToLower(newOption) == "exit" {
			exit = 0
		}
		if newSelection > lastOption {
			selected = lastOption
		}
		if newSelection < 0 {
			selected = 0
		} else {
			selected = newSelection
		}
	}
	return EXIT_OPTION
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	SELECTED = ">"
	BLANK    = " "
)
