# GCP IAP Oauth2 with Golang Sample
This is a sample code for access GAE standard through IAP.

## Before User
1. [Setup a GAE web server.](https://cloud.google.com/appengine/docs/standard/go/building-app/)
2. [Setup a IAP for GAE.](https://cloud.google.com/iap/docs/app-engine-quickstart)
3. [Create a service account only for access GAE through IAP.](https://cloud.google.com/iap/docs/managing-access)
4. Download the service account key json file.
5. [Get IAP client ID.](https://cloud.google.com/iap/docs/authentication-howto#authenticating_from_a_service_account)

[More detail](https://cloud.google.com/iap/docs/authentication-howto#authenticating_from_a_service_account)

## Use
```
$ GAE_WEB_URI="<YOUR_GAE_URL>" \
CLIENT_ID="<YOUR_IAP_CLIENT_ID>" \
SEVICE_ACCOUNT_KEY_PATH="<YOUR_JSON_KEY_PATH>" \
go run ./main.go
```

for exmaple:
```
$ GAE_WEB_URI="https://<PROJECT_ID>.appspot.com" \
CLIENT_ID="<RANDOM_STRING>.apps.googleusercontent.com" \
SEVICE_ACCOUNT_KEY_PATH="key.json" \
go run ./main.go
```
