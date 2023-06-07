FROM golang:1.20-buster

ENV USERNAME=devuser
ENV HOME=/home/lib/fira
ENV DATA=/var/run/fira

RUN set -eux; \
	addgroup --gid 70 $USERNAME; \
	adduser --uid 70 --gid 70 --home $HOME --shell /bin/sh $USERNAME; \
	mkdir -p $HOME/bin; \
	chown -R $USERNAME:$USERNAME $HOME

RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git make openssh-client nodejs sudo gosu && \
    go install github.com/cosmtrek/air@latest

RUN mkdir $HOME/go
RUN mkdir $HOME/app
RUN mkdir $HOME/npm
RUN npm config set prefix $HOME/npm

WORKDIR $HOME/app

RUN chown -R $USERNAME:$USERNAME $HOME

ENV PATH="$HOME/npm/bin:$HOME/go/bin:${PATH}"
ENV NODE_PATH="$HOME/npm/lib/node_modules:${NODE_PATH}"
ENV GOPATH=$HOME/go

VOLUME $HOME/app/workspace/node_modules
VOLUME $HOME/app/workspace/apps/fira-app/.next
VOLUME $GOPATH
VOLUME $DATA

ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=true
ENV FIRA_CLIENT_DIR=$HOME/app/workspace
ENV FIRA_BIND_HTTP=0.0.0.0:8080
ENV FIRA_EMBEDDED_POSTGRES_DATA_PATH=$DATA/data
ENV FIRA_EMBEDDED_POSTGRES_BINARIES_PATH=$DATA/bin
ENV FIRA_EMBEDDED_POSTGRES_RUNTIME_PATH=$DATA/runtime

COPY go.mod go.sum Makefile ./

# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs ./workspace/libs
COPY workspace/apps ./workspace/apps

RUN make reqs

COPY . .

CMD ["air"]
