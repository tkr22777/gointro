# simplepackage
Trying out Go lang

Execute all tests in a test files:
    go test -v simplepackage.go *the_test_filename.go*
    example:
        go test -v simplepackage.go goroutines_test.go

Execute a specific test in a test files:
    go test -cover -v simplepackage.go *the_test_filename.go* -run *test_name*
    example:
    go test -cover -v simplepackage.go goroutines_test.go -run TestChannelledFibonacci
