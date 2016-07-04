# Description

afrostream api v2

# Install

## Install golang

you need golang 1.6   
you need to set your *$GOPATH* env var (assuming ~/go here.)   
(FIXME: docker)

## Install this program

```
cd $GOPATH/src
mkdir -p $GOPATH/src/github.com/afrostream
cd $GOPATH/src/github.com/afrostream
git clone git@github.com:Afrostream/afrostream-api.git
```

## Install the dependency

```
go get ./...
```

to install afrostream-go-lib in "dev mode" ,   
@see https://github.com/Afrostream/afrostream-go-lib

# Usage

## start

```
./afrostream-api
```

## build

```
make
```

## clean

```
make clean
```

## install

should be install in $GOPATH/bin/afrostream-api

```
make install
```
