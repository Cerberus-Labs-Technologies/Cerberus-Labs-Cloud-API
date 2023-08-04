package commands

type SubCommand struct {
	Name    string
	Execute func(args ...string)
}

type Command struct {
	Name        string
	Description string
	Usage       string
	HasArgs     bool
	SubCommands []SubCommand
	Execute     func(args ...string)
}

type CommandInterface interface {
	Execute()
	getDescription() string
	getUsage() string
	getName() string
	getSubCommands() []SubCommand
}
