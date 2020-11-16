# Gopher Translator API

This API translates words into the Gopher language so the little friends can understand us.

Per specification, the API has 3 endpoints:

- `POST` `/word`: Translates a single English word into the Gopher language.
Send a JSON request with the following content
```json
{
    "english-word": "<your English word>"
}
```

- `POST` `/sentence`: Translates a single English sentence into the Gopher language.
Send a JSON request with the following content
```json
{
    "english-sentence": "<your English sentence>"
}
```

- `GET` `/history`: Returns the history of all translated words and sentences since the moment the API has been started

## Running

To run the API, just issue
```go
go run main.go
```

Optionally, a single parameter can be passed:
```sh
--port <your port number>
```
