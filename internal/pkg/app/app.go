package app

import (
	"fmt"
	"github.com/fatih/color"
	cliflag "github.com/iwinder/qyblog/internal/pkg/cli/flag"
	"github.com/iwinder/qyblog/internal/pkg/cli/globalflag"
	"github.com/iwinder/qyblog/internal/pkg/errors"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	"github.com/iwinder/qyblog/internal/pkg/term"
	"github.com/iwinder/qyblog/internal/pkg/version"
	"github.com/iwinder/qyblog/internal/pkg/version/verflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	progressMessage = color.GreenString("==>")

	usageTemplate = fmt.Sprintf(`%s{{if .Runnable}}
  %s{{end}}{{if .HasAvailableSubCommands}}
  %s{{end}}{{if gt (len .Aliases) 0}}

%s
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

%s{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  %s {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

%s{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "%s --help" for more information about a command.{{end}}
`,
		color.CyanString("Usage:"),
		color.GreenString("{{.UseLine}}"),
		color.GreenString("{{.CommandPath}} [command]"),
		color.CyanString("Aliases:"),
		color.CyanString("Examples:"),
		color.CyanString("Available Commands:"),
		color.GreenString("{{rpad .Name .NamePadding }}"),
		color.CyanString("Flags:"),
		color.CyanString("Global Flags:"),
		color.CyanString("Additional help topics:"),
		color.GreenString("{{.CommandPath}} [command]"),
	)
)

const flagNameGlobal = "global"

type App struct {
	basename  string
	name      string
	desc      string
	options   CliOptions
	runFunc   RunFunc
	silence   bool
	noVersion bool
	noConfig  bool
	commands  []*Command
	args      cobra.PositionalArgs
	cmd       *cobra.Command
}

type RunFunc func(basename string) error

// Option 将 func(*App) 命名为 Option，方便执行多个方法初始化 App
type Option func(*App)

func NewApp(name string, basename string, opts ...Option) *App {
	a := &App{name: name, basename: basename}

	for _, o := range opts {
		o(a)
	}

	return a
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           FormatBaseName(a.basename),
		Short:         a.name,
		Long:          a.desc,
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	cliflag.InitFlags(cmd.Flags())

	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.getCobraCommand())
		}
	}

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	var namedFlagSets cliflag.NamedFlagSets

	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)

		}
	}
	if !a.noVersion {
		verflag.AddFlag(namedFlagSets.FlagSet(flagNameGlobal))
	}
	if !a.noConfig {
		addConfigFlag(a.basename, namedFlagSets.FlagSet(flagNameGlobal))
	}
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet(flagNameGlobal), cmd.Name())
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet(flagNameGlobal))

	addCmdTemplate(&cmd, namedFlagSets)
	a.cmd = &cmd
}
func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}
func WithDesc(desc string) Option {
	return func(a *App) {
		a.desc = desc
	}
}
func WithNoVersion() Option {
	return func(a *App) {
		a.noVersion = true
	}
}
func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Println("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func (a *App) Command() *cobra.Command {
	return a.cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	printWorkingDir()
	cliflag.PrintFlags(cmd.Flags())
	if !a.noVersion {
		verflag.PrintAndExitIfRequested()
	}
	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}
	if !a.silence {
		log.Infof("%v Starting %s ...", progressMessage, a.name)
		if !a.noVersion {
			log.Infof("%v Version: `%s`", progressMessage, version.Get().ToJSON())
		}
		if !a.noConfig {
			log.Infof("%v Config file used: `%s`", progressMessage, viper.ConfigFileUsed())
		}
	}
	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}
	// run application
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

func (a *App) applyOptionRules() error {
	if completableOptions, ok := a.options.(CompleteableOptions); ok {
		if err := completableOptions.Complete(); err != nil {
			return err
		}
	}

	if errs := a.options.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		log.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}

func printWorkingDir() {
	wd, _ := os.Getwd()
	log.Infof("%v WorkingDir: %s", progressMessage, wd)
}
func addCmdTemplate(cmd *cobra.Command, nameFlagSets cliflag.NamedFlagSets) {
	usageFmt := "Usage:\n %s\n"
	clos, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStdout(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), nameFlagSets, clos)
		return nil
	})

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), nameFlagSets, clos)
	})

}
