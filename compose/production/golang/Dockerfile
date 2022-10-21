# Step 1: Modules caching
FROM golang:1.19-bullseye as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.19-bullseye as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

# createuser cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/createuser ./cmd/createuser

# migrate cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/migrate ./cmd/migrate

# app cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app .

# step3: copy from builder
FROM golang:1.19-bullseye 

RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=builder /bin/app /app
COPY --from=builder /bin/migrate /migrate
COPY --from=builder /bin/createuser /createuser
COPY --from=builder /app/config /config
COPY --from=builder /app/locales /locales

CMD ["/app"]
