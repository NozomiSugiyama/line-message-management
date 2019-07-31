# Line Message Management
LineのMessage APIを用いて、セグメント配信を行うためのテストリポジトリ

### Setup
1. Execute API setup based on api/README.md
2. Execute DB setup based on db/README.md
3. Execute UI setup based on ui/README.md

### Local Debug
```bash
$ docker-compose pull
$ docker-compose up
```

### Deploy

```bash
$ heroku login
$ git remote add heroku-api https://git.heroku.com/line-message-management-api.git
$ git remote add heroku-api https://git.heroku.com/line-message-management-ui.git
$ heroku plugins:install heroku-config
```

#### Deploy API
```bash
$ heroku config:push -f api/.env.production --app line-message-management-api # if modified .env file
$ git subtree push --prefix api/ heroku-api master
```
#### Deploy UI
```bash
$ git subtree push --prefix ui/ heroku-ui master
```

### Development
#### Modified postgres environment
```bash
$ docker-compose stop
$ rm -rf ./db/data
```
