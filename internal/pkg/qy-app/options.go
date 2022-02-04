package qy_app

import (
	cliflag "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/cli/flag"
)

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

type ConfigurableOptions interface {
	ApplyFlags() []error
}

type CompleteableOptions interface {
	Complete() error
}

type PrintableOptions interface {
	String() string
}
