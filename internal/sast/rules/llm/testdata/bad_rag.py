import chromadb
import openai

client = openai.OpenAI()
chroma = chromadb.Client()
collection = chroma.get_collection("docs")
results = collection.query(query_texts=["user question"])
docs = results["documents"][0]
prompt = f"Answer based on: {docs}"
response = client.chat.completions.create(
    model="gpt-4",
    messages=[{"role": "user", "content": prompt}]
)
