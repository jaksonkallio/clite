package clite

type Command struct {
	Keyword    string
	Arguments  []string
	Executable func(console *Console, args map[string]string) error
}
