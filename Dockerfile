FROM golang:1.14-alpine as builder
ADD . /go/src/github.com/onuryilmaz/k8s-device-plugin-example
WORKDIR /go/src/github.com/onuryilmaz/k8s-device-plugin-example/cmd
RUN go build -v 
 
FROM alpine:latest
COPY --from=builder /go/src/github.com/onuryilmaz/k8s-device-plugin-example/cmd/cmd /usr/local/bin/k8s-device-plugin-example
CMD ["k8s-device-plugin-example"]
