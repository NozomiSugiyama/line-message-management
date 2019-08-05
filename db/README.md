# LINE Message management db

Using PostgreSQL

### Initial Users
| ID | Name | Email | Password |
| -- | ---- | ----- | -------- |
| 1 | sample | user@example.com | password |
| 2 | sample2 | user2@example.com | password |
| 3 | sample3 | user3@example.com | password |
| 4 | sample4 | user4@example.com | password |
| 5 | sample5 | user5@example.com | password |

### Setup
```bash
$ cp .env.development.sample .env.development
$ vim .env # Edit .env file
$ docker-compose up db
$ docker-compose exec db psql -h db -U postgres -p 5432 -d line_message_management -f /sql/ddl.sql
$ docker-compose exec db psql -h db -U postgres -p 5432 -d line_message_management -f /sql/dml.sql
```

### Restore
```bash
$ docker-compose up -d db
$ docker-compose exec db psql -h db -U postgres -p 5432 -f /sql/drop.sql
$ docker-compose exec db psql -h db -U postgres -p 5432
> CREATE DATABASE line_message_management;
$ docker-compose exec db psql -h db -U postgres -p 5432 -d line_message_management -f /sql/ddl.sql
$ docker-compose exec db psql -h db -U postgres -p 5432 -d line_message_management -f /sql/dml.sql
```

### Debug
```bash
$ docker-compose up db
```

#### CLI
```bash
$ docker run -it --rm --network line-message-management_default postgres psql -h db -U postgres -p 5432 -d line_message_management
```

Document: https://www.postgresql.org/docs/11/app-psql.html

#### Web
##### pgweb
1. Open: http://localhost:8081/
2. Select Schema tab
3. Insert `postgres://postgres:password@db:5432/line_message_management?sslmode=disable`
4. Click Connect button

##### adminer
1. Open: http://localhost:8082/
3. Insert credentials
4. Click Login button

### Export SQL
#### DDL
```bash
$ docker run -it --rm --network line-message-management_default postgres pg_dump -s --dbname=postgresql://postgres:password@db:5432/line_message_management > sql/ddl.sql
```

#### DML
```bash
$ docker run -it --rm --network line-message-management_default postgres pg_dump -a --dbname=postgresql://postgres:password@db:5432/line_message_management > sql/dml.sql
```
