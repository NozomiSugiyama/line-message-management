# LINE Message management api

### Environments

```bash
$ go version
go version go1.12.1 darwin/amd64
```

### Setup
```bash
$ cp .env.production.sample .env.production # dev mode
$ cp .env.development.sample .env.development # prod mode
$ vim .env # Edit .env file
```

### Debug

```bash
$ go build
$ ./api dev
```

#### docker compose
```bash
$ docker-compose up
$ go build
$ docker-compose restart api
```
