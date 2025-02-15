⚠️⚠️⚠️ **THIS IS A WORK IN PROGRESS** ⚠️⚠️⚠️

# axelar-core

The axelar-core app based on the Cosmos SDK is the main application of the axelar network. This repository is used to
build the necessary binaries and docker image to run a core node.

## How To Build

_Note:_ For a release build, checkout the release tag via `git checkout vX.Y.Z` first.

Execute `make build` to create the `axelard` binary in the `./bin` folder.

## Creating docker images

To create a regular docker image for the node, execute `make docker-image`. This creates the image axelar/core:
latest.

To create a docker image for debugging (with [delve](https://github.com/go-delve/delve)),
execute `make docker-image-debug`. This creates the image axelar/core-debug:latest.

### Smart contracts bytecode dependency

In order to run/build the project locally we need to import the bytecode from gateway smart contracts.

1. Find the specific version of the bytecode here `contract-version.json`
2. Download the right version from the [gateway workflow](https://github.com/axelarnetwork/solidity-cgp-gateway/actions/workflows/publish-bytecode.yaml).
   Example: `Bytecode-v4.3.0`
3. Unzip the json files under `contract-artifacts/gateway/`
4. Run `make generate` to generate `x/evm/types/contracts.go`

## Interacting with a local node

With a local (dockerized) node running, the `axelard` binary can be used to interact with the node.
Run `./bin/axelard` or `./bin/axelard <command> --help` after building the binaries to get information about the available commands.

## Show API documentation

Execute `GO111MODULE=off go install -u golang.org/x/tools/cmd/godoc` to ensure that `godoc` is installed on the host.

After the installation, execute `godoc -http ":{port}" -index` to host a local godoc server. For example, with
port `8080` and `godoc -http ":8080" -index`, the documentation is hosted at
http://localhost:8080/pkg/github.com/axelarnetwork/axelar-core. The index flag makes the documentation searchable.

Comments at the beginning of packages, before types and before functions are automatically taken from the source files
to populate the documentation. See https://blog.golang.org/godoc for more information.

### CLI command documentation

For the full list of available CLI commands for `axelard` see [here](docs/cli/toc.md)

## Test tools

Dev tool dependencies, such as `moq` and `goimports`, can be installed via `make prereqs`.
Make sure they're on available on your `PATH`.

## Bug bounty and disclosure of vulnerabilities

See the [Axelar documentation website](https://docs.axelar.dev/bug-bounty).
