#!/bin/sh

chown -R "$USERNAME:$USERNAME" "$DATA"

su-exec "$USERNAME" "$@"
