# Copyright Axis Communications AB.
#
# For a full list of individual contributors, please see the commit history.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GIT = git
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGENERATE = $(GOCMD) generate
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test -race -cover
GOLANGCI_LINT = $(GOBIN)/golangci-lint

GOLANGCI_LINT_VERSION := v1.42.1
GOLANGCI_LINT_INSTALLATION_SHA256 := 294771225087ee48c8e0a45a99ac82ed8f9c6e9d384e692ab201986479c8594f
GOLANGCI_LINT_BINARY_SHA256 := f4d62220f2484f1584f91791a040dd2ba8af7a49c5ef151fad63c75f18a8db44

# Install tools locally instead of in $HOME/go/bin.
export GOBIN := $(CURDIR)/bin
export PATH := $(GOBIN):$(PATH)

.PHONY: all
all: gen
	$(GOBUILD) .

.PHONY: gen
gen:
	$(GOGENERATE) ./...

.PHONY: check
check: staticcheck test

.PHONY: staticcheck
staticcheck: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run

.PHONY: test
test:
	$(GOTEST) ./...

.PHONY: tidy
tidy:
	$(GOMOD) tidy

.PHONY: check-dirty
check-dirty:
	$(GIT) diff --exit-code HEAD

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(PROGRAM)

# Download the installation script for golangci-lint, verify its SHA-256 digest,
# run it if everything checks out, and verify the resulting binary.
$(GOLANGCI_LINT):
	mkdir -p $(dir $@)
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/$(GOLANGCI_LINT_VERSION)/install.sh \
		> $@.install-script-unverified
	echo "$(GOLANGCI_LINT_INSTALLATION_SHA256) $@.install-script-unverified" | sha256sum -c --quiet -
	sh -s -- -b $(dir $@) $(GOLANGCI_LINT_VERSION) < $@.install-script-unverified
	rm -f $@.install-script-unverified
	echo "$(GOLANGCI_LINT_BINARY_SHA256) $@" | sha256sum -c --quiet - || ( rm $@ ; exit 1 )

