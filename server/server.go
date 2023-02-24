package server

import (
	"fmt"
	"os"

	"github.com/gestor-gastos/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "all",
	Short: "All Command Application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Start() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	//Inicializamos la lectura del archivo de configuracion
	viper.SetConfigName("setting")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("No es posible leer el archivo setting %v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(versionMicroServer)
	rootCmd.AddCommand(servicesMicroServer)
}

var versionMicroServer = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		Name := viper.GetString("Microservice.name")
		Version := viper.GetString("Microservice.version")

		fmt.Printf("%s version %s", Name, Version)
	},
}

var servicesMicroServer = &cobra.Command{
	Use: "execute",
	Run: func(cmd *cobra.Command, args []string) {
		api.ExecuteAPI()
	},
}
