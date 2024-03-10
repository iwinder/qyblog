package app

import (
	cliflag "github.com/iwinder/qyblog/internal/pkg/cli/flag"
)

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

type CompleteableOptions interface {
	Complete() error
}

type PrintableOptions interface {
	String() string
}
