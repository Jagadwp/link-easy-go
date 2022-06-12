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
