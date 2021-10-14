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

GOLANGCI_LINT_VERSION = v1.42.1

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

$(GOLANGCI_LINT):
	curl -sfL \
			https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION)
