# spamtube

### Links

- [Google Drive](https://drive.google.com/drive/folders/1x6eskGhW1XJcFSHURoiRQhkAj4kD99N0)
- [Slack](https://theboardgamesgroup.slack.com/archives/C03PF2S0PFG)
- [Keybase](https://keybase.io/team/spamtube)

## Prerequisites

- [Docker](https://docs.docker.com/get-docker)

## Setup

### Docker

Run project in a docker container:

```bash
docker compose up --build
```

## Development

### Database

Connect to the database with the following command:

```bash
pgcli postgres://postgres:postgres@localhost:5337/spamtube_db
```

> Note: To run pgcli you need to install it from https://www.pgcli.com/install

Run database migrations manually with the following command:

```bash
cd database
go run . -dsn "postgres://postgres:postgres@localhost:5337/spamtube_db?sslmode=disable"
```

> Note: To run migrations manually you need to install golang https://go.dev/doc/install

### Backend

Run the backend manually with the following command:

```bash
cd backend
air
```

> Note: To run backend manually you need to install golang https://go.dev/doc/install

> Note: To run go in `air` (hot-reload mode) you need install air by `go install github.com/cosmtrek/air@latest` or just run `go run .`

### Frontend

Run the frontend manually with the following command:

```bash
cd frontend
npm install
npm run serve
```

> Note: To run frontend manually you need to install nvm https://github.com/nvm-sh/nvm
