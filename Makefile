BINDIR = $(PWD)/bin

.PHONY: all
all: db

.PHONY: install
install:
	GOBIN=$(BINDIR) go install ./examples/...

.PHONY: clean
clean: clean/db
	rm -f $(BINDIR)/*

.PHONY: clean/db
clean/db:
	psql -aq -d postgres -f .make/db-clean.sql

.PHONY: db
db: clean/db
	psql -aq -d postgres            -f .make/db-setup.sql
	psql -aq -d pgxer_todo          -f examples/todo/structure.sql
	psql -aq -d pgxer_url_shortener -f examples/url_shortener/structure.sql

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: build
build: $(BINDIR)/chat $(BINDIR)/todo $(BINDIR)/url_shortener

$(BINDIR)/chat: $(wildcard examples/chat/*)
	go build -o $(BINDIR)/chat          ./examples/chat

$(BINDIR)/todo: $(wildcard examples/todo/*)
	go build -o $(BINDIR)/todo          ./examples/todo

$(BINDIR)/url_shortener: $(wildcard examples/url_shortener/*)
	go build -o $(BINDIR)/url_shortener ./examples/url_shortener
