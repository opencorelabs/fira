FROM golang:1.20-buster

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git make openssh-client nodejs && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=true
ENV FIRA_CLIENT_DIR=/app/client
ENV FIRA_BIND_HTTP=0.0.0.0:8080

COPY go.mod go.sum Makefile ./
RUN mkdir client
COPY workspace/apps/fira-app/package.json workspace/apps/fira-app/yarn.lock ./workspace/apps/fira-app/
RUN make reqs

COPY . .

CMD ["air"]
