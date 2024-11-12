# Example Traefik setup using Docker Compose

### Relevant files

`config/traefik.yaml`

Main configuration file of Traefik

`config/rules.yaml`

Dynamic configuration for Traefik such as URL-based routing.


`.env`

You'll also need a `.env` with your cloudflare API token for TLS:

```
CF_DNS_API_TOKEN = 'token here'
```

### Notes

Make sure to `touch data/certs/acme.json && chmod 600 data/certs/acme.json`.

