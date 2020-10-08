# build
FROM golang:1.15 AS builder
WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -o apple-maintained-bot cmd/server/main.go

# run
FROM golang:1.15
COPY --from=builder /app/apple-maintained-bot /
CMD ["/apple-maintained-bot"]