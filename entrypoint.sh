#!/bin/sh
curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh 
make migrate-up
chmod 755 /app/bin/go-starter-template && /app/bin/go-starter-template