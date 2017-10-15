FROM golang:1.8-onbuild
MAINTAINER https://github.com/andersfylling

docker build -t discord-bot-bitcoin .
docker run -d discord-bot-bitcoin
