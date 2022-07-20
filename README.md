# NetOverlap

Go-based CLI utility that prints to standart output the relation between two network prefixes provided into CIDR format.


## Description

It expects two commands line arguments which must be valid IP prefixes of same type (IPv4 or IPv6).
Once the prefixes are parsed and validated, it goes through a set of checks to find the relation.
In case there is no error during the processing, it outputs the result on your terminal. Below table summarizes possible result.

| Status | Description |
| :--- | :--- |
| **`subset`** | the network of the second address is included in the first one |
| **`superset`** | the network of the second address includes the first one |
| **`different`** | the two networks are not overlapping |
| **`same`** | if both address are in the same network |


## Tests Coverage

You can see the detailed output at the tests at the bottom. Curent coverage is: **96.3%**

```shell
{gopher}:~$ go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
ok      github.com/jeamon/netoverlap    0.312s  coverage: 90.0% of statements
ok      github.com/jeamon/netoverlap/app        0.306s  coverage: 100.0% of statements
github.com/jeamon/netoverlap/app/app.go:32:     NewNetworkInfos         100.0%
github.com/jeamon/netoverlap/app/app.go:47:     IsSameAs                100.0%
github.com/jeamon/netoverlap/app/app.go:55:     IsSubsetOf              100.0%
github.com/jeamon/netoverlap/app/app.go:63:     IsSupersetOf            100.0%
github.com/jeamon/netoverlap/app/app.go:71:     CheckOverlapStatus      100.0%
github.com/jeamon/netoverlap/app/app.go:85:     IsComparableTo          100.0%
github.com/jeamon/netoverlap/app/app.go:90:     Init                    100.0%
github.com/jeamon/netoverlap/app/app.go:99:     Run                     100.0%
github.com/jeamon/netoverlap/main.go:22:        main                    0.0%
github.com/jeamon/netoverlap/main.go:28:        Execute                 100.0%
total:                                          (statements)            96.3%
```


## Setup

