# Variables
APP_NAME := htmx-server
GO_SRC := ./cmd/server.go
# CSS_SRC := ./web/css
# CSS_BUILD := ./web/dist

# Build the Go server
build:
	pnpm build
	go build -o $(APP_NAME) $(GO_SRC)


tailwind: 
	pnpm dev

air:
	air
