package ecosystem

import "strings"

type Package struct {
	Name        string
	Version     string
	IsAIPackage bool
}

type DepGraph struct {
	Packages []Package
}

func (g *DepGraph) ByName(name string) *Package {
	lower := strings.ToLower(name)
	for i := range g.Packages {
		if strings.ToLower(g.Packages[i].Name) == lower {
			return &g.Packages[i]
		}
	}
	return nil
}

type Parser interface {
	Name() string
	DetectFiles() []string
	Parse(filename string, content []byte) (*DepGraph, error)
}
