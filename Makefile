default: signalmon

.PHONY: signalmon
signalmon:
	@go build
	@go test

.PHONY: test
test:
	@go test

.PHONY: clean
clean:
	@rm -f signalmon
