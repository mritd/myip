BUILD_VERSION   := $(shell cat version)

all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

release: all
	ghr -u mritd -t $(GITHUB_RELEASE_TOKEN) -replace -recreate --debug ${BUILD_VERSION} dist

clean:
	rm -rf dist

docker: all
	docker build -t mritd/myip:${BUILD_VERSION} .

.PHONY : all release clean docker

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://athens.azurefd.net
