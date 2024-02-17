package cmd

import (
	"fmt"
	"os"

	"github.com/SaitoJP/winziper/cmd/components/menu"
	"github.com/SaitoJP/winziper/cmd/components/textinput"
	"github.com/SaitoJP/winziper/cmd/utils/file"
	"github.com/SaitoJP/winziper/cmd/utils/zip"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "winziper",
	Short: "MacからWindwos向けのzipファイル生成",
	Long:  `zip生成する際にMacの制御ファイルなどWindowsに不要なファイルを削除する`,
	Args:  cobra.ExactArgs(1), // 指定した数の引数を受け入れます。
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		var choices = []string{"パスワードなし", "パスワードあり"}
		choice := menu.Run(choices)
		fmt.Println("選択:", choices[choice])

		switch choice {
		case 0:
			result := file.IsDir(path)
			if result == -1 {
				fmt.Println("ファイルが存在しません")
				return
			}

			err := zip.Write(path, true)
			if err != nil {
				panic(err)
			}
		case 1:
			result := file.IsDir(path)
			if result == -1 {
				fmt.Println("ファイルが存在しません")
				return
			}

			pass := textinput.Run("パスワード入力")
			err := zip.WriteEncrypted(path, pass, true)
			if err != nil {
				panic(err)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.winziper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
