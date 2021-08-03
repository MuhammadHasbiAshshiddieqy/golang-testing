# golang-testing
Unit Test With Go

**MAIN COMMAND :**

go test -v -run=TestFunc -bench=BenchmarkFunc // RUN A TEST AND A BENCHMARK
go test -v -run=TestNothing -bench=. // ONLY RUN ALL BENCHMARK
go test -v ./... // RUN ALL TEST
go test -v bench=. ./... // RUN ALL TEST AND BENCHMARK
