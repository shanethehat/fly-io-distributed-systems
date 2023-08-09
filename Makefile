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

.PHONY: echo
echo:
	(cd c1-echo && go install .) && \
	$(MAELSTROM_BIN) test -w echo --bin /go/bin/echo --log-stderr --node-count 1 --time-limit 10