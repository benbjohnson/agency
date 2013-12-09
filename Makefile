DATASRC=$(wildcard data/*.csv)
DATABIN=$(patsubst %,%.go,${DATASRC})
FIXSRC=$(wildcard fixtures/*.csv)
FIXBIN=$(patsubst %,%.go,${FIXSRC})

default:
	@echo "usage:"
	@echo "    make bindata"
	@echo

bindata: data fixtures

data: $(DATABIN)

data/%.csv.go: data/%.csv
	rm -f data/$*.go
	cat data/$*.csv | go-bindata -func $(subst .,_,$*)_csv -pkg data | gofmt > data/$*.csv.go

fixtures: $(FIXBIN)

fixtures/%.csv.go: fixtures/%.csv
	rm -f fixtures/$*.go
	cat fixtures/$*.csv | go-bindata -func $(subst .,_,$*)_csv -pkg fixtures | gofmt > fixtures/$*.csv.go

clean:
	rm $(DATABIN) $(FIXBIN)

.PHONY: clean data fixtures
