module github.com/jbuchbinder/mock-data/mock

go 1.15

replace (
	github.com/jbuchbinder/mock-data/data => ../data
	github.com/jbuchbinder/mock-data/types => ../types
)

require (
	github.com/jbuchbinder/mock-data/data v0.0.0-00010101000000-000000000000
	github.com/jbuchbinder/mock-data/types v0.0.0-00010101000000-000000000000
)
