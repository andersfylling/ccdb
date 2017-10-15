# BTC Bitfinex [CCDB]
Crypto-Currency-Discord-Bot written in Go, using the Unison library to display a crypto currency.
This upates every 0.5sec or slower depending on the realtime updates.

Discord limits updates to 2 times per second, so obviusly this can't be realtime. I have currently just added a check that blocks any updates that takes place before 0.5s has passed, so it might feel slow at times.


## How it looks
[![ae43420d27c87146dbf64423f6b6ef87.png](http://pichoster.net/images/2017/08/19/ae43420d27c87146dbf64423f6b6ef87.png)](http://pichoster.net/image/aWow8)

With a seperate role given to the bot


## Add to your server
This uses docker to run, so make sure you set the environment variables `CCDB_TOKEN` and optinally `CCDB_COMMANDPREFIX`. The command prefix is `$` by default.
This bot doesn't need any permissions, so no need to waste time on that.
To add it to your server, click here: https://discordapp.com/oauth2/authorize?&client_id=348565617005297687&scope=bot&permissions=0

## Compile it yourself
The username of the bot is "BTC Bitfinex" for now.
Obviusly you can download and compile this project on your own. Use glide to handle dependencies on you're set.
