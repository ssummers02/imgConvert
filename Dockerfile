FROM golang:1.17

RUN go version
ENV GOPATH=/
ENV CGO_CFLAGS_ALLOW=-Xpreprocessor
RUN apt update
RUN apt install libvips -y
RUN apt install libvips-dev -y

COPY ./ ./
RUN go mod download
RUN go build -o app ./cmd/main.go

EXPOSE 8080 8080
ENTRYPOINT ["./app", "-mode","web"]
