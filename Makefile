NAME=afrostream-api

default: build

clean:
	rm -rf ./$(NAME)

clean-deps:
	rm -rf ./Godeps/
	rm -rf ./vendor/

build-deps:
	godep save

build:
	go build

heroku: clean clean-deps build-deps
	@echo "you should commit Godeps & vendor"

install:
	@echo "installed programm should be in "$(GOROOT)"/bin/"$(NAME)
	go install

all: clean build
