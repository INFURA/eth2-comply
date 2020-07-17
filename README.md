# eth2-comply

eth2-comply is a no-code platform for [Ethereum 2.0 API](https://github.com/ethereum/eth2.0-APIs) conformance testing. It provides blackbox conformance testing facilities and a simple syntax for specifying test cases.

## Install

Check the [release page](https://github.com/INFURA/eth2-comply/releases) for the latest eth2-comply binaries and test files.

## Usage

```
eth2-comply --target http://localhost:5051
```

CLI arguments:

- `--testsRoot` Path to a directory tree on the filesystem containing JSON test cases.
- `--testsRemote` URL to a zip file containing a valid tests directory tree. `--testsRemote` takes precedence over `--testsRoot` if both are specidied.
- `--outDir` Location on the filesystem to download and unpack a zip file specified in `--testsRemote`. Has no meaning if `--testsRemote` is not specified.
- `--target` URL of any appliance serving the Ethereum 2.0 API.
- `--timeout` Time after which to abandon waiting tests. Defaults to 10 minutes. Uses [Go duration syntax](https://golang.org/pkg/time/#ParseDuration).
- `--subset` The subset of paths to run tests for. For example, set this to "/v1/node" to only run tests for routes in that path. Defaults to "/" (all paths).
- `--failSilent` When true, return a 0 code even when tests fail. Defaults to false.

## Syntax of test cases

Tests are specified as normal JSON objects. Only one test should be specified per file. Test files should be placed in the appropriate directory according to the API route they activate. Files should be named `<expected-status-code>_<seq_num>.json` where `seq_num` is just a unique number to prevent the file name from colliding with any other test case in that directory testing for the same expected status code in that route. The names and filepaths of test files have no bearing on how eth2-comply processes them, this topology is just an organizational methodology.

The following JSON fields have meaning in eth2-comply test case syntax. Any fields not listed have no meaning to eth2-comply and are ignored.

| field              | required | value type | example                                          |
|--------------------|----------|------------|--------------------------------------------------|
| method             | yes      | string     | "GET"                                            |
| route              | yes      | string     | "/beacon/committees"                             |
| reqBody            | no       | object     | `{"epoch": "0", "pubkeys": ["0xdeadbeef"]}`      |
| queryParams        | no       | object     | `{"epoch": "0"}`                                 |
| awaitSlot          | no       | int        | 2666                                             |
| expectedRespStatus | no       | int        | 200                                              |
| expectedRespBody   | no       | object     | `[{"slot": "0", "index": "0", "committee": []}]` |

Most of the fields meaning should be self-explanatory. `awaitSlot` can be used to make `eth2-comply` wait until the target node has synced the specified slot before executing the test.

When specifying expected response bodies, know that received and expected responses are canonicalized before being compared. This means that whitespace and key order do not matter in general. Remember that list order does matter; the way a list is specified literally is its canonical form, though nested objects are themselves canonicalized.

## Build and run while developing

Build:

```
make build
```

Run:

```
bazel run -- cmd/eth2-comply:eth2-comply --target http://localhost:5051 --testsRoot $(PWD)/tests
```

