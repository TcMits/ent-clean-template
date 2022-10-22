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

Generate locale files:
```sh
goi18n extract -sourceLanguage=en-US -outdir=./locales/en-US/ -format=yaml ./
```

Generate docs files:
```sh
swag init -g ./internal/controller/http/v1/v1.go -o ./docs/v2
```

Convert OpenApi v2 to v3 (yaml):
```sh
# https://github.com/swaggo/swag/issues/386
docker run --rm -v $(PWD)/docs:/work openapitools/openapi-generator-cli:latest-release \
	  generate -i /work/v2/swagger.yaml -o /work/v3 -g openapi-yaml --minimal-update \
    && mv ./docs/v3/openapi/openapi.yaml ./docs/v3/ \
    && rm ./docs/v3/README.md \
    && rm ./docs/v3/.openapi-generator-ignore \
    && rm -rf ./docs/v3/.openapi-generator \
    && rm -rf ./docs/v3/openapi
```

Convert OpenApi v2 to v3 (json):
```sh
# https://github.com/swaggo/swag/issues/386
docker run --rm -v $(PWD)/docs:/work openapitools/openapi-generator-cli:latest-release \
	  generate -i /work/v2/swagger.json -o /work/v3 -g openapi --minimal-update \
    && rm ./docs/v3/README.md \
    && rm ./docs/v3/.openapi-generator-ignore \
    && rm -rf ./docs/v3/.openapi-generator 
```

### Changing between openapi v2/v3:
Change [docs.go](https://github.com/TcMits/ent-clean-template/blob/master/docs/docs.go)

```go
package docs

import _ "github.com/TcMits/ent-clean-template/docs/v3" // replace this
// import _ "github.com/TcMits/ent-clean-template/docs/v2"
```


## Overview

### Web framework
[Iris](https://www.iris-go.com/) is an efficient and well-designed, cross-platform, web framework with robust set of features. Build your own high-performance web applications and APIs powered by unlimited potentials and portability.

### Database - ORM
[ent](https://entgo.io/docs/getting-started/) is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:

-   Easily model database schema as a graph structure.
-   Define schema as a programmatic Go code.
-   Static typing based on code generation.
-   Database queries and graph traversals are easy to write.
-   Simple to extend and customize using Go templates.


### File system
[go-storage](https://github.com/beyondstorage/go-storage) is a **vendor-neutral** storage library for Go.

### Swagger urls

```
/api/v1/swagger
```

### Default urls

----
Login endpoint

* **URL**

  `/api/v1/login`

* **Method:**

  `POST`

* **Data Params**

```json
{
  "username": "string",
  "password": "string"
}
```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
```json
{
  "access_token": "string",
  "refresh_token": "string",
  "refresh_key": "string"
}
```

* **Error Response:**

  * **Code:** `400` | `500` | `401` <br />
    **Content:** 
```json
{
  "message": "string",
  "code": "string"
}
```

----
Refresh token endpoint

* **URL**

  `/api/v1/refresh-token`

* **Method:**

  `POST`

* **Data Params**

```json
{
  "refresh_token": "string",
  "refresh_key": "string"
}
```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
```json
{
  "token": "string"
}
```

* **Error Response:**

  * **Code:** `400` | `500` | `401` <br />
    **Content:** 
```json
{
  "message": "string",
  "code": "string"
}
```

----
Verify token endpoint

* **URL**

  `/api/v1/verify-token`

* **Method:**

  `POST`

* **Data Params**

```json
{
  "token": "string"
}
```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
```json
{}
```

* **Error Response:**

  * **Code:** `400` | `500` | `401` <br />
    **Content:** 
```json
{
  "message": "string",
  "code": "string"
}
```

----
Me endpoint

* **URL**

  `/api/v1/me`

* **Method:**

  `GET` | `PUT` | `PATCH`

* **Data Params**

```json
{
  "username": "string",
  "email": "string",
  "first_name": "string",
  "last_name": "string"
}
```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
```json
{
  "self": "string",
  "id": "string",
  "create_time": "string",
  "update_time": "string",
  "username": "string",
  "first_name": "string",
  "last_name": "string",
  "email": "string",
  "is_staff": "string",
  "is_superuser": "string",
  "is_active": "string"
}
```

* **Error Response:**

  * **Code:** `400` | `500` | `401` <br />
    **Content:** 
```json
{
  "message": "string",
  "code": "string"
}
```

