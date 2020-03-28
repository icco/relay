# relay

[![Go Report Card](https://goreportcard.com/badge/github.com/icco/relay)](https://goreportcard.com/report/github.com/icco/relay)

A service that runs as a discord bot and proxies web requests to it.

## Install

 - Create a new app at https://discordapp.com/developers/applications, and in the app, create a bot user.
 - Get the client ID from the app settings page and then visit this page to connect the bot:
   - https://discordapp.com/api/oauth2/authorize?client_id=$CLIENT_ID&scope=bot&permissions=51264
 - Start the app
 - Send any json to `/hook`
   - For example: `curl -svL -H "Content-Type: application/json" -d '{"test":"bar","hi":"xyz"}' http://localhost:8080/hook`
   - It'll show up in your discord.
   - [![photo](https://icco.imgix.net/photos/2020/18afc1ec-7ea4-4e8b-88e7-f1e74786b539.png?auto=format%2Ccompress&w=300)](https://icco.imgix.net/photos/2020/18afc1ec-7ea4-4e8b-88e7-f1e74786b539.png?auto=format%2Ccompress)
