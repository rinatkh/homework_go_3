package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func runCommand(t *testing.T, path string) string {
	t.Helper()
	cmd := exec.Command("go", "run", path)
	cmd.Dir = "../.."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run %s failed: %v\n%s", path, err, string(out))
	}
	return strings.TrimSpace(string(out))
}

func TestLoopsExample(t *testing.T) {
	out := runCommand(t, "./cmd/01_loops")
	for _, part := range []string{"sum=15", "factorial=120", "even=3", "negativeIndex=2", "Fizz"} {
		if !strings.Contains(out, part) {
			t.Fatalf("loops output %q does not contain %q", out, part)
		}
	}
}

func TestFunctionsExample(t *testing.T) {
	out := runCommand(t, "./cmd/02_functions")
	for _, part := range []string{"name=Мария Иванова", "div=3", "ok=true", "sum=10", "avg=4.0"} {
		if !strings.Contains(out, part) {
			t.Fatalf("functions output %q does not contain %q", out, part)
		}
	}
}

func TestArraysExample(t *testing.T) {
	out := runCommand(t, "./cmd/03_arrays")
	for _, part := range []string{"sum=19", "avg=3.8", "max=9", "contains=true", "changed=[100 1 9 1 3]"} {
		if !strings.Contains(out, part) {
			t.Fatalf("arrays output %q does not contain %q", out, part)
		}
	}
}

func TestSlicesExample(t *testing.T) {
	out := runCommand(t, "./cmd/04_slices")
	for _, part := range []string{"sum=18", "avg=3.6", "even=[2 8 2]", "unique=[5 2 8 1]", "sorted=[1 2 3 4]"} {
		if !strings.Contains(out, part) {
			t.Fatalf("slices output %q does not contain %q", out, part)
		}
	}
}
