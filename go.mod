module main

go 1.14

require (
	elasticsearch v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.2.0 // indirect
	github.com/olivere/elastic/v7 v7.0.19
	google.golang.org/api v0.3.1 // indirect
	kafka v0.0.0-00010101000000-000000000000
)

replace (
	elasticsearch => ./elasticsearch
	kafka => ./kafka
)
