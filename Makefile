MAELSTROM_BIN = bin/maelstrom/maelstrom

.PHONY: install
install: install-maelstrom

.PHONY: install-maelstrom
install-maelstrom:
	mkdir -p bin && cd bin && \
	curl -O -L https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 && \
	tar -xvf maelstrom.tar.bz2