
run:
	go run ./cmd/app 

run_prod:
	go run -tags prod ./cmd/app

build_with_cgo:
	go build --tags prod -o app ./cmd/app

build:
	CGO_ENABLED=0 go build --tags prod -o app ./cmd/app
