FROM alpine:3.18

RUN apk add --no-cache \
    bash \
    python3 \
    curl \
    git \
    nginx \
    nodejs \
    npm \
    openssh \
    supervisor \
    yarn

RUN mkdir /fira

WORKDIR /fira

VOLUME /fira/workspace/node_modules
VOLUME /fira/workspace/apps/fira-app/.next
VOLUME /fira/workspace/apps/fira-site/.next
VOLUME /data

ENV NEXT_TELEMETRY_DISABLED 1
ENV DEBUG=true

# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs ./workspace/libs
COPY workspace/apps ./workspace/apps

RUN cd workspace && yarn install --pure-lockfile --non-interactive

COPY conf/dev-nginx.conf /etc/nginx/nginx.conf
COPY conf/dev-supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 8080

COPY . .

CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
