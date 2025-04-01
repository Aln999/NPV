FROM golang:1.21 as builder

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /cluster-connectivity

FROM alpine:latest
WORKDIR /app
COPY --from=builder /cluster-connectivity .
COPY frontend/ ./frontend/

EXPOSE 8080
CMD ["./cluster-connectivity"]