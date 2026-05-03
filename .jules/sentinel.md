# Sentinel's Journal - GitHub CLI

This journal tracks critical security learnings for the GitHub CLI codebase.

## 2025-05-15 - [HIGH] Path Searching Vulnerability Mitigation
**Vulnerability:** Use of `os/exec.LookPath` instead of `github.com/cli/safeexec.LookPath`.
**Learning:** In Go, `os/exec.LookPath` can search the current working directory for executables on Windows, which can lead to command injection if the user runs `gh` from a malicious repository.
**Prevention:** Always use `github.com/cli/safeexec.LookPath` to ensure that the search does not include the current directory on Windows, unless explicitly intended.
