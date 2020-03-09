BIN_NAME=envssm

.PHONY: all
all: ${BIN_NAME}

.PHONY: ${BIN_NAME}
${BIN_NAME}:
	go build -o ./bin/${BIN_NAME} .


.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race
