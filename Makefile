APP_NAME=gomoji

d:
	@go run main.go

install: build
	@cp dist/$(APP_NAME) $(HOME)/bin/$(APP_NAME)
	@cp data/emojis.json $(HOME)/emojis.json

clean:
	@rm -rf dist
	@rm -rf $(HOME)/bin/$(APP_NAME) $(HOME)/data/emojis.json

build: clean
	@go build -o dist/$(APP_NAME) main.go