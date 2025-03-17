package environment

import (
	"os"
	"testing"
)

var TestEnvVar Name = "TEST_ENV_VAR"
var TestEnvVarMissing Name = "TEST_ENV_VAR_397237dd*"

func TestKnown(t *testing.T) {

	val := TestEnvVar.String()
	if val != "TEST_ENV_VAR" {
		t.Error("Failed to convert env var to string")
	}

	err := os.Setenv(TestEnvVar.String(), "6969")
	if err != nil {
		t.Error("Failed to set env var")
	}

	val = TestEnvVar.Value()
	if val != "6969" {
		t.Error("Failed to get env var value")
	}

	valInt, err := TestEnvVar.ValueInt()
	if err != nil {
		t.Error("Failed to convert env var value")
	}
	if valInt != 6969 {
		t.Error("Failed to convert env var value")
	}

	val = TestEnvVar.WithDefault("xx")
	if val != "6969" {
		t.Error("Failed to get env var value")
	}

	err = os.Setenv(TestEnvVar.String(), "")
	if err != nil {
		t.Error("Failed to set env var")
	}

	val = TestEnvVar.WithDefault("xx")
	if val != "xx" {
		t.Error("Failed to get env var default")
	}
}

func TestValueOrPanic(t *testing.T) {
	assertPanic(t, "calling ValueOrPanic on an unset environment variable", func() { TestEnvVarMissing.ValueOrPanic() })
	_ = os.Setenv(TestEnvVarMissing.String(), "abc")
	val := TestEnvVarMissing.ValueOrPanic()
	if val != "abc" {
		t.Error("Failed to get env value from ValueOrPanic")
	}
}

func assertPanic(t *testing.T, reason string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic: %s", reason)
		}
	}()
	f()
}

func TestDefault(t *testing.T) {

	const DefaultedVar = "TEST_ENV_VAR_DFLT"

	v := NewDefault(DefaultedVar, "default-value")
	if v.Name() != DefaultedVar {
		t.Error("Failed to get env var name")
	}

	if v.Default() != "default-value" {
		t.Error("Failed to get env var default")
	}

	if v.Value() != "default-value" {
		t.Error("Failed to get env var value")
	}

	_ = os.Setenv(DefaultedVar, "6969")

	if v.Value() != "6969" {
		t.Error("Failed to get env var value")
	}
}
