default: signalmon

.PHONY: signalmon
signalmon: templates/index.go
	@go build
	@go test

templates/index.go: templates/index.html

.PHONY: assets
assets: assets/bundle.js

assets/bundle.js: assets/app.js assets/app.css
	@cd assets && npm run build

.PHONY: test
test:
	@go test

.PHONY: clean_assets
clean_assets:
	@cd assets && npm run clean

.PHONY: clean
clean:
	@rm -f signalmon
