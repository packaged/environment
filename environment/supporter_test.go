package environment

import "testing"

type testSupporterOne struct {
	devTest bool
	env     Environment
}

func (t *testSupporterOne) SupportsEnvironment(env Environment) bool {
	return t.env == env || t.devTest && env.IsDevOrTest()
}

func TestSupports(t *testing.T) {
	devSupport := &testSupporterOne{devTest: true, env: Development}
	unitSupport := &testSupporterOne{devTest: true, env: UnitTest}
	unitOnlySupport := &testSupporterOne{devTest: false, env: UnitTest}
	prodSupport := &testSupporterOne{env: Production}

	if passed, ok := Supports(Development, devSupport, unitSupport); !ok {
		t.Error("Expected all environments to pass, got ", passed)
	}

	if passed, ok := Supports(Development, devSupport, unitSupport, unitOnlySupport); ok {
		t.Error("Expected unitOnly to fail, got ", passed)
	}

	if passed, ok := Supports(Production, prodSupport); !ok {
		t.Error("Expected production to pass, got ", passed)
	}

	passed, _ := Supports(Development, devSupport, unitSupport, unitOnlySupport)
	if len(passed) != 2 {
		t.Error("Expected 2 passed environments, got ", len(passed))
	}
}
