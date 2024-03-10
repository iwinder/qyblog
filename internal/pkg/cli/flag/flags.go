package flag

import (
	goflag "flag"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	"github.com/spf13/pflag"
	"strings"
)

func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormaliezFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

func WordSepNormaliezFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debugf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
