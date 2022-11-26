package clite

import "fmt"

var DefaultHelpCommand Command = Command{
	Keyword:     "help",
	Description: "Lists available commands.",
	Aliases: []string{
		"?",
	},
	Executable: func(console *Console, args map[string]string) error {
		fmt.Printf("%s\n", console.HelpString())
		return nil
	},
}

var DefaultQuitCommand Command = Command{
	Keyword:     "quit",
	Description: "Quits the program.",
	Aliases: []string{
		"exit",
	},
	Executable: func(console *Console, args map[string]string) error {
		console.End()
		return nil
	},
}
