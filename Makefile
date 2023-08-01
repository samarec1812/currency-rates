# EXAMPLE OPTIONS
EG_DATE=2022-10-08
EG_CODE=USD

.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run ./... -c ./.golangci.yml


.PHONY: test
test:
	$(info Testing project...)
	go test -v ./...




.PHONY: build
build:
	$(info Build project...)
	go mod tidy
	go build -o bin/currency-rates cmd/main.go


.PHONY: run
run:
	$(info Run project...)
	./bin/currency-rates --date=$(EG_DATE) --code=$(EG_CODE)