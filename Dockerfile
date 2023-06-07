# build the backend
FROM golang:1.20-alpine as backend

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download all

COPY . /code

RUN mkdir bin
RUN go build -o ./bin/fira ./cmd/fira

# build the client deps
FROM node:20-alpine as clientdeps
RUN apk add --no-cache libc6-compat nasm autoconf automake bash libltdl libtool gcc make g++ zlib-dev
WORKDIR /code
# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/
COPY workspace/apps/fira-app ./workspace/apps/fira-app/
COPY workspace/apps/fira-site ./workspace/apps/fira-site/

WORKDIR /code/workspace
RUN yarn install --pure-lockfile --non-interactive

# build the client app
FROM node:20-alpine as client

ARG NEXT_PUBLIC_BASE_URL
ARG NEXT_PUBLIC_VERIFICATION_BASE_URL

ENV NEXT_TELEMETRY_DISABLED 1
ENV NEXT_PUBLIC_BASE_URL=$NEXT_PUBLIC_BASE_URL
ENV NEXT_PUBLIC_VERIFICATION_BASE_URL=$NEXT_PUBLIC_VERIFICATION_BASE_URL

WORKDIR /code

# setup workspace
RUN mkdir workspace
COPY --from=clientdeps /code/workspace/node_modules ./workspace/node_modules
COPY --from=clientdeps /code/workspace/package.json ./workspace/package.json
COPY --from=clientdeps /code/workspace/yarn.lock ./workspace/yarn.lock
COPY ./workspace/.eslintrc.js ./workspace/.eslintrc.js
COPY ./workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/
COPY ./workspace/apps/fira-app ./workspace/apps/fira-app/
COPY ./workspace/apps/fira-site ./workspace/apps/fira-site/

WORKDIR /code/workspace
# build libs
RUN yarn workspace @fira/api-sdk build
# build apps
RUN yarn workspace @fira/app build
RUN yarn workspace @fira/site build

# final request serving image
FROM node:20-alpine

ENV USERNAME=fira
ENV HOME=/home/lib/fira
ENV DATA=/var/run/fira
ENV LANG en_US.utf8

RUN set -eux; \
	addgroup -g 70 -S $USERNAME; \
	adduser -u 70 -S -D -G $USERNAME -H -h $HOME -s /bin/sh $USERNAME; \
	mkdir -p $HOME/bin; \
	chown -R $USERNAME:$USERNAME $HOME

RUN apk --no-cache add ca-certificates su-exec

WORKDIR $HOME

# copy backend resources
COPY --from=backend /code/bin/fira ./bin/
COPY dist ./dist
COPY gen ./gen

# set up embedded postgres
RUN mkdir -p $DATA/pg/data && mkdir -p $DATA/pg/runtime
RUN chown -R $USERNAME:$USERNAME $DATA && chmod 3777 $DATA
VOLUME $DATA/pg/data

ENV FIRA_EMBEDDED_POSTGRES_DATA_PATH=$DATA/pg/data
ENV FIRA_EMBEDDED_POSTGRES_BINARIES_PATH=$DATA/pg/bin
ENV FIRA_EMBEDDED_POSTGRES_RUNTIME_PATH=$DATA/pg/runtime

RUN su-exec $USERNAME $HOME/bin/fira bootstrap && rm -rf $HOME/.embedded-postgres-go

# copy client resources
RUN mkdir $HOME/client
COPY --from=client /code/workspace/apps/fira-app/public ./client/public
COPY --from=client /code/workspace/apps/fira-app/package.json ./client/package.json
COPY --from=client /code/workspace/apps/fira-app/.next ./client/.next
COPY --from=client /code/workspace/node_modules ./client/node_modules
COPY --from=client /code/workspace/apps/fira-site/out ./dist/fira-site
COPY ./pg/migrations ./pg/migrations
COPY ./scripts/entrypoint.sh ./entrypoint.sh

ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=false
ENV FIRA_CLIENT_DIR=$HOME/client
ENV FIRA_MIGRATIONS_DIR=$HOME/pg/migrations

STOPSIGNAL SIGINT
EXPOSE 8080

ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]
CMD ["./bin/fira", "serve"]
