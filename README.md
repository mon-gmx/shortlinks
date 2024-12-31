# shortlinks

This is a backend to help redirection to provide a similar functionality like golinks.

## How to run?

Make sure you have a **Postgres** database beforehand. You will need it for your backend to be able to query what's existing or just to keep record of your short links.

To start the service, just run using `go mod run` or build a binary if you plan on hosting and automating.

If you feel like turning it into a service or just have a binary, run: `go build`

## How to test?

`CONFIG_FILE=$(pwd)/config.yaml go test ./tests/... -v`

This would work as well for execution, the default behavior is searching for `config.yaml` where the service runs.


## Examples:

### To create a new shortlink:
  You can have as many short links to the same URL as you want, there's no validation right now:
  `curl -X POST http://localhost:8080/shorts -H "Content-Type: application/json" -d '{"handle": "goog", "url": "http://www.google.com"}'`
  
  As an alternative, you can use the `/updates` endpoint to load a form in your browser, so you have a simple UI to add new entries.

  
### To be redirected:
  `curl -L http://localhost:8080/goog`

To have the best experience for this redirection project, you would want to have this project serving in port 80 and a record in your DNS so you can just use this in your browser like: `shortlinks/googl` and be redirected to the site associated to the `googl` handle.

If you want to expose this directly, you are advised to use a proxy from your HTTP server so you don't need to add `CAP_NET_BIND_SERVICE` permissions and add security concerns; you can add this directive to your nginx site:

```
	location / {
            proxy_pass http://<hostname>:8000;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
```

### How to add this to my DNS?

Well, you can't in the strict sense, even if you add an SRV entry, the result would be messy. But you can add an entry to a small DHCP/DNS instance like dnsmasq! Or you can simply add the entry to your hosts file. These require minimum effort and are very easy to setup.

```
# dnsmasq.conf entry
address=/go/<HOST RUNNING [PROXY FOR] SHORTLINKS>
address=/.go/<HOST RUNNING [PROXY FOR] SHORTLINKS>

# add your DNS downstream so recursive search happens
server=<YOUR DNS IP>
```
