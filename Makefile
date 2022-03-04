build:
	@go build -o ./dist/kabarda ./cmd/cli

## build_cli: builds the command line tool kabarda and copies it to myapp
build_cli:
	@go build -o ../myapp/kabarda ./cmd/cli