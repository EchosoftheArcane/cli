//go:build !windows

package findsh

import "github.com/cli/safeexec"

// Find locates the `sh` interpreter on the system.
func Find() (string, error) {
	return safeexec.LookPath("sh")
}
