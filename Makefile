# Configurations

_MICROSERVICES = client-api-ms port-domain-ms

# Testing all services
go-build:


proto:
	for service in $(_MICROSERVICES) ; do \
		protoc -I/usr/local/include \
				-I. \
				-I$(GOPATH)/src \
				--go_out=plugins=grpc:. \
				$$service/pb/$$service/$$service.proto; \
	done

test:
	go test ./...

check-code:
	go vet ./...
