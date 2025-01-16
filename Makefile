.PHONY: default build css run

default: build css run

build:
	templ generate

css:
	npx tailwindcss -i ./assets/input.css -o ./assets/dist/output.css --minify

run:
	go run ./cmd/wasessen