COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
CPP = $(shell which cpp)

.PHONY: all
all: fluxgo

BUILDENVVAR=CGO_CFLAGS="-I/usr/include" CGO_LDFLAGS="-L/usr/lib -L/usr/lib/flux -L/usr/lib/flux/resource -lflux-core -lflux-idset -lflux-hostlist -lstdc++ -ljansson"

.PHONY: fluxgo
fluxgo: 
	mkdir -p ./bin
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./bin/fluxgo-submit cmd/submit.go
