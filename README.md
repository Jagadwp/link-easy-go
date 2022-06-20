# Link Easy
## About
This is an url shortener API created by Echo (Go Framework) and PostgreSQL.
## Features
- Authentication
- Middleware
- Use clean architecture
![https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png]()
- CRUD user
- CRUD url
- Caching with Redis
## How to Run

Before you start the program, don't forget to do this on your terminal

```
go mod tidy
```

please make your .env on the root folder, you can use this example by changing on your preferences

```
DB_HOST=<string>
DB_DRIVER=<string>
DB_USER=<string>
DB_PASSWORD=<string>
DB_NAME=<string>
DB_PORT=<string> 

TestDbHost=<string>
TestDbDriver=<string>
TestDbUser=<string>
TestDbPassword=<string>
TestDbName=<string>
TestDbPort=<string>
```

Run app
```
go run app/main.go
```

## API Documentation