FROM golang:1.20-buster

ARG USERNAME=devuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG PUID=$USER_UID
ARG PGID=$USER_GID

RUN groupadd --gid $USER_GID $USERNAME
RUN useradd --uid $USER_UID --gid $USER_GID -m $USERNAME
RUN groupmod -o -g $PGID "$USERNAME"
RUN usermod -o -u $PUID "$USERNAME"

RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git make openssh-client nodejs sudo && \
    go install github.com/cosmtrek/air@latest

RUN echo "$USERNAME:$USERNAME" | chpasswd -e
RUN usermod -aG sudo $USERNAME
RUN echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER $USERNAME

RUN mkdir /home/$USERNAME/go
RUN mkdir /home/$USERNAME/app
RUN mkdir /home/$USERNAME/npm
RUN npm config set prefix /home/$USERNAME/npm

WORKDIR /home/$USERNAME/app

ENV PATH="/home/$USERNAME/npm/bin:/home/$USERNAME/go/bin:${PATH}"
ENV NODE_PATH="/home/$USERNAME/npm/lib/node_modules:${NODE_PATH}"

ENV GOPATH=/home/$USERNAME/go
ENV NEXT_TELEMETRY_DISABLED 1
ENV FIRA_DEBUG=true
ENV FIRA_CLIENT_DIR=/app/workspace/apps/fira-app
ENV FIRA_BIND_HTTP=0.0.0.0:8080

COPY go.mod go.sum Makefile ./

# Setup workspace
RUN mkdir workspace

# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs ./workspace/libs
COPY workspace/apps ./workspace/apps
RUN sudo chown -R $USERNAME:$USERNAME /home/$USERNAME

RUN make reqs

COPY . .

ENV USERNAME=$USERNAME
ENTRYPOINT ["/bin/sh", "scripts/entrypoint.sh"]
CMD ["./bin/server"]
