buildin:
	go build -o ./build/kvstore-buildin ./kvstore/buildin/main.go

all: buildin

.PHONY: buildin
