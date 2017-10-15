FROM golang:1.8
MAINTAINER https://github.com/andersfylling

WORKDIR /go/src/github.com/andersfylling/ccdb
COPY . .

# Get Glide for package management
RUN curl https://glide.sh/get | sh
RUN glide install

ENV CCDB_TOKEN DISCORD_TOKEN_HERE_PLEASE
ENV CCDB_COMMANDPREFIX $

# RUN go run main.go

# docker build -t discord-bot-bitcoin .
# docker run -d discord-bot-bitcoin
