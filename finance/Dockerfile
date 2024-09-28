# Dockerfile for business service
FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o finance-service .
CMD ["./finance-service"]