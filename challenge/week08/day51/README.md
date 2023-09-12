# Day 51 of #66DaysOfGo

_Last update:  Sep 11, 2023_.

---

Today, I continued working on the basic skeleton within a Clean Architecture series. This time replacing the chi library and testing the mongo driver.

---

> Based on _https://github.com/tensor-programming/hex-microservice/tree/master_

Check the code in [https://github.com/jp-chl/test-go-clean-architecture/tree/v10](https://github.com/jp-chl/test-go-clean-architecture/tree/v10)

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

* After cloning the v10 branch of the [repo](https://github.com/jp-chl/test-go-clean-architecture.git), set environment variables and run the code.

```bash
export MONGO_TIMEOUT=30
export MONGO_URL=mongodb://localhost:27017/mydb
export MONGO_DB=mydb
```

```bash
go run main.go
```

---

## References

- [https://github.com/tensor-programming/hex-microservice/tree/master](https://github.com/tensor-programming/hex-microservice/tree/master)
- [https://github.com/jp-chl/test-go-clean-architecture/tree/v10](https://github.com/jp-chl/test-go-clean-architecture/tree/v10)
