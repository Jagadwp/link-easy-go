You must make a migration
```
go run db/migration/main.go   
```


Before you start the program, don't forget to do this on your terminal
```
go mod tidy
```


please make your .env on the root folder, you can use this example by changing on your preferences
```
local_host_port="localhost:8080"
db_open_path="/Users/radityarestan/Projects/mini-project/db/mini-project.db"
```


if u want to debug, feel free to use debug mode on your VSCODE, i have adding .vscode in this repo
or u can just run 
```
go run cmd/main.go  
``` 