If your system has [Go >= 1.17](https://golang.org/dl/) and [git tool](https://git-scm.com/downloads) then you can build an executable from the codebase.
There is a makefile with predefined commands. Otherwise Open your terminal or Git Bash (if you are on windows system) then follow below steps:


* Move into the codebase folder and update depedencies

```shell
$ git clone https://github.com/jeamon/netoverlap.git
$ cd netoverlap
$ go mod tidy
```

* Run all unit tests and make sure all pass (nothing broken)

```shell
$ go test -v ./... -count=1
```

* Build the executable on unix-like system

```shell
$ go build -o bin/netoverlap -a -ldflags "-extldflags '-static' -X 'main.GitCommit=$(git rev-parse --short HEAD)' -X 'main.GitTag=$(git describe --tags --abbrev=0)' -X 'main.BuildTime=$(date -u '+%Y-%m-%d %I:%M:%S %p GMT')'" main.go
```

* Build the executable on windows system (run below batch script file)

```shell
$ build.windows.bat
```

* Check if the executable works on unix-like system

```shell
$ chmod +x ./bin/netoverlap
$ ./bin/netoverlap version
$ ./bin/netoverlap help
```

* Check if the executable works on windows system

```shell
$ bin/netoverlap.exe -v
$ bin/netoverlap.exe -h
```


## Get-Started

Below are some usage examples:

```shell
$> ./netoverlap 10.0.0.0/20 10.0.2.0/24
subset
```

```shell
$> ./netoverlap 10.0.2.0/24 10.0.0.0/20
superset
```

```shell
$> ./netoverlap 10.0.2.0/24 10.0.3.0/24
different
```

```shell
$> ./netoverlap 10.0.2.0/24 10.0.2.10/24
same 
```

```shell
$> ./netoverlap fe80::c845:ea23:ad2e/64 fe80::c845:ea23:ad2e:65b8/64
same 
```

```shell
$> ./netoverlap fe80::c845:ea23:ad2e/64 fe80::c845:ea23:ad2e:65b8/128
subset
```


## Tests Outputs


```shell

{gopher}:~$ go test -v ./... -count=1
=== RUN   TestExecute
=== RUN   TestExecute/valid_[version]_argument
=== RUN   TestExecute/valid_[--version]_argument
=== RUN   TestExecute/valid_[-v]_argument
=== RUN   TestExecute/valid_[help]_argument
=== RUN   TestExecute/valid_[--help]_argument
=== RUN   TestExecute/valid_[-h]_argument
=== RUN   TestExecute/valid_prefixes_arguments
=== RUN   TestExecute/invalid_[-x]_argument
=== RUN   TestExecute/invalid_syntax
=== RUN   TestExecute/invalid_prefixes
--- PASS: TestExecute (0.00s)
    --- PASS: TestExecute/valid_[version]_argument (0.00s)
    --- PASS: TestExecute/valid_[--version]_argument (0.00s)
    --- PASS: TestExecute/valid_[-v]_argument (0.00s)
    --- PASS: TestExecute/valid_[help]_argument (0.00s)
    --- PASS: TestExecute/valid_[--help]_argument (0.00s)
    --- PASS: TestExecute/valid_[-h]_argument (0.00s)
    --- PASS: TestExecute/valid_prefixes_arguments (0.00s)
    --- PASS: TestExecute/invalid_[-x]_argument (0.00s)
    --- PASS: TestExecute/invalid_syntax (0.00s)
    --- PASS: TestExecute/invalid_prefixes (0.00s)
PASS
ok      github.com/jeamon/netoverlap    0.280s
=== RUN   TestNewNetworkInfos
=== RUN   TestNewNetworkInfos/10.0.2.0/24
=== RUN   TestNewNetworkInfos/10.0.2.10/20
=== RUN   TestNewNetworkInfos/10.0.2.0/
=== RUN   TestNewNetworkInfos/10.0.2.0
=== RUN   TestNewNetworkInfos/10.0.2.
--- PASS: TestNewNetworkInfos (0.00s)
    --- PASS: TestNewNetworkInfos/10.0.2.0/24 (0.00s)
    --- PASS: TestNewNetworkInfos/10.0.2.10/20 (0.00s)
    --- PASS: TestNewNetworkInfos/10.0.2.0/ (0.00s)
    --- PASS: TestNewNetworkInfos/10.0.2.0 (0.00s)
    --- PASS: TestNewNetworkInfos/10.0.2. (0.00s)
=== RUN   TestIsSameAs
=== RUN   TestIsSameAs/same_ipv4_prefixes
=== RUN   TestIsSameAs/not_same_ipv4_prefixes
=== RUN   TestIsSameAs/same_ipv6_prefixes
=== RUN   TestIsSameAs/not_same_ipv6_prefixes
--- PASS: TestIsSameAs (0.00s)
    --- PASS: TestIsSameAs/same_ipv4_prefixes (0.00s)
    --- PASS: TestIsSameAs/not_same_ipv4_prefixes (0.00s)
    --- PASS: TestIsSameAs/same_ipv6_prefixes (0.00s)
    --- PASS: TestIsSameAs/not_same_ipv6_prefixes (0.00s)
=== RUN   TestIsSubsetOf
=== RUN   TestIsSubsetOf/subset_ipv4_prefixes
=== RUN   TestIsSubsetOf/not_subset_ipv4_prefixes
=== RUN   TestIsSubsetOf/subset_ipv6_prefixes
=== RUN   TestIsSubsetOf/not_subset_ipv6_prefixes
--- PASS: TestIsSubsetOf (0.00s)
    --- PASS: TestIsSubsetOf/subset_ipv4_prefixes (0.00s)
    --- PASS: TestIsSubsetOf/not_subset_ipv4_prefixes (0.00s)
    --- PASS: TestIsSubsetOf/subset_ipv6_prefixes (0.00s)
    --- PASS: TestIsSubsetOf/not_subset_ipv6_prefixes (0.00s)
=== RUN   TestIsSupersetOf
=== RUN   TestIsSupersetOf/superset_ipv4_prefixes
=== RUN   TestIsSupersetOf/not_superset_ipv4_prefixes
=== RUN   TestIsSupersetOf/superset_ipv6_prefixes
=== RUN   TestIsSupersetOf/not_superset_ipv6_prefixes
--- PASS: TestIsSupersetOf (0.00s)
    --- PASS: TestIsSupersetOf/superset_ipv4_prefixes (0.00s)
    --- PASS: TestIsSupersetOf/not_superset_ipv4_prefixes (0.00s)
    --- PASS: TestIsSupersetOf/superset_ipv6_prefixes (0.00s)
    --- PASS: TestIsSupersetOf/not_superset_ipv6_prefixes (0.00s)
=== RUN   TestCheckOverlapStatus
=== RUN   TestCheckOverlapStatus/same_ipv4_prefixes
=== RUN   TestCheckOverlapStatus/subset_ipv4_prefixes
=== RUN   TestCheckOverlapStatus/superset_ipv4_prefixes
=== RUN   TestCheckOverlapStatus/different_ipv4_prefixes
=== RUN   TestCheckOverlapStatus/same_ipv6_prefixes
=== RUN   TestCheckOverlapStatus/subset_ipv6_prefixes
=== RUN   TestCheckOverlapStatus/superset_ipv6_prefixes
=== RUN   TestCheckOverlapStatus/same_ipv6_prefixes#01
--- PASS: TestCheckOverlapStatus (0.00s)
    --- PASS: TestCheckOverlapStatus/same_ipv4_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/subset_ipv4_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/superset_ipv4_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/different_ipv4_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/same_ipv6_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/subset_ipv6_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/superset_ipv6_prefixes (0.00s)
    --- PASS: TestCheckOverlapStatus/same_ipv6_prefixes#01 (0.00s)
=== RUN   TestIsComparableTo
=== RUN   TestIsComparableTo/should_compare_v4tov4
=== RUN   TestIsComparableTo/should_compare_v6tov6
=== RUN   TestIsComparableTo/should_not_compare_v4tov6
=== RUN   TestIsComparableTo/should_not_compare_v6tov4
--- PASS: TestIsComparableTo (0.00s)
    --- PASS: TestIsComparableTo/should_compare_v4tov4 (0.00s)
    --- PASS: TestIsComparableTo/should_compare_v6tov6 (0.00s)
    --- PASS: TestIsComparableTo/should_not_compare_v4tov6 (0.00s)
    --- PASS: TestIsComparableTo/should_not_compare_v6tov4 (0.00s)
=== RUN   TestInit
=== RUN   TestInit/empty_build_flag
=== RUN   TestInit/non_empty_build_flag
=== RUN   TestInit/mixed_values_build_flag
--- PASS: TestInit (0.00s)
    --- PASS: TestInit/empty_build_flag (0.00s)
    --- PASS: TestInit/non_empty_build_flag (0.00s)
    --- PASS: TestInit/mixed_values_build_flag (0.00s)
=== RUN   TestRun
=== RUN   TestRun/same_ipv4_prefixes
=== RUN   TestRun/subset_ipv4_prefixes
=== RUN   TestRun/superset_ipv4_prefixes
=== RUN   TestRun/different_ipv4_prefixes
=== RUN   TestRun/same_ipv6_prefixes
=== RUN   TestRun/subset_ipv6_prefixes
=== RUN   TestRun/superset_ipv6_prefixes
=== RUN   TestRun/same_ipv6_prefixes#01
=== RUN   TestRun/not_comparable_prefixes
=== RUN   TestRun/invalid_first_prefix
=== RUN   TestRun/invalid_second_prefix
--- PASS: TestRun (0.00s)
    --- PASS: TestRun/same_ipv4_prefixes (0.00s)
    --- PASS: TestRun/subset_ipv4_prefixes (0.00s)
    --- PASS: TestRun/superset_ipv4_prefixes (0.00s)
    --- PASS: TestRun/different_ipv4_prefixes (0.00s)
    --- PASS: TestRun/same_ipv6_prefixes (0.00s)
    --- PASS: TestRun/subset_ipv6_prefixes (0.00s)
    --- PASS: TestRun/superset_ipv6_prefixes (0.00s)
    --- PASS: TestRun/same_ipv6_prefixes#01 (0.00s)
    --- PASS: TestRun/not_comparable_prefixes (0.00s)
    --- PASS: TestRun/invalid_first_prefix (0.00s)
    --- PASS: TestRun/invalid_second_prefix (0.00s)
PASS
ok      github.com/jeamon/netoverlap/app        0.285s

```

## Contact

Feel free to [reach out to me](https://blog.cloudmentor-scale.com/contact) before any action. Feel free to connect on [Twitter](https://twitter.com/jerome_amon) or [linkedin](https://www.linkedin.com/in/jeromeamon/)