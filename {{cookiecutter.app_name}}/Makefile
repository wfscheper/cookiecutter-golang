TARGET   = {{cookiecutter.app_name}}
{% if cookiecutter.go_gettable == "y" -%}
PACKAGE  = github.com/{{cookiecutter.github_username}}/$(TARGET)
{% else -%}
PACKAGE  = $(TARGET)
{% endif -%}
DATE    ?= $(shell date +%FT%T%z)
COMMIT   = $(shell git rev-parse HEAD)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.app_name}}"
BIN      = $(GOPATH)/bin
{% if cookiecutter.go_gettable == "y" -%}
PKGS     = $(or $(PKG),$(shell $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
{% raw %}TESTPKGS = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS)){% endraw %}
{% else -%}
GOPATH   = $(CURDIR)/.gopath~
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = {% raw %}$(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS)){% endraw %}
{% endif -%}

GO      = go
GOBUILD = $(GO) build -v
GODOC   = godoc
GOFMT   = gofmt
DOCKER  = docker
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: fmt lint vendor {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building executable…) @ ## Build program binary
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GO) build \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE) -X $(PACKAGE)/cmd.Commit=$(COMMIT)' \
		-o bin/$(TARGET) main.go

{% if cookiecutter.go_gettable != "y" -%}
$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@
{% endif -%}

{% if cookiecutter.use_docker == "y" -%}
build-alpine: fmt lint vendor ; $(info $(M) building apline executable…) @ ## Build program binary
	$Q $(GOBUILD) $(TAGS) \
		-ldflags '-w -linkmode external -extldflags "-static" -X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE) -X $(PACKAGE)/cmd.Commit=$(COMMIT)' \
		-o bin/$(TARGET) main.go

package: ; $(info $(M) building image…) @ ## Build docker image
	$Q $(DOCKER) build \
		--build-arg VERSION=${VERSION} \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg BUILD_DATE=$(DATE) \
		-t $(IMAGE_NAME):local .

tag: ; $(info $(M) tagging image…) @ ## Tag docker image
	$Q $(DOCKER) tag $(IMAGE_NAME):local $(IMAGE_NAME):$(VERSION)
	$Q $(DOCKER) tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag ; $(info $(M) push docker image…) @ ## Push docker imates
	$Q $(DOCKER) push $(IMAGE_NAME):$(VERSION)
	$Q $(DOCKER) push $(IMAGE_NAME):latest

{% endif -%}
# Tools

GODEP = $(BIN)/dep
$(BIN)/dep: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building go dep…)
	$Q go get github.com/golang/dep/cmd/dep

GOLINT = $(BIN)/golint
$(BIN)/golint: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building golint…)
	$Q go get github.com/golang/lint/golint

GOCOVMERGE = $(BIN)/gocovmerge
$(BIN)/gocovmerge: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building gocovmerge…)
	$Q go get github.com/wadey/gocovmerge

GOCOV = $(BIN)/gocov
$(BIN)/gocov: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building gocov…)
	$Q go get github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building gocov-xml…)
	$Q go get github.com/AlekSi/gocov-xml

GO2XUNIT = $(BIN)/go2xunit
$(BIN)/go2xunit: {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) building go2xunit…)
	$Q go get github.com/tebeka/go2xunit

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt lint vendor {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint vendor | {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GO2XUNIT) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}2>&1 $(GO) test -timeout 20s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

COVERAGE_MODE = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML)
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: fmt lint vendor test-coverage-tools {% if cookiecutter.go_gettable != "y" %}| $(BASE) {% endif %}; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}for pkg in $(TESTPKGS); do \
		$(GO) test \
			{% raw %}-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $$pkg | \{% endraw %}
					grep '^$(PACKAGE)/' | grep -v '^$(PACKAGE)/vendor/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	 done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: vendor | {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	{% raw %}@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \{% endraw %}
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Dependency management

Gopkg.toml: | {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GODEP); $(info $(M) generating Gopkg.toml…)
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GODEP) init

Gopkg.lock: | Gopkg.toml {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GODEP); $(info $(M) updating Gopkg.lock…)
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GODEP) ensure -no-vendor

vendor: Gopkg.lock | {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GODEP) ; $(info $(M) retrieving dependencies…)
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GODEP) ensure
	@touch $@
.PHONY: vendor-update
vendor-update: | {% if cookiecutter.go_gettable != "y" %}$(BASE) {% endif %}$(GODEP)
ifeq "$(origin PKG)" "command line"
	$(info $(M) updating $(PKG) dependency…)
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GODEP) ensure -update $(PKG)
else
	$(info $(M) updating all dependencies…)
	$Q {% if cookiecutter.go_gettable != "y" %}cd $(BASE) && {% endif %}$(GODEP) ensure -update
endif
	@touch vendor

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
{%- if cookiecutter.go_gettable != "y" %}
	@rm -rf $(GOPATH)
{%- endif %}
	@rm -rf bin
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
