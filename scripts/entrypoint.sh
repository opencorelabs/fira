#!/bin/sh

chown -R "$USERNAME:$USERNAME" "$DATA"
chown -R "$USERNAME:$USERNAME" "$FIRA_CLIENT_DIR"

su-exec "$USERNAME" "$@"
