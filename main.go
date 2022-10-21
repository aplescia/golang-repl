package main

import (
	"bufio"
	"fmt"
	"kv-store/transaction"
	"kv-store/util"
	"os"
	"strings"
)

var root *transaction.Transactions //basically just a stack

func main() {
	root = &transaction.Transactions{}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		res, err := util.ValidateInput(input)
		if err != nil {
			fmt.Fprint(os.Stderr, "Command is not valid. Please enter a valid command.\n")
			continue
		}

		var command = res.Command
		var key = res.Key
		var value = res.Value

		switch strings.ToLower(command) {
		case "quit":
			fmt.Fprint(os.Stderr, "exiting.")
			os.Exit(0)
		case "write":
			err := root.Write(key, value)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "read":
			val, err := root.Read(key)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(val)
			}
		case "delete":
			err := root.Delete(key)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "commit":
			err := root.Commit()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "start":
			root.StartTransaction()
		case "abort":
			err := root.Abort()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
