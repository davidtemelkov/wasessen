.PHONY: default build css run

default: build css run

build:
	templ generate

css:
	npx tailwindcss -i ./internal/assets/input.css -o ./internal/assets/dist/output.css --minify

run:
	go run ./cmd/wasessen