BIN=pointercast

.PHONY: c
c: $(BIN)
	./$<

$(BIN): pointercast.c
	gcc $^ -o $@

.PHONY: go
go: 
	go run pointercast.go
