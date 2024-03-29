HOSTNAME=registry.terraform.io
NAMESPACE=blameless
NAME=blameless
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_arm64

default: install

build:
	go build -o ${BINARY}

install: build
	rm -f -r ./modules/.terraform
	rm -f ./modules/.terraform.lock.hcl
	rm -f ./modules/*.tfstate*        
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -bench=. -cover ./...                                                             

lint:
	golangci-lint run
	tflint --chdir=./modules

doc:
	go generate ./...
