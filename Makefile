app := runandlint
out := /home/adithyaaaa/vsc/forked/go-refresher/bin/

build:
	@mkdir -p /home/adithyaaaa/vsc/forked/bin
	go build -o $(out)/anonymfns /home/adithyaaaa/vsc/forked/go-refresher/anonymousfunctions/anonymfns.go
	go build -o $(out)/simple_flags /home/adithyaaaa/vsc/forked/go-refresher/flags/simple_flags.go
	go build -o $(out)/parse_file /home/adithyaaaa/vsc/forked/go-refresher/jsonparsing/parse_file.go
	go build -o $(out)/data_collection /home/adithyaaaa/vsc/forked/go-refresher/structs/data_collection.go

lint:
	golangci-lint run


run:build
	$(out)/anonymfns
	$(out)/simple_flags
	$(out)/parse_file
	$(out)/data_collection
	