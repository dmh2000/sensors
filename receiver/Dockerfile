FROM golang:latest

ARG SHAPE
WORKDIR /app
COPY . .
COPY ../../.env .env
CMD ["go", "run", "."]
