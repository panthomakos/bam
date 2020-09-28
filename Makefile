version=$(shell cat VERSION)
name=bam
platforms=linux-amd64 linux-386 darwin-amd64
prefix=release/${name}-v${version}-
suffix=.tar.gz
packages=$(addsuffix ${suffix}, $(addprefix ${prefix}, ${platforms}))

.PHONY: clean/build clean/release clean
.PHONY: package build

build:
	go build -o ${name} main.go

clean/build:
	rm -f ${name}

clean/release:
	rm -rf release

clean: clean/build clean/release

package: ${packages}

release:
	mkdir -p release

${prefix}%${suffix}: release
	$(eval platform:=$(subst -, ,$*))
	$(eval goos:=$(word 1, ${platform}))
	$(eval goarch:=$(word 2, ${platform}))
	env GOOS=${goos} GOARCH=${goarch} go build -o release/${name} main.go
	cd ./release && tar -czf $(notdir $@) ${name}
	@rm release/${name}
