# use multi-stage container to build docker image
FROM golang:1.11 AS builder
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# compile inside container
WORKDIR $GOPATH/src/github.com/gogetth/webgook
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -o /main .

# Run inside anothor one container
FROM golang:1.11
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
