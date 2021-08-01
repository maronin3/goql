# GOQL - Golang with GraphQL
- go-gin
- graphql-go
- mongodb
- nginx
- docker, docker-compose

## How to Quick Start
- go mod tidy
- go run main.go

### Example :
```
1. Run Playground
http://localhost:3000/graphql

2. Sample Test

Input:
{
  SampleQuery(ID:"test",Password:"test") {
 Message
  }
}

Output:
{
  "data": {
    "SampleQuery": {
      "Message": "Hello, This is SampleQuery"
    }
  }
}

```
## Script
```bash
1. run_compose.sh
2. remove_dbdata.sh
3. redeploy.sh
4. log.sh
```

## Dockerfile

## Docker-Compose

## Nginx


