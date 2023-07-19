package environment

import (
	"os"
	"testing"
)

func TestVariable(t *testing.T) {
	key := "TEST_XXX_YYY_VAR"
	os.Unsetenv(key)

	if res := Variable(key, "test"); res != "test" {
		t.Errorf("Expected default return, got %s", res)
	}

	os.Setenv(key, "abc")
	if res := Variable(key, "test"); res != "abc" {
		t.Errorf("Expected env var value, got %s", res)
	}
}
