FROM golang:1.19.5-bullseye
COPY . /trainder-api
WORKDIR /trainder-api
RUN go mod tidy
CMD go run .