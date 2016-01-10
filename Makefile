BEATNAME=nagioscheckbeat
SYSTEM_TESTS=true

# Only crosscompile for linux because other OS'es use cgo.
GOX_OS=linux

include /go/src/github.com/elastic/beats/libbeat/scripts/Makefile

.PHONY: install-cfg
install-cfg:
	cp etc/nagioscheckbeat.template.json $(PREFIX)/nagioscheckbeat.template.json
	# linux
	cp etc/nagioscheckbeat.yml $(PREFIX)/nagioscheckbeat-linux.yml
	# binary
	cp etc/nagioscheckbeat.yml $(PREFIX)/nagioscheckbeat-binary.yml
	# darwin
	cp etc/nagioscheckbeat.osx.yml $(PREFIX)/nagioscheckbeat-darwin.yml
	# win
	cp etc/nagioscheckbeat.yml $(PREFIX)/nagioscheckbeat-win.yml
