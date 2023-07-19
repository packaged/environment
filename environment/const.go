// Package environment provides basic environment definitions
package environment

import (
	"fmt"
	"strings"
)

// Environment is a type that represents an environment name
type Environment string

// Production live environment
const Production Environment = "production"

// Sandbox public test environment
const Sandbox Environment = "sandbox"

// Development Development Environments
const Development Environment = "development"

// Local Local Development Environments
const Local Environment = "local"

// UnitTest Test Environments - same repo testing
const UnitTest Environment = "unittest"

// IntegrationTest - distributed testing, combining multiple repos & external services
const IntegrationTest Environment = "integrationtest"

// AllowTestData Check to allow for test data per environment
func (e Environment) AllowTestData() bool { return e != Production }

// IsDevOrTest returns true if the environment is development, local, or a test case
func (e Environment) IsDevOrTest() bool { return e.IsDevelopment() || e.IsLocal() || e.IsTestCase() }

// IsLocal returns true if the environment is local
func (e Environment) IsLocal() bool { return e == Local }

// IsDevelopment returns true if the environment is development
func (e Environment) IsDevelopment() bool { return e == Development }

// IsTestCase returns true if the environment is a test case either unit or integration
func (e Environment) IsTestCase() bool { return e == UnitTest || e == IntegrationTest }

// IsUnitTest returns true if the environment is a unit test
func (e Environment) IsUnitTest() bool { return e == UnitTest }

// IsIntegrationTest returns true if the environment is an integration test
func (e Environment) IsIntegrationTest() bool { return e == IntegrationTest }

// IsSandbox returns true if the environment is a sandbox
func (e Environment) IsSandbox() bool { return e == Sandbox }

// IsProduction returns true if the environment is production
func (e Environment) IsProduction() bool { return e == Production }

// Parse parses a string into an environment
func Parse(input string) (Environment, error) {
	switch strings.ToLower(input) {
	case "p", "production":
		return Production, nil
	case "s", "sandbox":
		return Sandbox, nil
	case "d", "development", "dev":
		return Development, nil
	case "l", "local", "local-dev":
		return Local, nil
	case "i", "itest", "integrationtest":
		return IntegrationTest, nil
	case "u", "utest", "unittest", "test":
		return UnitTest, nil
	}
	return "", fmt.Errorf("invalid input: %s", input)
}

// GetIdentifier returns the identifier for the environment, this is a single character for known environments, and 2 characters for unknown environments
func (e Environment) GetIdentifier() string {
	switch e {
	case Production:
		return "P"
	case Sandbox:
		return "S"
	case Development:
		return "D"
	case UnitTest:
		return "U"
	case IntegrationTest:
		return "I"
	case Local:
		return "L"
	}

	// default to the first letter of the provided environment
	if len(e) > 1 {
		return strings.ToUpper(string(e[:2]))
	}
	return ""
}
