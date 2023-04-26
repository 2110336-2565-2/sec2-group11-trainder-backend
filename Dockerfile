FROM golang:1.19.5-bullseye AS builder


COPY . /trainder-api
WORKDIR /trainder-api
RUN go mod tidy
RUN go build

FROM golang:1.19.5-bullseye AS runner
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /trainder-api/trainder-api /app

EXPOSE 8080

CMD ["/app/trainder-api"]
