package environment

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	env, err := Parse("invalidenv")
	if env != "" {
		t.Error("Expected empty string for an invalid environment")
	}
	if err == nil || !strings.Contains(err.Error(), "invalid input") {
		t.Error("Expected error for an invalid environment")
	}

	env, err = Parse("proDUCtion")
	if env != Production {
		t.Error("Expected production environment")
	}

	env, err = Parse("sandbox")
	if env != Sandbox {
		t.Error("Expected sandbox environment")
	}

	for _, dVal := range []string{"development", "dev"} {
		env, err = Parse(dVal)
		if env != Development {
			t.Error("Expected development environment")
		}
		if !env.IsDevelopment() {
			t.Error("Expected development environment")
		}
		if !env.IsDevOrTest() {
			t.Error("Expected development environment")
		}
	}

	for _, dVal := range []string{"local"} {
		env, err = Parse(dVal)
		if env != Local {
			t.Error("Expected local environment")
		}
		if !env.IsLocal() {
			t.Error("Expected development environment")
		}
		if !env.IsDevOrTest() {
			t.Error("Expected development environment")
		}
	}

	for _, dVal := range []string{"test", "unittest"} {
		env, err = Parse(dVal)
		if env != UnitTest {
			t.Error("Expected test environment")
		}
		if !env.IsTestCase() {
			t.Error("Expected test case to pass")
		}
		if !env.IsDevOrTest() {
			t.Error("Expected dev or test to be true")
		}
	}

	for _, dVal := range []string{"itest", "integrationtest"} {
		env, err = Parse(dVal)
		if env != IntegrationTest {
			t.Error("Expected IntegrationTest environment")
		}
		if !env.IsTestCase() {
			t.Error("Expected test case to pass")
		}
		if !env.IsDevOrTest() {
			t.Error("Expected dev or test to be true")
		}
	}
}

func TestValidators(t *testing.T) {
	tests := []struct {
		env               Environment
		allowTestData     bool
		isDevOrTest       bool
		isTestCase        bool
		isDevelopment     bool
		isSandbox         bool
		isProduction      bool
		isUnitTest        bool
		isIntegrationTest bool
		isLocal           bool
	}{
		{Production, false, false, false, false, false, true, false, false, false},
		{Sandbox, true, false, false, false, true, false, false, false, false},
		{Local, true, true, false, false, false, false, false, false, true},
		{Development, true, true, false, true, false, false, false, false, false},
		{UnitTest, true, true, true, false, false, false, true, false, false},
		{IntegrationTest, true, true, true, false, false, false, false, true, false},
		{"unknown45", true, false, false, false, false, false, false, false, false},
	}

	for _, test := range tests {
		if test.env.AllowTestData() != test.allowTestData {
			t.Errorf("Expected AllowTestData in %s to be %t", test.env, test.allowTestData)
		}
		if test.env.IsDevelopment() != test.isDevelopment {
			t.Errorf("Expected IsDevelopment in %s to be %t", test.env, test.isDevelopment)
		}
		if test.env.IsLocal() != test.isLocal {
			t.Errorf("Expected isLocal in %s to be %t", test.env, test.isLocal)
		}
		if test.env.IsTestCase() != test.isTestCase {
			t.Errorf("Expected IsTestCase in %s to be %t", test.env, test.isTestCase)
		}
		if test.env.IsDevOrTest() != test.isDevOrTest {
			t.Errorf("Expected IsDevOrTest in %s to be %t", test.env, test.isDevOrTest)
		}
		if test.env.IsSandbox() != test.isSandbox {
			t.Errorf("Expected IsSandbox in %s to be %t", test.env, test.isSandbox)
		}
		if test.env.IsProduction() != test.isProduction {
			t.Errorf("Expected IsProduction in %s to be %t", test.env, test.isProduction)
		}
		if test.env.IsUnitTest() != test.isUnitTest {
			t.Errorf("Expected IsProduction in %s to be %t", test.env, test.isProduction)
		}
		if test.env.IsIntegrationTest() != test.isIntegrationTest {
			t.Errorf("Expected IsProduction in %s to be %t", test.env, test.isProduction)
		}
	}
}

func TestGetIdentifier(t *testing.T) {
	tests := []struct {
		env    Environment
		expect string
	}{
		{Production, "P"},
		{Sandbox, "S"},
		{Development, "D"},
		{UnitTest, "U"},
		{IntegrationTest, "I"},
		{Local, "L"},
		{"unknown45", "UN"},
		{"X", ""}, // Single character environments are not supported with identifiers
	}

	for _, test := range tests {
		ident := test.env.GetIdentifier()
		if ident != test.expect {
			t.Errorf("Expected identifier %s for %s, got %s", test.expect, test.env, ident)
		}
	}
}
