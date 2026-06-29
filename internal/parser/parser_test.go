package parser_test

import (
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
	"google.golang.org/genai"
)

func TestParserPackageExists(t *testing.T) {
	// Verify the parser package is importable and the project compiles.
	// Importing parser and calling a sentinel function confirms linkage.
	got := parser.PackageName()
	if got != "parser" {
		t.Fatalf("expected package name %q, got %q", "parser", got)
	}
}

func TestGenaiDependencyAvailable(t *testing.T) {
	// Verify that the google.golang.org/genai dependency is resolvable.
	// Importing the genai package and referencing a type proves the
	// dependency is correctly added to go.mod and downloadable.
	var client *genai.Client
	if client != nil {
		t.Fatal("unexpected non-nil client")
	}
}

