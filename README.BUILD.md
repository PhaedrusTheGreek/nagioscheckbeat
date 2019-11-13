1. Set up your golang environment

2. Download Me and my Dependencies

```
go get github.com/PhaedrusTheGreek/nagioscheckbeat
```

3.  Checkout the desired version of `elastic/beats` 

```
cd $GOPATH/src/github.com/elastic/beats
git checkout 7.x
```

4. Build

```
cd $GOPATH/src/github.com/PhaedrusTheGreek/nagioscheckbeat
make setup
mage build 
go get -u github.com/kardianos/govendor
govendor fetch mgutz/str +out
make release
```
