FROM golang:1.20-alpine as builder

WORKDIR /code

COPY go.mod /code/
COPY go.sum /code/

RUN go mod download all

COPY . /code

RUN mkdir bin
RUN go build -o ./bin/server ./cmd/server 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /code/bin/server ./
COPY dist /root/dist
COPY gen /root/gen

CMD ["./server"]
