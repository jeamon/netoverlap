// NetOverlap is a cli utility that prints to standard output the relation between two CIDRs.
package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/jeamon/netoverlap/app"
)

var ( // Useful build flags.
	BuildTime string
	GitCommit string
	GitTag    string
	Contact   = "https://cloudmentor-scale.com/contact"
)

// main is the entry function which initializes the build flags then parse and
// processes the arguments to display the result & exit code of the operations.
func main() {
	app.Init(&BuildTime, &GitCommit, &GitTag)
	os.Exit(Execute(os.Stdout))
}

// Execute function processes argument(s) passed and output the result along with the exit code.
func Execute(out io.Writer) int {
	// validate commands to display help or version details.
	if len(os.Args) == 2 {
		if os.Args[1] == "version" || os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Fprintf(out, "Version: %s\nGo version: %s\nGit commit: %s\nOS/Arch: %s/%s\nBuilt: %s\nContact: %s\n",
				GitTag, runtime.Version(), GitCommit, runtime.GOOS, runtime.GOARCH, BuildTime, Contact)
			return 0
		}
		if os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Fprintf(out, "%s\n", usage)
			return 0
		}
		fmt.Fprintf(out, "%s %s: Unknown command. Run '%s help' for usage.\n", os.Args[0], os.Args[1], os.Args[0])
		return 1
	}

	// ensure only 2 arguments are provided.
	if len(os.Args) != 3 {
		fmt.Fprintf(out, "Invalid syntax. Run '%s help' for usage.\n", os.Args[0])
		return 1
	}

	// pass network prefixes to check for the overlap status.
	status, err := app.Run(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(out, err)
		return 1
	}

	fmt.Fprintln(out, status)
	return 0
}

const usage = `Usage:
    
    This tool provides the overlap status between two given network prefixes of same type (IPv4 or IPv6)
    in their CIDR notation. It does by evaluating the second network prefix status from the first prefix
	perspective. The result of this checks is either [<subset> or <superset> or <same> or <different>].
    
    netoverlap [version | help | <first-network-prefix> <second-network-prefix>]

    Examples:
	
	$ ./netoverlap 10.0.0.0/20 10.0.2.0/24
	$ ./netoverlap 10.0.0.0/20 10.0.2.0/24
	$ ./netoverlap 10.0.2.0/24 10.0.3.0/24

	$ ./netoverlap ::/0 fe80:c845:ea23::/8
	$ ./netoverlap fd74:5909::/8 fe80::/64
	
	$ ./netoverlap help
	$ ./netoverlap version

`
