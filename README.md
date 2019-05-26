[![GoDoc](https://godoc.org/github.com/isacikgoz/libgit2-api?status.svg)](https://godoc.org/github.com/isacikgoz/libgit2-api) [![CircleCI](https://img.shields.io/circleci/build/github/isacikgoz/libgit2-api.svg)](https://circleci.com/gh/isacikgoz/libgit2-api/tree/master)

# libgit2-api

This project aim to be an idiomatic and simple interface to libgit2 and eventually to a git repository. The main idea is to use it just as you use in command line.

The main reason of developing such library is to use it on my various git tools. For now basic functionalities will be implemented. These are:

- Clone
- Checkout
- Pull
- Log (kind of)

# Using this API

First of all you should get this project by:

`go get -d github.com/isacikoz/libgit2-api`

After you downloaded it you will need a compiled libgit2 library on your operating system. If you are on macOS and using brew, you can install libgit2 via `brew install libgit2`, if you want to build it by yourself here you go:
- make sure you have following libraries and tools installed:
  - cmake
  - pkg-config
  - libssl
  - libssh2
- run the script file (Linux): `scripts/install-libgit2.sh`

After you install required library, you can use this API. Also, I am considering to supply a sample make file so that you can build your Go application with this project.
