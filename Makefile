.PHONY: start tw

start:
	@echo "Starting the server..."
	@templ generate -watch & echo $$! > templ.pid
	@rm templ.pid
	@go run main.go
tw:
	@tailwindcss -i ./public/css/input.css -o ./public/css/styles.css --watch