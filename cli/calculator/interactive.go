package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func interactive() {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		if key == '\x03' {
			// Quit without error on Ctrl^C
			exit(oldState)
		}
		return "", 0, false
	}
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit(oldState)
			}
			panic(err)
		}

		text = strings.TrimSpace(text)
		if text == "exit" || text == "quit" {
			break
		}
	}

}

func exit(state *terminal.State) {
	terminal.Restore(0, state)
	fmt.Println()
	os.Exit(0)
}
