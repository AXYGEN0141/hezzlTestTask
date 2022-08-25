![Go Report](https://goreportcard.com/badge/github.com/AXYGEN0141/hezzlTestTask)
![Repository Top Language](https://img.shields.io/github/languages/top/AXYGEN0141/hezzlTestTask)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AXYGEN0141/hezzlTestTask)
![Github Repository Size](https://img.shields.io/github/repo-size/AXYGEN0141/hezzlTestTask)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/AXYGEN0141/hezzlTestTask)
![GitHub contributors](https://img.shields.io/github/contributors/AXYGEN0141/hezzlTestTask)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

<img align="right" width="50%" src="./images/big-gopher.jpg">

# gRPC Hezzl Test Task

## Task description

1. Describe the **.proto** file with the service using 3 methods: to **ADD** user, to **DELETE** user, to **LIST** users
2. Implement an rpc service based on a proto file on Go
3. Use PostgreSQL for data storage
4. Upon request for a list of users, the data will be cached in redis for a minute and taken from radish
5. When adding a user, make a log in ClickHouse
6. Add logs to ClickHouse via Kafka queue