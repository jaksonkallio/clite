package clite

type Command struct {
	Keyword     string
	Description string
	Aliases     []string
	Arguments   []string
	Executable  func(console *Console, args map[string]string) error
}
