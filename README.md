# BTC Bitfinex [CCDB]
Crypto-Currency-Discord-Bot written in Go, using [DisGord](https://github.com/andersfylling/disgord). It updates every 12 second due to Discord rate limits for status updates.

![ccdb bot in action](https://raw.githubusercontent.com/andersfylling/ccdb/master/looks.png)

## Add to your server
[Click here.](https://discordapp.com/oauth2/authorize?&client_id=348565617005297687&scope=bot&permissions=0)

Or you can compile and host it yourself. Use the Dockerfile provided or the docker image at `hub.docker.com/andersfylling/ccdb-disgord`. You must provide the environment variable named `CCDB_TOKEN`, which is your discord bot token. This bot doesn't need any permissions, so no need to waste time on that.
