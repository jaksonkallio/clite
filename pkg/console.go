package clite

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Console struct {
	commands map[string]nameMappedCommand
	ended    bool
	logger   Logger
}

type nameMappedCommand struct {
	Command *Command
	IsAlias bool
}

type helpEntry struct {
	Command *Command
	Entry   string
}

func NewConsole(commands []Command, opts ...InteractiveConsoleOption) *Console {
	interactiveConsoleOptions := defaultInteractiveConsoleOptions
	for _, opt := range opts {
		opt(&interactiveConsoleOptions)
	}
	commandsMap := make(map[string]nameMappedCommand, len(commands))
	for _, command := range commands {
		command := command
		nameMappedCommand := nameMappedCommand{
			Command: &command,
			IsAlias: false,
		}
		commandsMap[command.Keyword] = nameMappedCommand
		nameMappedCommand.IsAlias = true
		for _, alias := range command.Aliases {
			commandsMap[alias] = nameMappedCommand
		}
	}
	return &Console{
		commands: commandsMap,
		logger:   interactiveConsoleOptions.Logger,
	}
}

func (console *Console) End() {
	console.ended = true
}

func (console *Console) HelpString() string {
	sortedCommandDefs := make([]*Command, 0, len(console.commands))
	for _, nameMappedCommand := range console.commands {

		if !nameMappedCommand.IsAlias {
			sortedCommandDefs = append(sortedCommandDefs, nameMappedCommand.Command)
		}
	}
	sort.Slice(sortedCommandDefs, func(i int, j int) bool {
		return sortedCommandDefs[i].Keyword < sortedCommandDefs[j].Keyword
	})
	var maxCommandEntryLen int
	helpEntries := make([]helpEntry, 0, len(sortedCommandDefs))
	for _, command := range sortedCommandDefs {
		entryStr := strings.Builder{}
		entryStr.WriteString(command.Keyword)
		for _, arg := range command.Arguments {
			entryStr.WriteString(fmt.Sprintf(" <%s>", arg))
		}
		if len(entryStr.String()) > maxCommandEntryLen {
			maxCommandEntryLen = len(entryStr.String())
		}
		helpEntries = append(
			helpEntries,
			helpEntry{
				Command: command,
				Entry:   entryStr.String(),
			},
		)
	}
	helpStr := strings.Builder{}
	for _, helpEntry := range helpEntries {
		helpStr.WriteString(fmt.Sprintf("%-*s -> %s\n", maxCommandEntryLen, helpEntry.Entry, helpEntry.Command.Description))
	}
	return helpStr.String()
}

func (console *Console) Listen() {
	stdInReader := bufio.NewReader(os.Stdin)
	for !console.ended {
		inputStr, _ := stdInReader.ReadString('\n')
		inputStr = strings.TrimSpace(inputStr)
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
		if len(command.Command.Arguments) != len(args) {
			console.logger.Errorf("expected %d arguments, %d arguments were provided", len(command.Command.Arguments), len(args))
			break
		}
		argMap := make(map[string]string)
		for i := range args {
			argMap[command.Command.Arguments[i]] = args[i]
		}
		err := command.Command.Executable(console, argMap)
		if err != nil {
			console.logger.Errorf("executing command failed: %w", err)
		}
	}
}
