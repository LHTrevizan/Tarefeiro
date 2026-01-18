#build stage
FROM golang::1.24-alpine AS build

WORKDIR /app

copy go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tarefeiro ./cmd/tarefeiro/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/tarefeiro .

ENTRYPOINT ["./tarefeiro"]