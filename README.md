# TestTaskBaverlyHills
test task 

### Create new migration file
```
$ migrate create -ext sql -dir db/migrations -seq {file_name}
```
Example:
```
$ migrate create -ext -dir db/migration -seq sql create_users
```

### Migrate
In order to migrate the tables, run the following command:
```
$ migrate -verbose -source file://db/migration -database "postgres://postgres:password@localhost:5432/beverly?sslmode=disable" up [N]
```

### Rollback
In order to rollback the tables, run the following command:
```
$ migrate -verbose -source file://db/migration -database "postgres://postgres:password@localhost:5432/beverly?sslmode=disable" down [N]
```

N defines the number of migration files to migrate or rollback
