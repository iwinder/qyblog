package verflag

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"strconv"
)

// Define some const.
const (
	VersionFalse    versionValue = 0
	VersionTrue     versionValue = 1
	VersionRaw      versionValue = 2
	versionFlagName              = "version"
)
const strRawVersion string = "raw"

type versionValue int

func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}
func (v *versionValue) Get() interface{} {
	return v
}
func (v *versionValue) IsBoolFlag() bool {
	return true
}
func (v *versionValue) Type() string {
	return "version"
}

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit.")

func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	flag.Var(p, name, usage)
	// "--version" will be treated as "--version=true"
	flag.Lookup(name).NoOptDefVal = "true"
}

// Version wraps the VersionVar function.
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

func AddFlag(fs *flag.FlagSet) {
	fs.AddFlag(flag.Lookup(versionFlagName))
}

func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
