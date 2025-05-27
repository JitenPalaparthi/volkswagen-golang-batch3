- create docker network

```sh
docker network create demo-network
```

- create postgres container

```sh
docker run -d --name pgdb -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=usersdb --network=demo-network postgres 
```

- create adminer UI for db

```sh
docker run -d --name dbui -p 38080:8080 --network demo-network adminer
```

- test a func

```sh
go test -timeout 30s -run ^TestUserValidate$ demo/models
```

- test whole package

```sh
go test -timeout 30s -coverprofile=coverage.out ./models/
```

- html coverage
```sh
go tool cover -html=coverage.out -o coverage.html
open coverate.html
```

- directly open in webbrowser

```sh

  go tool cover -html=coverage.out
  ```

  - assertion library 

  go get github.com/stretchr/testify/assert