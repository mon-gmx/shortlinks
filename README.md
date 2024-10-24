# shortlinks

This is a backend to help redirection to provide a similar functionality like golinks.

## How to run?

Make sure you have a **Postgres** database beforehand. You will need it for your backend to be able to query what's existing or just to keep record of your short links.

To start the service, just run using `go mod run` or build a binary if you plan on hosting and automating.

## How to test?

`CONFIG_FILE=$(pwd)/config.yaml go test ./tests/... -v`


## Examples:

### To create a new shortlink:
  You can have as many short links to the same URL as you want, there's no validation right now
  `curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"handle": "goog", "url": "http://www.google.com"}'`
  
### To be redirected:
  `curl -L http://localhost:8080/goog`


To have the best experience for this redirection project, you would want to have this project serving in port 80 and a record in your DNS so you can just use this in your browser like: `shortlinks/googl` and be redirected to the site associated to the `googl` handler.
