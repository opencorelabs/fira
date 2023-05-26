#!/bin/sh

set -eux

chmod 03777 "$DATA"
chmod 03777 "$HOME/app/bin"
chown -R "$USERNAME:$USERNAME" "$DATA"
chown -R "$USERNAME:$USERNAME" "$HOME/app/workspace"

gosu "$USERNAME" "$@"
