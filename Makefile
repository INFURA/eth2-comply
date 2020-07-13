ETH2SPECREMOTE=git@github.com:ethereum/eth2.0-APIs.git
ETH2SPECFS=/tmp/eth2-api
ETH2SPECROOT=beacon-node-oapi.yaml
ETH2SPECBUNDLE=bundle.yaml
ETH2PKGPATH=pkg/eth2spec

build: lint
	go build ./... \
	  && bazel run //:gazelle

lint:
	golangci-lint --skip-dirs $(ETH2PKGPATH) run 

oapi-pull-latest:
	rm -rf $(ETH2SPECFS) \
	  ; git clone $(ETH2SPECREMOTE) $(ETH2SPECFS)

oapi-build:
	swagger-cli bundle -r $(ETH2SPECFS)/$(ETH2SPECROOT) > $(ETH2SPECFS)/$(ETH2SPECBUNDLE) \
	  && openapi-generator generate \
	  -i $(ETH2SPECFS)/$(ETH2SPECBUNDLE) \
	  -g go \
	  -o $(ETH2PKGPATH) \
	  --additional-properties=packageName=eth2spec,isGoSubmodule=true \
	  && rm $(ETH2PKGPATH)/go.mod \
	  && rm $(ETH2PKGPATH)/go.sum \
	  && rm $(ETH2PKGPATH)/.travis.yml \
	  && rm $(ETH2PKGPATH)/.openapi-generator-ignore \
	  && rm $(ETH2PKGPATH)/git_push.sh \
	  && go build $(ETH2PKGPATH)/*.go \
	  && bazel run //:gazelle \
 	  && bazel run //:gazelle -- update-repos -from_file=go.mod
