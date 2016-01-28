# BangGO

Simple two app project represents simplest bank in the world ;-)

![BankGo Logo](https://miolini.github.io/bankgo/bankgo.jpg)

## Deploy

```bash
$ go get github.com/miolini/bankgo

$ cd $GOPATH/src/github.com/miolini/bankgo

$ make docker
```

## Documentation

[HTTP API Documentation](http://miolini.github.io/bankgo/)

## Test Coverage

### Generate local

```bash
$ make testcoverage
```

### Hosted on Github

[Report](https://miolini.github.io/bankgo/testcoverage.html)

## Source code stats

```bash
$ cloc .
      30 text files.
      30 unique files.                              
      10 files ignored.

http://cloc.sourceforge.net v 1.64  T=0.11 s (200.0 files/s, 19002.9 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               9            104             23            728
XML                              8              0              0            711
HTML                             2             49              6            384
Protocol Buffers                 1              9              0             30
make                             1             10              0             26
YAML                             1              0              0             10
-------------------------------------------------------------------------------
SUM:                            22            172             29           1889
-------------------------------------------------------------------------------
```