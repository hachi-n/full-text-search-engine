APP_NAME := ftse

go-clean:
	rm -f ./bin/$(APP_NAME)

go-build:
	go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)

build:
	$(MAKE) go-clean
	$(MAKE) go-build

clean:
	$(MAKE) go-clean
