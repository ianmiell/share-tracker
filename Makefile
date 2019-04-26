run: init safe_clean shares
	rm -f shares.db
	cat db_export.sql | sqlite3 shares.db
	./shares
	echo ".dump" | sqlite3 shares.db > db_export.sql
	#./insert_profits.sh
	./queries.sh
	echo ".dump" | sqlite3 shares.db > db_export.sql
	python graphs.py
	git add graph_data
	git add graphs
	git commit -am re-run || true
	git pull --rebase -s recursive -X ours
	git push

db:
	rm -f shares.db
	cat db_export.sql | sqlite3 shares.db
	./queries.sh

.PHONY: install

init:
	go version
	PATH=/usr/local/bin:${PATH} pip --version
	PATH=/usr/local/bin:${PATH} pip list --format=columns | grep shutit
	sqlite3 --version
	git --version
	python --version
	gnuplot --version
	GOPATH=/space/go go get github.com/mattn/go-sqlite3

shares: *.go
	GOPATH=/space/go go build

# Removes everything (tho we still have git history).
# Removes data that is recreated with a make
safe_clean:
	rm -f share-tracker.db
	cp /dev/null db_export.db
	for f in $(ls reports/*txt); do cp /dev/null $f; done
	cp /dev/null graph_data/dividend_12mth.dat
