1. Set up your golang environment

2. Fetch Dependencies

```
go get github.com/PhaedrusTheGreek/nagioscheckbeat
cd $GOPATH/src/github.com/PhaedrusTheGreek/nagioscheckbeat
go get -u github.com/kardianos/govendor
govendor fetch mgutz/str +out
```

3.  Checkout the desired version of `elastic/beats` 

```
cd $GOPATH/src/github.com/elastic/beats
git checkout 7.x
```

4. Do it

```
cd $GOPATH/src/github.com/PhaedrusTheGreek/nagioscheckbeat
make setup
mage build 
make release
```
