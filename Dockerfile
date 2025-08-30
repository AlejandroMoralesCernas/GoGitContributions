FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base-debian12 as runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]