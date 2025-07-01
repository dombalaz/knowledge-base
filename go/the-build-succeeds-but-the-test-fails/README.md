# The build succeeds, but the test fails.

This example demonstrates the behvaiour of the `go` tooling, when the `go test`
command apparently fails, because it cannot build the binary.

```
➜  vet-test git:(master) ✗ go test main.go
# command-line-arguments
./main.go:10:2: fmt.Printf does not support error-wrapping directive %w
FAIL    command-line-arguments [build failed]
FAIL
```

However, the `go build` finishes successfully:

```
➜  vet-test git:(master) ✗ go build main.go
➜  vet-test git:(master) ✗ ./main
v: I am an error, w: %!w(*errors.errorString=&{I am an error})%
```

This is because the `go test` also executes `go vet` before running.

```
➜  vet-test git:(master) ✗ go vet main.go
# command-line-arguments
# [command-line-arguments]
./main.go:10:2: fmt.Printf does not support error-wrapping directive %w
```

So don't forget to run `go test` or `go vet` before commiting any changes to go.
