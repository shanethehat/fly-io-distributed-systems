MAELSTROM_BIN = bin/maelstrom/maelstrom

.PHONY: install
install: install-packages install-maelstrom

.PHONY: install-packages
install-packages:
	sudo apt-get update && sudo apt-get install -y graphviz gnuplot

.PHONY: install-maelstrom
install-maelstrom:
	mkdir -p bin && \
	cd bin && \
	curl -O -L https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 && \
	tar -xvf maelstrom.tar.bz2

.PHONY: test
test: test-common

.PHONY: test-common
test-common:
	cd common && go test

.PHONY: echo
echo:
	(cd c1-echo && go install .) && \
	$(MAELSTROM_BIN) test -w echo --bin /go/bin/echo --log-stderr --node-count 1 --time-limit 10

.PHONY: generate
generate:
	(cd c2-generate && go install .) && \
	$(MAELSTROM_BIN) test -w unique-ids --bin /go/bin/generate --log-stderr --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

.PHONY: broadcast-1
broadcast-1:
	(cd c3-broadcast && go install .) && \
	$(MAELSTROM_BIN) test -w broadcast --bin /go/bin/broadcast --log-stderr --node-count 1 --time-limit 20 --rate 10

.PHONY: broadcast-2
broadcast-2:
	(cd c3-broadcast && go install .) && \
	$(MAELSTROM_BIN) test -w broadcast --bin /go/bin/broadcast --log-stderr --node-count 5 --time-limit 20 --rate 10