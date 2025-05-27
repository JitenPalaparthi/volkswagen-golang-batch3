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
