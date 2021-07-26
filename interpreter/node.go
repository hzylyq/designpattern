package interpreter

type node interface {
	parse() error
}
