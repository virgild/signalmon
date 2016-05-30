default: signalmon

.PHONY: signalmon
signalmon: templates/index.go
	@go build
	@go test

templates/index.go: templates/index.html

.PHONY: assets
assets: assets/bundle.js

assets/bundle.js: assets/app.js assets/app.css
	@cd assets && browserify -t [ babelify ] -t browserify-css app.js -o bundle.js

.PHONY: test
test:
	@go test

.PHONY: clean_assets
clean_assets:
	@rm -f assets/bundle.js

.PHONY: clean
clean:
	@rm -f signalmon
