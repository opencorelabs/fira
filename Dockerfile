# build the backend
FROM golang:1.20-alpine as backend

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download all

COPY . /code

RUN mkdir bin
RUN go build -o ./bin/server ./cmd/server

RUN mkdir bin/pg
ENV FIRA_EMBEDDED_POSTGRES_BINARIES_PATH=/code/bin/pg
RUN go run ./cmd/bootstrap

# build the client deps
FROM node:16-alpine as clientdeps
RUN apk add --no-cache libc6-compat nasm autoconf automake bash libltdl libtool gcc make g++ zlib-dev
WORKDIR /code
# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/
COPY workspace/apps/fira-app ./workspace/apps/fira-app/

WORKDIR /code/workspace
RUN yarn install --pure-lockfile --non-interactive --cache-folder ./ycache; rm -rf ./ycache

# build the client app
FROM node:16-alpine as client

ARG NEXTAUTH_URL

ENV NEXT_TELEMETRY_DISABLED 1
ENV NEXTAUTH_URL=$NEXTAUTH_URL

WORKDIR /code

# setup workspace
RUN mkdir workspace
COPY --from=clientdeps /code/workspace/node_modules ./workspace/node_modules
COPY --from=clientdeps /code/workspace/package.json ./workspace/package.json
COPY --from=clientdeps /code/workspace/yarn.lock ./workspace/yarn.lock
COPY ./workspace/.eslintrc.js ./workspace/.eslintrc.js
COPY ./workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/
COPY ./workspace/apps/fira-app ./workspace/apps/fira-app/

WORKDIR /code/workspace
# build libs
RUN yarn workspace @fira/api-sdk build
# build app
RUN yarn workspace @fira/app build

# final request serving image
FROM node:16-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy backend resources
COPY --from=backend /code/bin/server ./
COPY --from=backend /code/bin/pg ./pg
COPY dist ./dist
COPY gen ./gen

# copy client resources
RUN mkdir /root/client
COPY --from=client /code/workspace/apps/fira-app/public ./client/public
COPY --from=client /code/workspace/apps/fira-app/package.json ./client/package.json
COPY --from=client /code/workspace/apps/fira-app/.next ./client/.next
COPY --from=client /code/workspace/node_modules ./client/node_modules
COPY ./pg/migrations ./pg/migrations

ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=false
ENV FIRA_CLIENT_DIR=/root/client
ENV FIRA_MIGRATIONS_DIR=/root/pg/migrations
ENV FIRA_EMBEDDED_POSTGRES_BINARIES_PATH=/root/pg/bin

CMD ["./server"]
