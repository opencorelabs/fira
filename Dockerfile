# build the backend
FROM golang:1.20-alpine as backend

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download all

COPY . /code

RUN mkdir bin
RUN go build -o ./bin/server ./cmd/server 

# build the client deps
FROM node:16-alpine as clientdeps
RUN apk add --no-cache libc6-compat nasm autoconf automake bash libltdl libtool gcc make g++ zlib-dev
WORKDIR /code
COPY client/package.json client/yarn.lock ./
RUN yarn install

FROM node:16-alpine as client
ENV NEXT_TELEMETRY_DISABLED 1
WORKDIR /code
COPY --from=clientdeps /code/node_modules ./node_modules
COPY ./client .
RUN yarn build

# final request serving image
FROM node:16-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy backend resources
COPY --from=backend /code/bin/server ./
COPY dist ./dist
COPY gen ./gen

# copy client resources
RUN mkdir /root/client
COPY --from=client /code/public ./client/public
COPY --from=client /code/package.json ./client/package.json
COPY --from=client /code/.next ./client/.next
COPY --from=client /code/node_modules ./client/node_modules

ENV NEXT_TELEMETRY_DISABLED 1

CMD ["./server"]
