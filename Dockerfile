FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./main.go

FROM alpine:3.19

# ✅ เพิ่มบรรทัดนี้
RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/system ./system

EXPOSE 8080
CMD ["./server"]