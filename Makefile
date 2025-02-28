
run:
	DOCKER_API_VERSION=1.47 go run ./cmd/app
	# go run ./cmd/app --registry https://localhost:5000 

run_prod:
	go run -tags prod ./cmd/app
	# go run -tags prod ./cmd/app --registry https://localhost:5000

build_with_cgo:
	go build --tags prod -o app ./cmd/app

build:
	CGO_ENABLED=0 go build --tags prod -o wdu ./cmd/app
