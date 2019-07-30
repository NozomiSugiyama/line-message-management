# Line Message Management
LineのMessage APIを用いて、セグメント配信を行うためのテストリポジトリ

### Deploy
#### Setup
```bash
$ heroku login
$ git remote add heroku-api https://git.heroku.com/line-message-management-api.git
$ git remote add heroku-api https://git.heroku.com/line-message-management-ui.git
$ heroku plugins:install heroku-config
```

#### Deploy API
```bash
$ heroku config:push -f api/.env --app line-message-management-api # if modified .env file
$ git subtree push --prefix api/ heroku-api master
```
#### Deploy UI
```bash
$ git subtree push --prefix ui/ heroku-ui master
```
