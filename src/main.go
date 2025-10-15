package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	l27 "cobaltdb-local/L-27"
)

func main() {
	fmt.Println("Interpreter DB CLI")
	inter, err := l27.NewInterpreter("/db")
	if err != nil {
		fmt.Println("Error initializing interpreter:", err)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.ToLower(line) == "exit" {
			//inter.Save()
			fmt.Println("Saving changes to disk...")
			fmt.Println("bye, see you later.")
			break
		}

		lex := l27.NewLexer(line)
		parser := l27.NewParser(lex)

		cmd, err := parser.ParseCommand()
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		err = inter.Execute(cmd)
		if err != nil {
			fmt.Println("Error ejecutando comando:", err)
		}
	}
}
