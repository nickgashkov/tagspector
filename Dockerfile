FROM golang:1.14-alpine3.11

ENV CGO_ENABLED="0"

WORKDIR /app/
COPY . /app/

RUN go build ./...

CMD ["/app/tagspector"]
