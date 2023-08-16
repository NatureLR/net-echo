package cmd

import (
	"fmt"
	netecho "github/naturelr/net-echo/pkg/net-echo"
	"github/naturelr/net-echo/pkg/versions"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "net-echo",
	Short: versions.ShortDescribe,
	Long:  versions.LongDescribe,
	Run: func(cmd *cobra.Command, Args []string) {
		netecho.Run()
	},
	Version: versions.Strings(),
}

// Execute 将所有的子命令加入到根命令并设置适当的flag
// 这是main.main()调用的,只调用一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//在这里，您将定义标志和配置设置。Cobra支持持久性标志，如果在这里定义的话，这里的配置是全局的。
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "config file")
}

// initConfig 读取配置文件和环境变量
func initConfig() {
	self := filepath.Base(os.Args[0])
	if cfgFile != "" {
		// 从flag读取文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 寻找home目录
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
		}
		//用户配置目录
		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			fmt.Println(err)
		}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)                                        //家目录
		viper.AddConfigPath(".")                                         //当前目录
		viper.AddConfigPath(userConfigDir)                               //用户配置目录
		viper.AddConfigPath(fmt.Sprint(filepath.Join("/", "etc", self))) //etc目录下程序的名字下的config.yaml
	}

	// 读取环境环境变量以为程序名字大写开头
	prefix := strings.ToUpper(self)
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	// 如果找到一个配置文件就读取它
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
