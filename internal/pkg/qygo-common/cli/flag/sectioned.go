package flag

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"strings"
)

type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}

	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		wideFs := pflag.NewFlagSet("", pflag.ExitOnError)
		wideFs.AddFlagSet(fs)

		var az string
		if cols > 24 {
			az = strings.Repeat("z", cols-24)
			wideFs.Int(az, 0, strings.Repeat("z", cols-24))
		}

		var buf bytes.Buffer
		fmt.Fprintf(&buf, "\n%s flags:\n\n%s", strings.ToUpper(name[:1])+name[1:], wideFs.FlagUsagesWrapped(cols))

		if cols > 24 {
			i := strings.Index(buf.String(), az)
			lines := strings.Split(buf.String()[:i], "\n")
			fmt.Fprintf(w, strings.Join(lines[:len(lines)-1], "\n"))
			fmt.Fprintln(w)
		} else {
			fmt.Fprintln(w, buf.String())
		}
	}
}
