package main

import (
	"testing"
)

func TestEnv(t *testing.T) {
	t.Run("get enviroment variable", func(t *testing.T) {
		variable := GetEnvVariable("TEST_VARIABLE")
		expect := "TEST"

		if variable != expect {
			t.Errorf("got: %q ,expected: %q", variable, expect)
		}
	})
}
