FROM golang:1.11 as builder
RUN wget http://storage.googleapis.com/kubernetes-release/release/v1.13.3/bin/linux/amd64/kubectl -O /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN cd /go/src/github.com/nlopes/slack&&git checkout tags/v0.4.0&&cd -
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app -v ./...

FROM alpine:3.6 as alpine
RUN apk update && apk add -U --no-cache ca-certificates && rm -rf /var/cache/apk/*

FROM scratch
WORKDIR /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/bin/kubectl /usr/bin/kubectl
COPY --from=builder /go/src/app/app .
CMD ["/app"]
EXPOSE 80
