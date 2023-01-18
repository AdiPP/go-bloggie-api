#!/bin/sh

make init-schema
make init-seeder

chmod +x /app/bin/bloggie
/app/bin/bloggie