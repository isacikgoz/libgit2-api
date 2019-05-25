GOPATH_DIR?=$(shell go env GOPATH | cut -d: -f1)

GIT2GO_VERSION=27
GIT2GO_DIR:=$(GOPATH_DIR)/src/gopkg.in/libgit2/git2go.v$(GIT2GO_VERSION)

all: update
	make -C $(GIT2GO_DIR) install-static

.PHONY: update
update:
	git submodule -q foreach --recursive git reset -q --hard
	git submodule update --init --recursive
