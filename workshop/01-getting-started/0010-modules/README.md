- Go uses modules to manage projects with external dependencies and internal packages
- If you only have one file and no external dependency, you do not need Go modules
    - `go run filename.go`
- For projects, you usually use
    - `go mod init projectname`
    - `projectname` is a unique identifier, usually the repo url it is hosted under
    - `go run filename.go` to run go mod projects

## Example go.mod
```
module projectname

go 1.22.3
```