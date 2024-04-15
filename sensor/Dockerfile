FROM golang:latest

ARG SHAPE
WORKDIR /app
COPY . .
COPY ./${SHAPE} ./config.yaml
CMD ["go", "run", "."]
