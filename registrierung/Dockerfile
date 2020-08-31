FROM golang:1 as builder

WORKDIR /go/src/registration

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o registration ./cmd

FROM alpine:latest
COPY --from=builder /go/src/registration .
CMD ["./registration"]
