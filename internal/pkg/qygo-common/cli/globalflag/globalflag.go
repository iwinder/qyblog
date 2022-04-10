package globalflag

import (
	"fmt"
	"github.com/spf13/pflag"
)

func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}
