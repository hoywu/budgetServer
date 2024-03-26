# budgetServer

## Build

Binary: see [GitHub Actions](https://github.com/hoywu/budgetServer/actions)

Docker: Use `--network=host` for docker build if you are behind a proxy.

`docker build --network=host --tag budget .`

## Run

If you have your own MySQL environment, copy `config/config.yaml` and modify it.

- `docker run -d --name budget --network host -v /PATH/TO/YOUR/config.yaml:/app/config.yaml budget`

Or you can use docker compose to run both MariaDB and budget server. Copy `docker-compose.yml.example` to `docker-compose.yml` and modify it.

- `docker compose up -d`

## Web Frontend

see https://github.com/hoywu/budget-web

![desktop-shot](https://raw.githubusercontent.com/hoywu/budget-web/main/assets/desktop_shot.png)
