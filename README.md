# Example Traefik setup using Docker Compose

Should just be able to `docker compose up --build` and it will start working.

Uses the `traefik.yml` as a file provider for routing rules.
You can get to either the `whoami` or `testy` containers via either 
`http://whoami.localhost` or `http://test.localhost` respectively, or by using 
the router rules to access them via `http://me.localhost/whoami` or `http://me.localhost/testy`.




