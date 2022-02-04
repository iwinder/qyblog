package qy_errors

import (
	"errors"
)

type Aggregate interface {
	error
	Errors() []error
	Is(error) bool
}

func NewAggregate(errList []error) Aggregate {
	if len(errList) == 0 {
		return nil
	}

	var errs []error
	for _, e := range errList {
		if e != nil {
			errs = append(errs, e)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return aggregate(errs)
}

type aggregate []error

func (agg aggregate) Error() string {
	if len(agg) == 0 {
		return ""
	}

	if len(agg) == 1 {
		return agg[0].Error()
	}
	seenerrs := NewString()
	result := ""

	agg.visit(func(err error) bool {
		msg := err.Error()
		if seenerrs.Has(msg) {
			return false
		}
		seenerrs.Insert(msg)
		if len(seenerrs) > 1 {
			result += ","
		}
		result += msg
		return false
	})
	if len(seenerrs) == 1 {
		return result
	}
	return "[" + result + "]"
}

func (agg aggregate) Errors() []error {
	return []error(agg)
}
func (agg aggregate) Is(target error) bool {
	return agg.visit(func(err error) bool {
		return errors.Is(err, target)
	})
}

func (agg aggregate) visit(f func(err error) bool) bool {
	for _, err := range agg {
		switch err := err.(type) {
		case aggregate:
			if match := err.visit(f); match {
				return match
			}
		case Aggregate:
			for _, nestedErr := range err.Errors() {
				if match := f(nestedErr); match {
					return match
				}
			}
		default:
			if match := f(err); match {
				return match
			}
		}
	}
	return false
}
