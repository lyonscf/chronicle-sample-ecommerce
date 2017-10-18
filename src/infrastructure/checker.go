package infrastructure

import (
	"github.com/with-hindsight/chronicle/domain/rule"
	"errors"
)

type StubChecker struct {

}

func (self *StubChecker) Check(invariant interface{}) rule.Rule {

	return rule.NewRule(
		true,
		errors.New("Satisfied!"),
		errors.New("Not satisfied!"),
	)
}
