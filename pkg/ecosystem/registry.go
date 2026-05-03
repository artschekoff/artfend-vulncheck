package ecosystem

import (
	"fmt"
	"sync"
)

var (
	mu      sync.RWMutex
	parsers = map[string]Parser{}
)

func Register(name string, p Parser) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := parsers[name]; exists {
		panic(fmt.Sprintf("ecosystem: duplicate registration for %q", name))
	}
	parsers[name] = p
}

func Get(name string) (Parser, bool) {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := parsers[name]
	return p, ok
}
