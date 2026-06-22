# How to use

Utilises asdf for go versioning management (only for Linux/Mac).
To check you are running the correct version of go run:
```asdf install golang```

To spin up the server run: 
```go build ./... && go run main.go```

# Implementation Plan
## Initial plan
- [x] Create implementation plan
- [x] Implement initial basic http server with ping endpoint
- [ ] Add initial test for ping endpoint
- [ ] Add failing tests for endpoints in spec.yaml
- [ ] Add internal model structs consistent with spec.yaml
- [ ] Add additional endpoints in line with spec, including validation and error returns
- [ ] Add persistence with internal memory storage (not perfect)
- [ ] Ensure tests are passing

## Stretch goal
- [ ] Move to real backing storage e.g. mongo/postgres

## Further work
- [ ] Dockerise the deployment
- [ ] Set up with CI/CD environment (test building, security)