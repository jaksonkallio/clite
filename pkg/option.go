package clite

type InteractiveConsoleOption func(*InteractiveConsoleOptions)

type InteractiveConsoleOptions struct {
	Logger Logger
}

var defaultInteractiveConsoleOptions = InteractiveConsoleOptions{
	Logger: PrintLogger{},
}

func WithLogger(logger Logger) InteractiveConsoleOption {
	return func(options *InteractiveConsoleOptions) {
		options.Logger = logger
	}
}
