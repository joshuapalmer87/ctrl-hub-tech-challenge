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
- [x] JP-5: Add get all endpoint
- [x] JP-6: Add get single endpoint
- [x] JP-7: Add get summary endpoint

## Stretch goal
- [ ] Supplement testing with more paths tested, particularly error paths
- [ ] JP-9: Move to real backing storage e.g. mongo/postgres
- [ ] Return more meaningful responses for failed validation, etc

## Further work
- [ ] JP-10: Dockerise the deployment
- [ ] JP-11: Set up with CI/CD environment (test building, security)
- [ ] JP-12: Add some better linting
- [ ] JP-13: Add error monitoring and stats
- [ ] JP-14: User validation and access controls
- [ ] Ensure debugging support
- [ ] Investigate whether there are any official inputs and outputs to the calculations to ensure alignment
- [ ] Optimise to use bucketed times for summary storage to improve retrieval rates

# Event Driven Discussion
Realistically, this example is simple enough that the calculations won't take much time to perform when adding data to making the process asynchronous wouldn't be worth the payoff.
If we were to assume that the processing became more intensive then we would be able to add a queue on the items added via post, and ensure they were processed.
As these results are required for compliance, we need to recognise if any errors are retryable later (connection issue or similar) or will not be retryable (user wasn't found)
In the case of the latter, we would be obliged to let the users know that the request had failed, as they would otherwise have not visibility.
We might need to invent a polling driven service to allow users to ensure they know the state of their requests.