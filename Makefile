VERSION=`date -u +v%Y%m%d.%H%M%S`

clean:
	rm -f 10bisbar*
	rm -f *.tar.gz

test:
	@gom test

get:
	go get

build: clean get
	go build -ldflags "-X main.version=$(VERSION) -s"
	upx -9 10bisbar

dist: build
	mv 10bisbar 10bisbar.15m.sh

release: build
	tar czvf 10bisbar.$(VERSION).tar.gz 10bisbar
	git tag -a $(VERSION) -m "release $(VERSION)"


.PHONY: test build dist release clean
