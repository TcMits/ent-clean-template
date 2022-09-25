# Ent Clean template

## Quick start
Local development:
```sh
# Run app with migrations
docker compose -f local.yml up
```

Production:
```sh
docker compose -f production.yml up
```

Integration tests (can be run in CI):
```sh
docker compose -f integration_test.yml up --abort-on-container-exit --build --exit-code-from http_v1_integration
```

Unit tests (can be run in CI):
```sh
go test -cover -race $(go list ./... | grep -v /integration_test/)
```
