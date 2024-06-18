FROM golang:1.21.4

WORKDIR /opt/app

COPY ./cmd ./cmd
COPY ./go.* .
RUN go build -o app ./cmd/...

ENTRYPOINT [ "/opt/app/app" ]