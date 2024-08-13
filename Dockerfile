FROM golang:1.22-alpine3.19 as builder

WORKDIR /build

COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o fiberApp ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /build/fiberApp ./
COPY --from=builder /build/.env ./

CMD ["./fiberApp"]