FROM golang:latest
RUN apt-get update && apt-get install -y default-mysql-client
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]