package python

import (
	"os"
	"strings"
	"testing"
)

func TestParseRequirements(t *testing.T) {
	content, err := os.ReadFile("testdata/requirements.txt")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	graph, err := ParseRequirements(content)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(graph.Packages) != 3 {
		t.Fatalf("expected 3 packages, got %d", len(graph.Packages))
	}
	lc := graph.ByName("langchain")
	if lc == nil {
		t.Fatal("langchain not found")
	}
	if lc.Version != "0.1.0" {
		t.Errorf("langchain version: want 0.1.0, got %s", lc.Version)
	}
	if !lc.IsAIPackage {
		t.Error("langchain should be AI package")
	}
	req := graph.ByName("requests")
	if req == nil {
		t.Fatal("requests not found")
	}
	if req.IsAIPackage {
		t.Error("requests should not be AI package")
	}
}

func TestParsePyproject(t *testing.T) {
	content, err := os.ReadFile("testdata/pyproject.toml")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	graph, err := ParsePyproject(content)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if graph.ByName("anthropic") == nil {
		t.Error("anthropic not found")
	}
	if graph.ByName("transformers") == nil {
		t.Error("transformers not found")
	}
	if graph.ByName("python") != nil {
		t.Error("python key should not be in packages")
	}
}

func TestEcosystemPluginParse(t *testing.T) {
	plugin := &EcosystemPlugin{}

	reqContent := []byte("flask==3.0.0\n")
	graph, err := plugin.Parse("requirements.txt", reqContent)
	if err != nil {
		t.Fatalf("Parse requirements.txt: %v", err)
	}
	if graph == nil {
		t.Fatal("expected non-nil graph for requirements.txt")
	}

	pyprojectContent := []byte("[tool.poetry.dependencies]\nflask = \"3.0.0\"\n")
	graph2, err := plugin.Parse("pyproject.toml", pyprojectContent)
	if err != nil {
		t.Fatalf("Parse pyproject.toml: %v", err)
	}
	if graph2 == nil {
		t.Fatal("expected non-nil graph for pyproject.toml")
	}

	_, err = plugin.Parse("unknown.txt", []byte("something"))
	if err == nil {
		t.Fatal("expected error for unknown.txt, got nil")
	}
}

func TestParseRequirementsNonEqualSpecifier(t *testing.T) {
	content := []byte("requests>=2.0.0\nflask==3.0.0\n")
	graph, err := ParseRequirements(content)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	// flask (==) must be present
	if graph.ByName("flask") == nil {
		t.Error("flask should be in graph")
	}

	// requests (>=) must be skipped from packages
	if graph.ByName("requests") != nil {
		t.Error("requests should NOT be in packages (non-== specifier)")
	}

	// warning must mention the skipped line
	found := false
	for _, w := range graph.Warnings {
		if strings.Contains(w, "requests>=2.0.0") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected warning about skipped line 'requests>=2.0.0', got warnings: %v", graph.Warnings)
	}
}

func TestEcosystemPluginDetectFiles(t *testing.T) {
	p := &EcosystemPlugin{}
	files := p.DetectFiles()
	if len(files) != 2 {
		t.Fatalf("expected 2 detect files, got %d", len(files))
	}
	found := map[string]bool{}
	for _, f := range files {
		found[f] = true
	}
	if !found["requirements.txt"] {
		t.Error("requirements.txt not in DetectFiles")
	}
	if !found["pyproject.toml"] {
		t.Error("pyproject.toml not in DetectFiles")
	}
}
