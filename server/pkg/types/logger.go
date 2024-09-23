package types

type Logger interface {
	Debug(string, ...any)
	Error(string, ...any)
	Info(string, ...any)
}
