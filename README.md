# API for Trainder
## Requirements
- [Go](https://go.dev)

## Development
1. `go mod tidy` to get all the Go requirements 
2. `go run .` to start the API 

The API will be available at [localhost:8080](http://localhost:8080) 

## Development using Docker
1. Build the image `docker build -t trainder-api .
2. Run the built images `docker run -p 8080:8080 -it trainder-api`

The API will be available at [localhost:8080](http://localhost:8080) 