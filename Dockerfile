# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine

WORKDIR /web-sales

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /web-sales/websales ./cmd/main.go

#run app
ENTRYPOINT [ "/web-sales/websales" ]