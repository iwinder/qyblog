package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/cobra/cmd"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func init() {
	// "github.com/spf13/cobra" 需要 github.com/spf13/viper ,但目前其对1.17不兼容，暂时记录不添加。
	rootCmd.AddCommand(versionCmd)
}

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	// ./___1go_build_gitee_com_windcoder_qingyublog_test_Pflag -v=false
	var version bool
	pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.")
	fmt.Printf("the version : %v", version)
	pflag.Parse()
	//if *help {
	//	pflag.Usage()
	//	return
	//}
	cmd.Execute()
	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
