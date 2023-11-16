package railway_test

import (
	"testing"
	"vrata/legacy/internal/railway"
)

func Test_Create_Step(t *testing.T) {

	step := railway.Step{Name: "Create"}

	if step.Name != "Create" {
		t.Error("Expected a step")
	}
}
