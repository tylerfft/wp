SRC = $(wildcard *.go)

export GOPATH=/root/wp:$(shell pwd)

export GOBIN=$(shell pwd)/bin

ws:$(SRC)
	go install ./src/main $^
	mv ./bin/main ./bin/pspsrv
	@echo
	@echo "请确认，当前 gopath 是：" $(GOPATH)
	@echo
