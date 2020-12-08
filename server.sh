#!/bin/sh

ng serve &
gin --port 4201 --path . --build ../cli-md-backend/src/main --i --all &

wait