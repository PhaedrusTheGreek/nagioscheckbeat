1. Fetch Dependencies

```
cd nagioscheckbeat
go get -u github.com/kardianos/govendor
govendor fetch mgutz/str +out
```

2.  Checkout the desired version of `elastic/beats` 

```
cd $GOPATH/github.com/elastic/beats
git checkout 7.x
```

3. Do it

```
make setup
mage build 
make release
```
