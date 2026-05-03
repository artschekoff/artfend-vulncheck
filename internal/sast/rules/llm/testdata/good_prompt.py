import openai
import html

client = openai.OpenAI()
user_input = input("Ask anything: ")
safe_input = html.escape(user_input)[:500]

response = client.chat.completions.create(
    model="gpt-4",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": safe_input},
    ]
)
