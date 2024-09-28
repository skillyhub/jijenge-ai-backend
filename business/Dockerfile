# Dockerfile for business service
FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o business-service .
CMD ["./business-service"]