tapper
====

## Description

tapper is test tool for Golang.
It convert `go test` output to TAP.

## Usage

```bash
$ tapper
ok 1 - TestOk (0.00 seconds)
not ok 2 - TestNg (0.00 seconds)
1..2
```

To use with `prove`, use `tapper-runner`:

```bash
$ prove -v --exec="bash" tapper-runner
tapper-runner ..
ok 1 - TestOk (0.00 seconds)
not ok 2 - TestNg (0.00 seconds)
1..2
Failed 1/2 subtests

Test Summary Report
-------------------
tapper-runner (Wstat: 0 Tests: 2 Failed: 1)
  Failed test:  2
Files=1, Tests=2,  1 wallclock secs ( 0.03 usr  0.00 sys +  0.21 cusr  0.07 csys =  0.31 CPU)
Result: FAIL
```

You can set options for `go test` command:

```bash
$ tapper --options="-run Ok"
ok 1 - TestOk (0.00 seconds)
1..1
```

`tapper-runner` too:

```bash
$ prove -v --exec="bash" tapper-runner :: "-run Ok"
tapper-runner ..
ok 1 - TestOk (0.00 seconds)
1..1
ok
All tests successful.
Files=1, Tests=1,  0 wallclock secs ( 0.03 usr  0.01 sys +  0.25 cusr  0.13 csys =  0.42 CPU)
Result: PASS
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/handlename/tapper
```

After `get`, link `tapper-runner` to your bin directory.

```bash
$ ln -s $GOPATH/src/github.com/handlename/tapper/tapper-runner $GOPATH/bin/tapper-runner
```

## Contribution

1. Fork ([https://github.com/handlename/tapper/fork](https://github.com/handlename/tapper/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[handlename](https://github.com/handlename)
