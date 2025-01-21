# urlshorten

## Testing with curl

- `curl -H "Content-Type: application/json" -d '{"url": "https://www.youtube.ca"}' -X POST http://localhost:8080/shorten`

- `curl -X GET http://localhost:8080/{shortcode}`
- `curl -X GET http://localhost:8080/stats/{shortcode}`
