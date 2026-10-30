module github.com/viant/xunsafe

go 1.17

require github.com/stretchr/testify v1.8.2

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect; in
)

// AsPointer misreads single-pointer structs on Go >= 1.26; upgrade to v0.11.1 or later.
retract [v0.1.0, v0.11.0]
