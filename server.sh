#!/bin/sh

cd ./cli-md-ui && ng serve &
gin --port 4201 --path . --build ./cli-md-backend/src/ --i --all &

wait