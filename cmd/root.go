package cmd

import (
	"fmt"
	"novem/cmd/fsys"
	"novem/cmd/index"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// When called with no args, novem will stat the current
// directory for any files it is tracking, without delving
// into any subdirs(at least not without `-r` flag).
var rootCmd = &cobra.Command{
	Use:   "novem",
	Short: "Dead simple wrapper for managing dotfiles. \n" + flavorText,
	Args:  cobra.NoArgs,
	Run:   statCwd,
}

func statCwd(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		tics.ThrowSys(statCwd, err)
	}

	names, nodes := fsys.StatDir(cwd, &SubDirFlag)
	idx := index.From(names, nodes)

	fmt.Println(idx)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&SubDirFlag, "recursive", "r", false, "traverses subdirectories wherever novem was called")
}

func initConfig() {

	confMap = tics.CobraRootInitBoilerPlate(confMap, true)
	confPath := confMap[confFile]
	viper.SetConfigFile(confPath)
	viper.AutomaticEnv()

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		tics.RunConfPrompts("novem", confMap, defaultSettings)
		tics.ThrowQuiet("")
	}
}
