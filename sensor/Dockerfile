FROM golang:latest

ARG SHAPE
WORKDIR /app
COPY . .
COPY ./${SHAPE} ./config.yaml
RUN go mod download
CMD ["go", "run", "."]
