.PHONY: default build css run

default: build css run

build:
	templ generate

css:
	npx tailwindcss -i ./internal/assets/input.css -o ./internal/assets/dist/output.css --minify

run:
	go run ./cmd/wasessen

zip:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./infra/main ./cmd/wasessen
	zip -j ./infra/wasessen.zip ./infra/main
	rm ./infra/main

deploy: terraform-plan terraform-apply