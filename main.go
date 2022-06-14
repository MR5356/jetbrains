package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"jetbrains/utils"
	"net/http"
	"os"
)

func RunWeb(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GetLinks(c.Param("os")))
}

var lsCmd = &cobra.Command{
	Use: "ls",
	Short: `获取Jetbrains产品最新版本的下载地址
		ls 获取所有系统安装包
		ls windows 获取Windows安装包
		ls limux 获取linux安装包
		ls mac 获取mac安装包
		ls mac-m1 获取m1芯片mac安装包`,
	Run: func(cmd *cobra.Command, args []string) {
		var linkItem utils.LinkItemSort
		if len(args) != 1 {
			linkItem = utils.GetLinks("all")
		} else {
			linkItem = utils.GetLinks(args[0])
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"name", "PLATFORM", "size", "version", "build", "date", "link"})
		table.SetBorder(false)
		for _, item := range linkItem {
			table.Append([]string{item.Name, item.PlatFrom, item.Size, item.Version, item.Build, item.Date, item.Link})
		}
		table.Render()
	},
}

var startCmd = &cobra.Command{
	Use: "start",
	Short: `启动web api获取Jetbrains产品最新版本的下载地址
		jetbrains start : 在8999端口启动web API
		jetbrains start 2345 : 在2345(自定义端口)启动web API`,
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.GET("/:os/links", RunWeb)
		if len(args) == 1 {
			r.Run(":" + args[0])
		} else {
			r.Run(":8999")
		}
	},
}

var rootCmd = &cobra.Command{
	Use:   "jetbrains",
	Short: "获取Jetbrains产品最新版本的下载地址",
}

func main() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}
