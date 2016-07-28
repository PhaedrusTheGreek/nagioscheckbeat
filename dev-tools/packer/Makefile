BUILDID=$(shell git rev-parse HEAD)

.PHONY: all
all: nagioscheckbeat/deb nagioscheckbeat/rpm nagioscheckbeat/darwin nagioscheckbeat/win nagioscheckbeat/bin \
	build/upload/build_id.txt

ES_BEATS?=../../vendor/github.com/elastic/beats
include $(ES_BEATS)/dev-tools/packer/scripts/Makefile


.PHONY: nagioscheckbeat
nagioscheckbeat: build/upload
	# cross compile on ubuntu
	docker run --rm \
		-v $(abspath build):/build \
		-v $(abspath $(ES_BEATS)/dev-tools/packer/xgo-scripts):/scripts \
		-v $(abspath ../..):/source \
		-e PUREGO="yes" \
		-e PACK=$@ \
		-e BEFORE_BUILD=before_build.sh \
		-e SOURCE=/source \
		-e TARGETS="linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64" \
		-e BUILDID=${BUILDID} \
		tudorg/beats-builder \
		github.com/PhaedrusTheGreek/nagioscheckbeat
