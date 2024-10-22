package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Help   bool
	Add    string
	Del    int
	Edit   string
	Update int
	List   bool
}

func NewCmdflags() *CmdFlags {
	cf := CmdFlags{}

	flag.BoolVar(&cf.Help, "help", false, "List existing commands")
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing todo, enter #index and specify a new title. \"id:new title\"")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by #index to delete")
	flag.IntVar(&cf.Update, "update", -1, "Specify a todo #index to update")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !isValidFlag(arg) {
			fmt.Printf("Unknown flag: %s\n", arg)
			fmt.Println("try --help to know more")

			os.Exit(0)
		}
	}

	flag.Parse()
	return &cf
}

func isValidFlag(flag string) bool {

	validFlags := []string{
		"-help", "--help",
		"-add", "--add",
		"-edit", "--edit",
		"-del", "--del",
		"-update", "--update",
		"-list", "--list",
	}

	if idx := strings.Index(flag, "="); idx != -1 {
		flag = flag[:idx]
	}

	for _, validFlag := range validFlags {
		if flag == validFlag {
			return true
		}
	}

	return false
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Printf("Error, invalid format for edit.\nCorrect Format: \"id:new title\" ")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Printf("Error, Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Update != -1:
		todos.toggle(cf.Update)

	case cf.Del != -1:
		todos.delete(cf.Del)

	case cf.Help:
		fmt.Println("usage:")
		fmt.Println("--help\t\t| List existing commands")
		fmt.Println("--add\t\t| Add new task")
		fmt.Println("--del\t\t| Delete an existing task")
		fmt.Println("--update\t| Check/Uncheck existing task")
		fmt.Println("--edit\t\t| Edit an existing task")
	}
}