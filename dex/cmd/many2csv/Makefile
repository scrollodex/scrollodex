
all: data

SRC = main.go tidy.go ../../reslist/main.go ../../reslist/model.go ../../reslist/fs.go ../../reslist/git.go

clean:
	rm -f many2csv

many2csv:  $(SRC)
	go build -o many2csv main.go tidy.go

.PHONY: data
data: many2csv
	#./many2csv file:/Users/tlimoncelli/gitthings/beta-scollodex-db-poly data/poly-
	#./many2csv file:/Users/tlimoncelli/gitthings/beta-scollodex-db-bi data/bi-
	#
	./many2csv /Users/tlimoncelli/gitthings/scrollodex-db-bi data/bi-
	./many2csv /Users/tlimoncelli/gitthings/scrollodex-db-poly data/poly-
	#
	#./many2csv git@github.com:scrollodex/scrollodex-db-poly.git data/poly-
	#./many2csv git@github.com:scrollodex/scrollodex-db-bi.git data/bi-
