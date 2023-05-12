FROM golang:1.20-buster

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git make openssh-client nodejs && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=true
ENV FIRA_CLIENT_DIR=/app/workspace/apps/fira-app
ENV FIRA_BIND_HTTP=0.0.0.0:8080

COPY go.mod go.sum Makefile ./

# Setup workspace
RUN mkdir workspace


# Copy root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/

# Copy libs first
COPY workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/

# Copy apps
COPY workspace/apps/fira-app ./workspace/apps/fira-app/

RUN make reqs

COPY . .

CMD ["air"]
