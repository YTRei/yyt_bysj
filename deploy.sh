#! /bin/bash

cp -R ./tmpl ./bin/

mkdir ./bin/videos

cd bin
nohup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"