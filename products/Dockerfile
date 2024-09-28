# Dockerfile for business service
FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o products-service .
CMD ["./products-service"]