package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	format string
	path   string
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "生成文档",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		err = os.Mkdir(path, 0751)
		if err != nil {
			log.Println(err)
		}
		switch format {
		case "manpage":
			header := &doc.GenManHeader{
				Title:   os.Args[0],
				Section: "3",
			}
			err = doc.GenManTree(rootCmd, header, path)
		case "markdown":
			err = doc.GenMarkdownTree(rootCmd, path)
		case "yaml":
			err = doc.GenYamlTree(rootCmd, path)
		case "rest":
			err = doc.GenReSTTree(rootCmd, path)
		default:
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
	docCmd.PersistentFlags().StringVarP(&format, "format", "f", "markdown", "生成文档格式")
	docCmd.PersistentFlags().StringVarP(&path, "path", "p", "docs", "生成文档路径")
}
