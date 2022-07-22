CFLAG=-std=c11 -g -static
9cc:
	go build -o 9cc main.go

.PHONY: deps test clean
deps:
	go mod tidy
test:
	./test.sh
clean:
	rm -f tmp* 9cc