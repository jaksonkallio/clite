package main

import (
	"fmt"

	clite "github.com/jaksonkallio/clite/pkg"
)

func main() {
	console := clite.NewConsole(
		[]clite.Command{
			{
				Keyword:     "hello",
				Description: "Will say hello to the user and will say their favorite color.",
				Aliases: []string{
					"hi",
					"👋",
					"howdy",
				},
				Arguments: []string{
					"name",
					"favorite_color",
				},
				Executable: func(console *clite.Console, args map[string]string) error {
					fmt.Printf("Hello %s! 🎨 Looks like your favorite color is %s.\n", args["name"], args["favorite_color"])
					return nil
				},
			},
			clite.DefaultHelpCommand,
			clite.DefaultQuitCommand,
		},
	)

	console.Listen()
}
