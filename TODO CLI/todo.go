package main

import (
	"TODO_CLI/services/cli_parser"
	"TODO_CLI/services/todo_handler"
	"fmt"
	"os"
	"strings"
)

func main() {
	service, strArg, err := cli_parser.Parse()
	if err != nil {
		fmt.Println(strings.Title(err.Error()))
		os.Exit(0)
	}
	todo_handler.Handle(service, strArg)
}
