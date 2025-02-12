FROM golang:1.22-alpine AS builder

COPY ./src/ /go/src/
WORKDIR /go/src/
RUN go mod download
RUN go build -o main .

FROM scratch

COPY --from=builder /go/src/main /main
CMD ["/main"]
