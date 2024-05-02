bootstrap:
	go mod download
	cp config.example.yaml config.yaml
	touch /tmp/sgs.pid
run:
	@go run main.go serve
config:
	@go run main.go config
test:
	@go test -race -json ./...
forward-port:
	@./scripts/forward-port.sh
total-cover:
	@go test -race ./... -coverprofile=coverage.out > /dev/null && go tool cover -func coverage.out | grep total


# Setup hot reload. Ref: https://medium.com/@olebedev/live-code-reloading-for-golang-web-projects-in-19-lines-8b2e8777b1ea
PID      = /tmp/sgs.pid
GO_FILES = $(shell find . -name "*.go" | grep -v "_test.go" | xargs)
APP      = ./build/sgs
dev: restart
	@fswatch -o ./internal -o config.yaml | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@clear
	@echo "File change"
$(APP): $(GO_FILES)
	@echo "Rebuilding..."
	@go build -o $(APP) --gcflags "all=-N -l" .
restart: kill before $(APP)
	@$(APP) serve & echo $$! > $(PID)

.PHONY: serve restart kill before
