# Sample Go api server

## Tech stack

| STT | Type                         | Stack Name              | Support Reason                                                                                                 | Version                    | Source                                     |
|-----|------------------------------|-------------------------|----------------------------------------------------------------------------------------------------------------|----------------------------|--------------------------------------------|
| 1   | Programming Language         | Golang                  |                                                                                                                | 1.21                       |                                            |
| 2   | REST API Framework           | Gin                     | We'll need to serve REST API for Admin Web. Gin is the framework that have the best performance in the market. | v1.4.0                     | https://github.com/gin-gonic/gin           |
| 3   | Input Validator              | go-playground/validator | Validate input from Admin Web                                                                                  | v10                        | https://github.com/go-playground/validator |
| 4   | GRPC library                 | grpc/grpc-go            | Support gRPC implementation for service to service call                                                        | v1.58.1                    | https://github.com/grpc/grpc-go            |
| 5   | Thrift library               | apache/thrift           | Support Thrift implementation for service to service call                                                      | v0.19.0                    | https://thrift.apache.org/tutorial/go.html |
| 6   | Dependency Injection library | google/wire             | Since gin not have built-in DI mechanism. We need a DI mechanism to maintain dependencies in the process.      |                            | https://github.com/google/wire             |
| 7   | Unit Testing                 | ginkgo                  | Testing library provide by Gin                                                                                 | 2.12.1                     | https://github.com/onsi/ginkgo             |
| 8   | Unit Testing                 | gomega                  | matcher/assertion library                                                                                      | 1.27.0                     | https://github.com/onsi/gomega             |
| 9   | Code Formatter               | go fmt                  | Common code formatter tools provide by go                                                                      | compatible with go version | https://pkg.go.dev/fmt                     |
| 10  | Logging Library              | log/slog                | A built-in logging library in golang that support structured log                                               | compatible with go version | https://pkg.go.dev/golang.org/x/exp/slog   |
