module hello

go 1.16

require (
	github.com/rs/zerolog v1.23.0
	github.com/runningzyp/calculator v0.0.0
	rsc.io/quote v1.5.2
)

replace github.com/runningzyp/calculator => ../calculator
