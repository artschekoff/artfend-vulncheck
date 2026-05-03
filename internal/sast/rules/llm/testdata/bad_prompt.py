import openai

client = openai.OpenAI()
user_input = input("Ask anything: ")

response = client.chat.completions.create(
    model="gpt-4",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": f"Answer this: {user_input}"},
    ]
)
