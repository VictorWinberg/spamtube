# spamtube

### Links
- [Google Drive](https://drive.google.com/drive/folders/1x6eskGhW1XJcFSHURoiRQhkAj4kD99N0)
- [Slack](https://theboardgamesgroup.slack.com/archives/C03PF2S0PFG)
- [Keybase](https://keybase.io/team/spamtube)

## Setup

### Docker
Initiate a new docker container:

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
go run main.go -dsn "postgres://postgres:postgres@localhost:5337/spamtube_db?sslmode=disable"
```
