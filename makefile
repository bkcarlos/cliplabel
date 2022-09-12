# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PUTPATH=output
BINARY_NAME=lucas-go-be

#all: build
.PHONY: build
build: cliplabel
	if [ -d $(PUTPATH) ] ; then rm -rf ${PUTPATH} ; fi
	if [ ! -d $(PUTPATH) ];then mkdir $(PUTPATH); fi
	mv -f $^ ./$(PUTPATH)/
	cp -r ./configs/*.yaml ./$(PUTPATH)/
	cp -r lib ./$(PUTPATH)/
	
	echo 'git version:' > output/version.txt
	git log | head -n 3 >> output/version.txt
	echo '' >> output/version.txt
	echo 'build time:' >> output/version.txt
	date >> output/version.txt

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	if [ -d $(PUTPATH) ] ; then rm -rf ${PUTPATH} ; fi

check:
	go fmt ./...
	go vet ./...
	golint ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	nohup ./$(BINARY_NAME) &

cliplabel:
	wget -qO- https://github.com/wangfenjin/simple/releases/latest/download/libsimple-osx-x64.zip | tar xf - 
	if [ ! -d ./lib/libsimple ];then mkdir ./lib/libsimple; fi
	cp -rf libsimple-osx-x64/* ./lib/libsimple/
	rm -r libsimple-osx-x64
	$(GOBUILD) --tags fts5 -o $@ main.go
