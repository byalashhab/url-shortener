# URL Shortener

A simple URL shortener server with two endpoints:

- `POST /api/v1/shorten`  
  Accepts a JSON body `{ "longURL": string }` and returns a shortened URL.

- `GET /api/v1/{id}`  
  Redirects to the original long URL based on the provided ID.

## License

This project is licensed under the [MIT License](LICENSE).