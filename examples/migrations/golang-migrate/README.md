## install

```bash
go install github.com/golang-migrate/migrate@latest
```

## add new migration file

```bash
migrate create -ext sql -dir ./migrations -seq create_users
```

## up

```bash
URL="mysql://root:@tcp(127.0.0.1)/examples-golang-migrate?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
migrate -database $URL -path ./migrations up
```

## down

```bash
URL="mysql://root:@tcp(127.0.0.1)/examples-golang-migrate?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
migrate -database $URL -path ./migrations down
```
