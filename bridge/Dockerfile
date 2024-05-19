FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
CMD ["go", "run", ".", "w/sin","w/triangle","w/square"]


