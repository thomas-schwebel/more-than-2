# more-than-2

Using the Spotify API.

Scan your own playlists or likes to check out from which artists you like 2 or more tracks. Gives you a hint about which artists you might wanna check more.

## Prerequesites: 
- Spotify Dev Account + 1 App setup
Or:
- 1 valid API access token with scope: `playlist-read-private`

## How to use:
### Getting an API access token:
Using code form: https://github.com/spotify/web-api-examples, in `./web-api-examples/authentication/authorization_code`:
- run `npm i`
- setup `client_id`, `client_secret` and `redirect_uri` in `app.js` (use the values from your Spotify App)
- then run `node app.js` and follow the Oauth2 flow to get an API access token

### Running the app:
In `./cmd/more-than-2/main.go` replace:
- `accessToken`: you token from above
- `owner`: your spotify id

Then run: `go run ./cmd/more-than-2/main.go`

Output example:
```
******* 13 MATCHES *******
- Teho
- JAY-Z

******* 10 MATCHES *******
- Todd Terje
- Disiz
- Justice

******* 9 MATCHES *******
- Christopher Schwarzwälder
- James Blake

******* 8 MATCHES *******
- Christian Smith
- NTO
- Kavinsky
- Fatima Yamaha

[...]
```
