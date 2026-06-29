package parser

// PackageName returns the name of this package.
// It serves as a build-verification sentinel confirming
// the package structure is correctly configured.
func PackageName() string {
	return "parser"
}
