DATASRC=$(wildcard data/*.csv)
DATABIN=$(patsubst %,%.go,${DATASRC})

default:
	@echo "usage:"
	@echo "    make bindata"
	@echo

bindata: data

data: $(DATABIN)

data/%.csv.go: data/%.csv
	rm -f data/$*.csv.go
	cat data/$*.csv | go-bindata -func $(subst .,_,$*)_csv -pkg data | gofmt > data/$*.csv.go

clean:
	rm $(DATABIN)

.PHONY: clean data
