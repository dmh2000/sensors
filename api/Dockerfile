FROM golang:latest

EXPOSE 8002
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "."]
