#!/bin/sh

sudo chown -R $USERNAME:$USERNAME /home/$USERNAME/pg

exec "$@"
