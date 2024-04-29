FROM golang:latest

ARG SHAPE
WORKDIR /app
COPY . .
COPY ./${SHAPE} ./config.yaml
COPY ../../.env .env
CMD ["go", "run", "."]
