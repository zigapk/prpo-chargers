FROM golang:1.17-bullseye as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Get dependencies
RUN go mod download -x

# Build
COPY cmd cmd
COPY internal internal
RUN CGO_ENABLED=0 GOOS=linux go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOOS=linux swag init -d cmd/chargers/,internal/handle --parseDependency --parseInternal --parseDepth 1 -g main.go
RUN CGO_ENABLED=0 GOOS=linux go build github.com/zigapk/prpo-chargers/cmd/chargers

#####################
# Running container #
#####################
FROM debian:bullseye

# Get certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR /app

COPY configs configs
COPY --from=builder /app/chargers .

ENTRYPOINT ["./chargers"]
CMD ["serve"]