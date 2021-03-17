.PHONY: gomodgen deploy delete build run clean

run: build-cli
	dist/azureog -title "$(TITLE)"

build-cli:
	go build \
		-trimpath \
		-ldflags "-s -w -extldflags '-static'" \
		-o dist/azureog \
		./cmd/azureog && \
		chmod +x ./dist/azureog

build-local-func:
	go build \
		-trimpath \
		-ldflags "-s -w -extldflags '-static'" \
		-o handler \
		./cmd/handler && \
		chmod +x ./handler

build-deploy-func:
	GOOS=linux GOARCH=amd64 go build \
		-trimpath \
		-ldflags "-s -w -extldflags '-static'" \
		-o handler \
		./cmd/handler && \
		chmod +x ./handler



run-func: build-local-func
	func start

clean:
	rm -rf ./dist && rm og-standard.png && rm handler

