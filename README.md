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

#### Modified postgres environment
```bash
$ docker-compose stop
$ rm -rf ./db/data
```

### Deploy

#### Setup
```bash
$ heroku login
$ git remote add heroku-api https://git.heroku.com/line-message-management-api.git
$ git remote add heroku-ui https://git.heroku.com/line-message-management-ui.git
$ heroku plugins:install heroku-config
```

#### Setup DB
```bash
$ heroku pg:psql -a line-message-management-api -f ./db/sql/ddl.sql
$ heroku pg:psql -a line-message-management-api -f ./db/sql/dml.sql
```

#### Deploy API
```bash
$ heroku config:push -f api/.env.production --app line-message-management-api -o # If modified .env file
$ git subtree push --prefix api/ heroku-api master
$ heroku run --app line-message-management-api ./bin/api dev # In dev mode
```

#### Deploy UI
```bash
$ heroku config:push -f ui/.env.production --app line-message-management-ui -o # If modified .env file
$ git subtree push --prefix ui/ heroku-ui master
```

#### Debug
##### API
```bash
$ heroku logs --tail --app line-message-management-api
```

##### UI
```bash
$ heroku logs --tail --app line-message-management-ui
```

##### DB
```bash
$ heroku pg:info -a line-message-management-api
$ heroku pg:psql -a line-message-management-api
```

<!--

```bash
$ git push heroku-api `git subtree split --prefix api/ master`:master --force
$ git push heroku-ui `git subtree split --prefix ui/ master`:master --force
```

-->
