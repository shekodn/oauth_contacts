# OAuth Contacts

Git repository for *The OAuth Waltz* [article](https://medium.com/@sergiodn/the-oauth-2-0-waltz-957879e5316d).

## Description
Sample application used to retrieve your google's account contacts using OAuth 2.0

## Setup
### Google Developer
1. Create the authorization server [here](https://console.developers.google.com/?angularJsUrl=)
2. In the Authorized redirect URIs don't forget to include the application's callback URI: http://localhost:8000/callback

### Your machine
1. Clone this git repository in the respective $GOPATH
2. Create an ```.env``` file with the respective Client ID and Client Secret
```
GOOGLE_CLIENT_ID=xxx
GOOGLE_CLIENT_SECRET=xxx
PORT=8000
```
3. In your terminal just do ```make run```
4. Go to the client that should be running in http://localhost:8000/contacts
5. Click on Import Contacts From Google
Select your Google account and give permission to the client to get the respective resources
6. Sit back and wait for the client to retrieve the resource owner's contacts
