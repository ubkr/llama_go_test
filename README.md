# llama_go_test

## Setting the API Endpoint

You can set the API Endpoint for accessing the ollama API using either a command line argument or an environment variable.

### Using Command Line Argument

To set the API Endpoint using a command line argument, run the program with the API Endpoint as the first argument:

```sh
go run main.go <API_ENDPOINT>
```

Replace `<API_ENDPOINT>` with the actual API Endpoint URL.

### Using Environment Variable

To set the API Endpoint using an environment variable, set the `OLLAMA_API_ENDPOINT` environment variable before running the program:

```sh
export OLLAMA_API_ENDPOINT=<API_ENDPOINT>
go run main.go
```

Replace `<API_ENDPOINT>` with the actual API Endpoint URL.
