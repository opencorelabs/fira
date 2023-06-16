# build the client deps
FROM node:18-alpine as clientdeps
RUN apk add --no-cache libc6-compat nasm autoconf automake bash libltdl libtool gcc make g++ zlib-dev
WORKDIR /code
# root workspace
COPY workspace/package.json workspace/yarn.lock ./workspace/
COPY workspace/libs/fira-api-sdk ./workspace/libs/fira-api-sdk/
COPY workspace/apps/fira-app ./workspace/apps/fira-app/

WORKDIR /code/workspace
RUN yarn install --pure-lockfile --non-interactive

# build the client app
FROM node:18-alpine as client

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

WORKDIR /code/workspace
# build libs
RUN yarn workspace @fira/api-sdk build
# build app
RUN yarn workspace @fira/app build

# final request serving image
FROM alpine:3.18

RUN apk add --no-cache python3 py3-pip py3-psycopg2 nginx nodejs npm yarn supervisor

RUN mkdir -p /fira/client && mkdir /fira/backend

WORKDIR /fira

COPY --from=client /code/workspace/apps/fira-app/public ./client/public
COPY --from=client /code/workspace/apps/fira-app/package.json ./client/package.json
COPY --from=client /code/workspace/apps/fira-app/.next ./client/.next
COPY --from=client /code/workspace/node_modules ./client/node_modules

COPY backend /fira/backend

RUN pip install -r backend/requirements.txt && pip install gunicorn

RUN cd backend && python manage.py collectstatic --noinput && chown -R 1000 static

ENV NEXT_TELEMETRY_DISABLED 1
ENV DEBUG=false

EXPOSE 8080

COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
