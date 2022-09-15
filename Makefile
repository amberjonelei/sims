# go through all the sims

# when building for a release:
# IMPORTANT: for mac, export MACOSX_DEPLOYMENT_TARGET=10.8

TOPTARGETS := all clean mac linux windows

SUBDIRS := $(wildcard */.)

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

tidy: export GO111MODULE = on
tidy:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go mod tidy

old:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go list -u -m all | grep '\['
	
# updates go.mod to master for all of the goki dependencies
# note: must somehow remember to do this for any other depend
# that we need to update at any point!
master: export GO111MODULE = on
master:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go get -u github.com/emer/etable@master
	go get -u github.com/emer/emergent@master
	go get -u github.com/emer/leabra@master
	go list -m all | grep emer
	go mod tidy
	
mod-update: export GO111MODULE = on
mod-update:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go get -u ./...
	go mod tidy

# gopath-update is for GOPATH to get most things updated.
# need to call it in a target executable directory
gopath-update: export GO111MODULE = off
gopath-update:
	@echo "GO111MODULE = $(value GO111MODULE)"
	cd cmd/pi; go get -u ./...

# NOTE: MUST update version number here prior to running 'make release' and edit this file! 
VERS=v1.3.2
PACKAGE=sims
GIT_COMMIT=`git rev-parse --short HEAD`
VERS_DATE=`date -u +%Y-%m-%d\ %H:%M`
VERS_FILE=version.go

release:
	/bin/rm -f $(VERS_FILE)
	@echo "// WARNING: auto-generated by Makefile release target -- run 'make release' to update" > $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	@echo "package $(PACKAGE)" >> $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	@echo "const (" >> $(VERS_FILE)
	@echo "	Version = \"$(VERS)\"" >> $(VERS_FILE)
	@echo "	GitCommit = \"$(GIT_COMMIT)\" // the commit JUST BEFORE the release" >> $(VERS_FILE)
	@echo "	VersionDate = \"$(VERS_DATE)\" // UTC" >> $(VERS_FILE)
	@echo ")" >> $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	goimports -w $(VERS_FILE)
	/bin/cat $(VERS_FILE)
	git commit -am "$(VERS) release"
	git tag -a $(VERS) -m "$(VERS) release"
	git push
	git push origin --tags

