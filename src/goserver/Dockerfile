FROM golang:1.15-alpine as builder
WORKDIR /src
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"
# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -gcflags='-N -l' -o /goserver .

FROM alpine as release

COPY --from=builder /goserver /goserver
#ENV PORT 8082
EXPOSE 8080
ENTRYPOINT ["/goserver"]
