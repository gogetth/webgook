# use multi-stage container to build docker image
FROM golang:1.11 AS builder
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# compile inside container
WORKDIR $GOPATH/src/github.com/gogetth/webgook
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /webhook .

# Run inside anothor one container
FROM docker:latest
COPY --from=builder /webhook ./
ENTRYPOINT ["./webhook"]