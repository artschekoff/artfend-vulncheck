package python

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/artfend/vulncheck/pkg/ecosystem"
)

// compile-time interface guard
var _ ecosystem.Parser = (*EcosystemPlugin)(nil)

func ParseRequirements(content []byte) (*ecosystem.DepGraph, error) {
	var graph ecosystem.DepGraph
	for _, raw := range bytes.Split(content, []byte("\n")) {
		line := strings.TrimSpace(string(raw))
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "==", 2)
		if len(parts) != 2 {
			graph.Warnings = append(graph.Warnings, fmt.Sprintf("skipped line: %s", line))
			continue
		}
		name := strings.ToLower(strings.TrimSpace(parts[0]))
		version := strings.TrimSpace(parts[1])
		graph.Packages = append(graph.Packages, ecosystem.Package{
			Name:        name,
			Version:     version,
			IsAIPackage: isAIPackage(name),
		})
	}
	return &graph, nil
}

type pyprojectTOML struct {
	Tool struct {
		Poetry struct {
			Dependencies map[string]interface{} `toml:"dependencies"`
		} `toml:"poetry"`
	} `toml:"tool"`
}

func ParsePyproject(content []byte) (*ecosystem.DepGraph, error) {
	var doc pyprojectTOML
	if _, err := toml.Decode(string(content), &doc); err != nil {
		return nil, fmt.Errorf("toml decode: %w", err)
	}
	var graph ecosystem.DepGraph
	for k, v := range doc.Tool.Poetry.Dependencies {
		name := strings.ToLower(k)
		if name == "python" {
			continue
		}
		version := ""
		switch val := v.(type) {
		case string:
			version = strings.TrimLeft(val, "^~>=<!")
		case map[string]interface{}:
			if vs, ok := val["version"].(string); ok {
				version = strings.TrimLeft(vs, "^~>=<!")
			}
		default:
			graph.Warnings = append(graph.Warnings, fmt.Sprintf("skipped dep %s: unsupported value type", name))
			continue
		}
		graph.Packages = append(graph.Packages, ecosystem.Package{
			Name:        name,
			Version:     version,
			IsAIPackage: isAIPackage(name),
		})
	}
	return &graph, nil
}

type EcosystemPlugin struct{}

func (p *EcosystemPlugin) Name() string { return "python" }

func (p *EcosystemPlugin) DetectFiles() []string {
	return []string{"requirements.txt", "pyproject.toml"}
}

func (p *EcosystemPlugin) Parse(filename string, content []byte) (*ecosystem.DepGraph, error) {
	switch {
	case strings.HasSuffix(filename, "requirements.txt"):
		return ParseRequirements(content)
	case strings.HasSuffix(filename, "pyproject.toml"):
		return ParsePyproject(content)
	default:
		return nil, fmt.Errorf("unsupported file: %s", filename)
	}
}
