package python

var knownAIPackages = map[string]struct{}{
	"langchain":           {},
	"langchain-core":      {},
	"langchain-community": {},
	"openai":              {},
	"anthropic":           {},
	"google-generativeai": {},
	"transformers":        {},
	"sentence-transformers": {},
	"llama-index":         {},
	"llama-cpp-python":    {},
	"chromadb":            {},
	"pinecone-client":     {},
	"weaviate-client":     {},
	"guardrails-ai":       {},
	"nemoguardrails":      {},
	"rebuff":              {},
	"langfuse":            {},
	"mlflow":              {},
	"bentoml":             {},
}

func isAIPackage(name string) bool {
	_, ok := knownAIPackages[name]
	return ok
}
