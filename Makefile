default: signalmon

.PHONY: signalmon
signalmon:
	@go build

.PHONY: test
test:
	@go test

.PHONY: clean
clean:
	@rm -f signalmon
