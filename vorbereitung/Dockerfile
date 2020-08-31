FROM golang:1 as builder

WORKDIR /go/src/vorbereitung

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o vorbereitung ./cmd

FROM alpine:latest
COPY --from=builder /go/src/vorbereitung .
CMD ["./vorbereitung"]
