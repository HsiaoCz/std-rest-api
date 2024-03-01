# Hotel reservation backend

## Resources

### MongoDB dirver

```bash
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client

```go
go get go.mongodb.org/mongo-driver/mongo
```

### go-fiber

Documentation

```bash
https://gofiber.io
```

Installing go-fiber

```go
go get github.com/gofiber/fiber/v2
```

### Docker

Installing mongodb as a Docker container

```docker
docker run --name mongodb -d mongo:latest -p 27017:27017
```

### Design

- users ----> book room from an hotel
- admins ----> going to check reservation/bookings
- Authentication and autherization ---> jwt tokens
- Hotels ----> crud API ---> json
- Rooms -----> crud API -----> json
- Scripts ----> database management ---> seeding, migration
