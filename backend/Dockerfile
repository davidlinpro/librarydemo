FROM golang:alpine as builder
WORKDIR /build 
COPY . /build
RUN apk add git && \
    go get -d ./... && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o backend .

FROM scratch
WORKDIR /app
COPY --from=builder /build/backend /app
EXPOSE 8081
CMD ["./backend"]
