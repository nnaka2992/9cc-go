SRCS=$(wildcard *.go)
OBJ=$(9cc)

9cc:
	go build -o 9cc *.go
$(SRCS): go.mod
$(OBJ): $(SRCS)

.PHONY: deps test clean cc run
deps:
	go mod tidy
test:
	./test.sh
clean:
	rm -f tmp* 9cc
cc:
	cc -o tmp tmp.s
