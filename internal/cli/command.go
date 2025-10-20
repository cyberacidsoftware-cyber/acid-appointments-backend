package cli

type Command interface {
	Name() string
	Init()
	Parse(args []string)
	Usage()
	Help()
	Execute(s *State) error
}
