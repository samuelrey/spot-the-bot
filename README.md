# spot-the-bot (aka Spot)
Build collaborative playlists with friends.

## Installation
In this section we cover how to setup applications for message & music services, and how to configure environment variables for Spot.

### **Discord**
Start by creating a new application for Spot in the [Discord developer portal](https://discord.com/developers/applications).
Next install Spot to your server. Replace your client ID and navigate to the authorization URL. Select the server on which you want to install Spot.

> https://discord.com/oauth2/authorize?client_id=REPLACEME&scope=bot&permissions=2048

Spot requires permission to send messages, hence 2048.

### **Spotify**
Start by creating a new application for Spot in the [Spotify developer portal](https://developer.spotify.com/dashboard/applications).
In order to authenticate, you will have to set the redirect URL to the address where you deploy Spot, eg `http://example.com/callback`.

### **Environment Variables**
Spot uses environment variables to interact with message & music services.

Service | Variable | Description
------- | -------- | -----------
Discord | `DISCORD_TOKEN` | Token found under the Bot section of the application.
Spotify | `CLIENT_ID` | ID found in the overview of the application.
Spotify | `SECRET` | Client secret found in the overview of the application.
Spotify | `REDIRECT_URL` | The address where you deploy Spot, eg `http://example.com/callback`.
Spotify | `STATE` | This variable is used to validate tokens received from Spotify. It can be set to anything, eg `spot-the-bot`. 

## Usage
In this section we cover how to run and interact with Spot.

### **Run**
`go build && ./spot-the-bot`

You will then be prompted to authorize Spot to create playlists for a Spotify user. This _could_ be your own personal account, but **we recommend you create a separate account for Spot**.

### **Commands**
At this point you can interact with Spot.

Command | Description
------- | -----------
`!join` | Jump into the queue of people that want to start a playlist.
`!leave` | Jump out of the queue.
`!list` | View the queue.
`!next` | Move yourself to the back of the queue if you are at the front.
`!create` | Make a new playlist if you are at the front of the queue. Spot will send you the link to the playlist. You will add a handful of songs (4-6) and share the playlist with the rest of the group.

**Anybody can add to the playlist, not just people in the queue!**

## Contibuting
I'm happy to work together--feel free to reach out!

Checkout the GitHub Issues for feature and improvement ideas.

Start at the **Installation** section. Most steps listed there are required for local development and testing. Any additional steps are covered below.

### **Spotify**
Add `http://localhost:8080/callback` to the list of redirect URLs in the developer portal and set the `REDIRECT_URL` environment variable accordingly.
