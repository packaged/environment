package environment

import (
	"fmt"
	"strconv"
)

// Name is a type that represents an environment variable name
type Name string

// Returns the string representation of the Name
func (v Name) String() string {
	return string(v)
}

// Value returns the value of the environment variable, or an empty string if not set
func (v Name) Value() string {
	return Variable(string(v), "")
}

// ValueInt returns the value of the environment variable as an integer, or an error if not set
func (v Name) ValueInt() (int, error) {
	return strconv.Atoi(Variable(string(v), ""))
}

// WithDefault returns the value of the environment variable, or a default value if not set
func (v Name) WithDefault(defaultValue string) string {
	return Variable(string(v), defaultValue)
}

// ValueOrPanic returns the value of the environment variable, or panics if not set
func (v Name) ValueOrPanic() string {
	if val := v.Value(); val != "" {
		return val
	}
	panic(fmt.Sprintf("environment variable [%s] not set", v.String()))
}

type Var struct {
	name         string
	defaultValue string
}

func NewDefault(name, defaultValue string) Var {
	return Var{name: name, defaultValue: defaultValue}
}

func (v Var) Name() string    { return v.name }
func (v Var) Default() string { return v.defaultValue }
func (v Var) Value() string   { return Variable(v.name, v.defaultValue) }
