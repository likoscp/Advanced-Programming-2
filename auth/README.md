# auth (microservice)

## to deploy

- golang version >= 1.23.0
- bit of luck

download all depencies
```env
go mod tidy
```

compile all files
```cli
make
```

run the project
```env
./a
```

## endpoints

| endpoint   | stand for    |
|------------|--------------|
| /user/login| give jwt   |
|/user/register| give jwt |