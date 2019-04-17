.PHONY run, test, build, cert, rootca

run:
	go run ./...

test:
	go test ./...

build:
	go build ./...

cert:
	echo "Generate self-sign certificate"
	echo "You must have rootCA certificate, if not run 'make rootca'"
	openssl req -new -sha256 -nodes -out server.csr -newkey rsa:2048 -keyout server.key -config <( cat configs/server.csr.cnf )
	openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500 -sha256 -extfile configs/v3.ext

rootca:
	echo "Generate rootCA certificate"
	openssl genrsa -des3 -out rootCA.key 2048
	openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.pem
	echo "For Mac OS users :"
	echo "Open Keychain Access on your Mac and go to the Certificates category in your System keychain." echo "Once there, import the rootCA.pem using File > Import Items."
	echo "Double click the imported certificate and change the “When using this certificate:” dropdown to Always Trust in the Trust section"
	echo ""
	echo "For Windows users :"
	echo "Follow this tutorial https://www.thewindowsclub.com/manage-trusted-root-certificates-windows"
	echo "=)"