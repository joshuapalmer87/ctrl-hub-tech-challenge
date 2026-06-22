# How to use

Utilises asdf for go versioning management (only for Linux/Mac).
To check you are running the correct version of go run:
```asdf install golang```

To spin up the server run: 
```go build ./... && go run main.go```

# Implementation Plan
## Initial plan
- [x] JP-1: Create implementation plan
- [x] JP-2: Implement initial basic http server with ping endpoint
- [x] JP-3: Add initial test for ping endpoint
- [ ] JP-4: Add failing tests for endpoints in spec.yaml
- [ ] JP-5: Add internal model structs consistent with spec.yaml
- [ ] JP-6: Add additional endpoints in line with spec, including validation and error returns
- [ ] JP-7: Add persistence with internal memory storage (not perfect)
- [ ] JP-8: Ensure tests are passing

## Stretch goal
- [ ] JP-9: Move to real backing storage e.g. mongo/postgres

## Further work
- [ ] JP-10: Dockerise the deployment
- [ ] JP-11: Set up with CI/CD environment (test building, security)
- [ ] JP-12: Add some better linting