.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -v && mv BillPage publish/BillPage && chmod +x publish/BillPage && cp static/* publish/static/ && cp templates/* publish/templates/
.PHONY: run
run:
	go run main.go
.PHONY: clean
clean:
	rm cmd/watchdog/watchdog