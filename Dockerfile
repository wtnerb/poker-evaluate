# built from template https://www.cloudreach.com/blog/containerize-this-golang-dockerfiles/

# when go to https and things inevitably break, check out https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
# make a builder container
FROM golang:alpine as builder
RUN mkdir /build
ADD . /go/src/build
RUN apk add git
WORKDIR /go/src/build 
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /main .
# make the actual end container that only holds that static binary
FROM scratch
COPY --from=builder /main /app/
WORKDIR /app
CMD ["./main"]