package cqrs

type command struct {
	TypeInfo
	Request
}

type Command interface {
	isCommand()

	Request
	TypeInfo
}

func NewCommandByT[T any]() Command {
	c := &command{
		TypeInfo: NewTypeInfo[T](),
		Request:  NewRequest(),
	}

	return c
}

// https://github.com/EventStore/EventStore-Client-Go/blob/master/esdb/position.go#L29
func (c *command) isCommand() {}

func IsCommand(obj interface{}) bool {
	if _, ok := obj.(Command); ok {
		return true
	}

	return false
}
