package ecosystem

var parsers = map[string]Parser{}

func Register(name string, p Parser) { parsers[name] = p }
func Get(name string) (Parser, bool) { p, ok := parsers[name]; return p, ok }
