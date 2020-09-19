FROM golang:1.15
ADD . /go/src/github.com/furikuri/mtls
WORKDIR /go/src/github.com/furikuri/mtls
RUN go get ./
RUN go build

FROM alpine:3.10
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /root/
COPY --from=0 /go/bin/mtls .
EXPOSE 80
ENTRYPOINT ["/root/mtls"]