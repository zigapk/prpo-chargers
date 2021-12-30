bindata:
	@go-bindata -o internal/config/bindata.go -prefix="configs" -pkg=config configs/...

docs:
	@swag fmt -d cmd/chargers/,internal/handle && swag init -d cmd/chargers/,internal/handle

build:
	@go build github.com/zigapk/prpo-chargers/cmd/chargers