package clite

import (
	"bufio"
	"os"
	"strings"
)

type Console struct {
	commands map[string]Command
	ended    bool
	logger   Logger
}

func NewConsole(commands []Command, opts ...InteractiveConsoleOption) *Console {
	interactiveConsoleOptions := defaultInteractiveConsoleOptions

	for _, opt := range opts {
		opt(&interactiveConsoleOptions)
	}

	commandsMap := make(map[string]Command, len(commands))
	for _, command := range commands {
		commandsMap[command.Keyword] = command
	}

	return &Console{
		commands: commandsMap,
		logger:   interactiveConsoleOptions.Logger,
	}
}

func (console *Console) End() {
	console.ended = true
}

func (console *Console) Listen() {
	stdInReader := bufio.NewReader(os.Stdin)

	for !console.ended {
		inputStr, _ := stdInReader.ReadString('\n')
		inputStr = strings.TrimSpace(inputStr)

		// If command is empty, do nothing.
		if len(inputStr) == 0 {
			continue
		}

		inputStrParts := strings.Split(inputStr, " ")

		command, found := console.commands[inputStrParts[0]]
		if !found {
			console.logger.Errorf("command %q not found", inputStrParts[0])
			break
		}

		args := inputStrParts[1:]

		if len(command.Arguments) != len(args) {
			console.logger.Errorf("expected %d arguments, %d arguments were provided", len(command.Arguments), len(args))
			break
		}

		argMap := make(map[string]string)
		for i := range args {
			argMap[command.Arguments[i]] = args[i]
		}

		go func() {
			err := command.Executable(console, argMap)
			if err != nil {
				console.logger.Errorf("executing command failed: %w", err)
			}
		}()
	}
}
