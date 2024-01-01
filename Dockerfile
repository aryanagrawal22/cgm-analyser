# Dockerfile
FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY src/ .
COPY .env ./

RUN go get github.com/joho/godotenv
RUN go build -o main ./main.go
RUN go install github.com/pilu/fresh@latest

EXPOSE 9876

CMD ["./main"]