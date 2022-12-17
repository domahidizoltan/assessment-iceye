FROM golang:1.19.3-alpine3.16 AS builder
ENV CGO_ENABLED=0
WORKDIR $GOPATH/src/iceye/
COPY . .
RUN go install ./...
RUN go build -o /go/bin/poker main.go

FROM scratch
COPY --from=builder /go/bin/* /bin/
ENTRYPOINT ["/bin/poker"]