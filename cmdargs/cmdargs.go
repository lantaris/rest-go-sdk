package cmdargs

import (
	"github.com/alexflint/go-arg"
)

// ***********************************************************************
func Parse(Args interface{}) {
	arg.MustParse(Args)
}
