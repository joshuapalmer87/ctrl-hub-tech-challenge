# How to use

Utilises asdf for go versioning management (only for Linux/Mac).
To check you are running the correct version of go run:
```asdf install golang```

To spin up the server run: 
```go build ./... && go run . &```

# Implementation Plan
## Initial plan
- [x] JP-1: Create implementation plan
- [x] JP-2: Implement initial basic http server with ping endpoint
- [x] JP-3: Add initial test for ping endpoint
- [x] JP-4: Add post endpoint
- [ ] JP-5: Add get all endpoint
- [ ] JP-6: Add get single endpoint
- [ ] JP-7: Add get summary endpoint
- [ ] JP-8: Ensure tests are passing

## Stretch goal
- [ ] JP-9: Move to real backing storage e.g. mongo/postgres

## Further work
- [ ] JP-10: Dockerise the deployment
- [ ] JP-11: Set up with CI/CD environment (test building, security)
- [ ] JP-12: Add some better linting
- [ ] JP-13: Add error monitoring and stats
- [ ] JP-14: User validation and access controls