package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestGreetDefault(t *testing.T) {
	var buf bytes.Buffer
	if err := greet(&buf, "World"); err != nil {
		t.Fatalf("greet returned error: %v", err)
	}
	got := buf.String()
	want := "Hello, World!\n"
	if got != want {
		t.Errorf("greet(World) = %q, want %q", got, want)
	}
}

func TestGreetCustomName(t *testing.T) {
	var buf bytes.Buffer
	if err := greet(&buf, "Jan"); err != nil {
		t.Fatalf("greet returned error: %v", err)
	}
	if got := buf.String(); got != "Hello, Jan!\n" {
		t.Errorf("greet(Jan) = %q", got)
	}
}

func TestRootCmdRun(t *testing.T) {
	cmd := NewRootCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--name", "Go"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	if !strings.Contains(buf.String(), "Hello, Go!") {
		t.Errorf("unexpected output: %q", buf.String())
	}
}
