# Ecommerce - Product Service

## Tech Stack
- [Golang](https://github.com/golang/go)
- [MySQL (Database)](https://github.com/mysql/mysql-server)

## Framework
- [GoFiber (HTTP Framework)](https://github.com/gofiber/fiber)
- [Viper (Configuration)](https://github.com/spf13/viper)
- [Go Playground Validator (Validation)](https://github.com/go-playground/validator)
- [Golang Migration](https://github.com/golang-migrate/migrate)

## Development Tools
- [air (hot reload)](https://github.com/cosmtrek/air)

## How to install
```bash
go get .
```

## Run with air (hot reload)
```bash
air
```

## Create Migration
```bash
migrate create -ext sql -dir db/migrations create_table_xxx
```

## Run Migration
```bash
migrate -database "mysql://root:password@tcp(localhost:3306)/ecommerce_product_service?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

## Fix Database Dirty
```bash
migrate -database "mysql://root:password@tcp(localhost:3306)/ecommerce_product_service?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations force VERSION
```