default: build

BIN=time-tweet

all: build

deps: .vendor

.vendor: Goopfile.lock
	goop install

clean:
	rm $(BIN)

build: deps
	goop exec go build -o $(BIN)

