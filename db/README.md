# LINE Message management db

Using PostgreSQL

### Environments

### Setup
```bash
$ cp .env.development.sample .env.development
$ vim .env # Edit .env file
$ go build
```

### Debug
```bash
$ docker run -it --rm --network line-message-management_default postgres psql -h db -U postgres -p 5432
```
