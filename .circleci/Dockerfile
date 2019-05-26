FROM golang:1.12.5

RUN apt-get update && apt-get install -y --no-install-recommends \
		git \
		ssh \
		tar \
		gzip \
		ca-certificates \
		libssl-dev \
		cmake \
		libssh2-1 \
		libcurl4-openssl-dev \
	&& rm -rf /var/lib/apt/lists/*

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH