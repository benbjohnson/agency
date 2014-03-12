bindata.go: data/%.csv
	go-bindata -pkg=agency -prefix=data data

data/%.csv:
	
test: bindata.go
	go test -v

.PHONY: test
