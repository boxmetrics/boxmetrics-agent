.PHONY: run test coverage coverage-html prebuild build clean cert rootca

prebuild:
	@echo "Building dependencies"
	mkdir -p bin
	cd bin && go build -v ../internal/app/server.go

run:
	@echo "Starting module"
	make prebuild
	go run main.go

test:
	@echo "Running module test"
	go test -v ./...

coverage:
	@echo "Running coverage"
	go test -v -cover ./...

coverage-html:
	@echo "Running coverage html"
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html

build:
	@echo "Building module to ./bin"
	make prebuild
	cd bin && go build -v ../.

clean:
	@echo "Cleaning workspace"
	rm bin/*
	go mod tidy
	go clean ./...

cert:
	@echo "Generate self-sign certificate"
	@echo "You must have rootCA certificate, if not run 'make rootca'"
	openssl req -new -sha256 -nodes -out certificates/server.csr -newkey rsa:2048 -keyout certificates/server.key -config <( cat configs/server.csr.cnf )
	openssl x509 -req -in certificates/server.csr -CA certificates/rootCA.pem -CAkey certificates/rootCA.key -CAcreateserial -out certificates/server.crt -days 500 -sha256 -extfile configs/v3.ext

rootca:
	@echo "Generate rootCA certificate"
	openssl genrsa -des3 -out certificates/rootCA.key 2048
	openssl req -x509 -new -nodes -key certificates/rootCA.key -sha256 -days 1024 -out certificates/rootCA.pem
	@echo "For Mac OS users :"
	@echo "Open Keychain Access on your Mac and go to the Certificates category in your System keychain."
	@echo "Once there, import the rootCA.pem using File > Import Items."
	@echo "Double click the imported certificate and change the “When using this certificate:” dropdown to Always Trust in the Trust section"
	@echo ""
	@echo "For Windows users :"
	@echo "Follow this tutorial https://www.thewindowsclub.com/manage-trusted-root-certificates-windows"
	@echo "=)"
