module windows

go 1.17

replace github.com/myco/dns v1.0.0 => ./dns
replace github.com/myco/utility v1.0.0 => ./utility

require (
	github.com/gorilla/mux v1.8.0
	github.com/myco/dns v1.0.0
)
