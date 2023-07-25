FROM alpine:3.18

RUN apk add --no-cache \
    bash \
    python3 \
    py3-pip \
    py3-psycopg2 \
    curl \
    git \
    nginx \
    nodejs \
    npm \
    openssh \
    supervisor \
    yarn

RUN mkdir -p /fira/backend

WORKDIR /fira

VOLUME /fira/workspace/node_modules
VOLUME /fira/workspace/apps/fira-app/.next
VOLUME /fira/workspace/apps/fira-app/node_modules
VOLUME /data

ENV NEXT_TELEMETRY_DISABLED 1
ENV DEBUG=true
ENV DB_DIR=/data
ENV DJANGO_SUPERUSER_PASSWORD=password
ENV DJANGO_SUPERUSER_EMAIL=admin@opencorelabs.org
ENV DATABASE_URL=sqlite:////data/db.sqlite3

# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/apps ./workspace/apps

RUN --mount=type=cache,target=/root/.yarn cd workspace && YARN_CACHE_FOLDER=/root/.yarn yarn install --pure-lockfile --non-interactive

COPY backend/requirements.txt ./backend/

RUN --mount=type=cache,target=/root/.cache/pip pip install -r backend/requirements.txt

COPY conf/dev-nginx.conf /etc/nginx/nginx.conf
COPY conf/dev-supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 8080

COPY . .

CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
