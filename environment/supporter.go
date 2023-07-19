package environment

// Supporter is an interface that can be implemented by any component to identify if it supports a given environment
type Supporter interface {
	SupportsEnvironment(environment Environment) bool
}

// Supports returns a list of components that support the given environment, and a boolean indicating if all components support the environment
func Supports(env Environment, components ...interface{}) ([]interface{}, bool) {
	passed := make([]interface{}, 0)
	allOk := true
	for _, c := range components {
		if s, ok := c.(Supporter); ok {
			if s.SupportsEnvironment(env) {
				passed = append(passed, c)
			} else {
				allOk = false
			}
		}
	}
	return passed, allOk
}
