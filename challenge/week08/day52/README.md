# Day 52 of #66DaysOfGo

_Last update:  Sep 12, 2023_.

---

Today, I continued working on the basic skeleton within a Clean Architecture series. This time exposing a POST method to save a document into a Mongo DB.

---

> Based on _https://github.com/tensor-programming/hex-microservice/tree/master_

Check the code in [https://github.com/jp-chl/test-go-clean-architecture/tree/v11](https://github.com/jp-chl/test-go-clean-architecture/tree/v11)

---

Setup

* Use docker

```bash
docker run -d \
  --name my-mongo-container \
  -e MONGO_INITDB_DATABASE=mydb \
  -p 27017:27017 \
  mongo
```

* After cloning the v11 branch of the [repo](https://github.com/jp-chl/test-go-clean-architecture.git), set environment variables and run the code.

```bash
export DB=mongo
export MONGO_TIMEOUT=30
export MONGO_URL=mongodb://localhost:27017/mydb
export MONGO_DB=mydb
```

* Run the server

```bash
$ go run main.go
Connected...
Ping ok...
```

* Hit the endpoint

```bash
curl -s --location 'http://localhost:8000' --header 'Content-Type: application/json' --data '{"url": "https://www.amazon.com"}' | jq .
```

```json
{
  "Code": "2SvSwjAJ",
  "URL": "https://www.amazon.com",
  "CreatedAt": 1694569116
}
```

* Query the DB

Run mongo client

```bash
$ docker run -it --rm --link my-mongo-container:mongo mongo mongo --host mongo --authenticationDatabase admin
MongoDB shell version v4.4.5
connecting to: mongodb://mongo:27017/?authSource=admin&compressors=disabled&gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("3e1f8ced-7e02-427a-929f-df48378e4539") }
MongoDB server version: 4.4.5
Welcome to the MongoDB shell.
# ...

> use mydb
switched to db mydb

> db.temp.findOne({})
{
  "_id" : ObjectId("65010ffd80fae124c72563e0"),
  "code" : "2SvSwjAJ",
  "url" : "https://www.amazon.com",
  "createdat" : NumberLong(1694569116)
}
```

---

## References

- [https://github.com/tensor-programming/hex-microservice/tree/master](https://github.com/tensor-programming/hex-microservice/tree/master)
- [https://github.com/jp-chl/test-go-clean-architecture/tree/v11](https://github.com/jp-chl/test-go-clean-architecture/tree/v11)
