package cmd

import (
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/api"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/arashrasoulzadeh/go-game-engine/worker"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		server := &models.Server{
			Port:     "3000",
			Hostname: "0.0.0.0",
			Engine:   gin.Default(),
		}

		dsn := "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		models.Migrate(db)

		cn := worker.Init()
		go agent.Server(cn)

		api.RunServer(server, db, cn)

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
