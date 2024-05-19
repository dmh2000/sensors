FROM golang:latest

ARG SHAPE
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
COPY ./${SHAPE} ./config.yaml
CMD ["go", "run", "."]
