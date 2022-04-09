// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"go-api-gin-swagger/client/client"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.TrinityMetadataAPI, error) {
	hostname := viper.GetString("hostname")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, client.DefaultBasePath, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/json"] = runtime.JSONConsumer()

	r.Producers["application/json"] = runtime.JSONProducer()

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	// add all operation groups
	operationGroupAlbumCmd, err := makeOperationGroupAlbumCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAlbumCmd)

	operationGroupAlbumsCmd, err := makeOperationGroupAlbumsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAlbumsCmd)

	operationGroupExampleCmd, err := makeOperationGroupExampleCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupExampleCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

func makeOperationGroupAlbumCmd() (*cobra.Command, error) {
	operationGroupAlbumCmd := &cobra.Command{
		Use:  "album",
		Long: ``,
	}

	operationGetAlbAlbumsCmd, err := makeOperationAlbumGetAlbAlbumsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAlbumCmd.AddCommand(operationGetAlbAlbumsCmd)

	return operationGroupAlbumCmd, nil
}
func makeOperationGroupAlbumsCmd() (*cobra.Command, error) {
	operationGroupAlbumsCmd := &cobra.Command{
		Use:  "albums",
		Long: ``,
	}

	operationGetAlbAlbumsIDCmd, err := makeOperationAlbumsGetAlbAlbumsIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAlbumsCmd.AddCommand(operationGetAlbAlbumsIDCmd)

	return operationGroupAlbumsCmd, nil
}
func makeOperationGroupExampleCmd() (*cobra.Command, error) {
	operationGroupExampleCmd := &cobra.Command{
		Use:  "example",
		Long: ``,
	}

	operationGetExampleHelloworldCmd, err := makeOperationExampleGetExampleHelloworldCmd()
	if err != nil {
		return nil, err
	}
	operationGroupExampleCmd.AddCommand(operationGetExampleHelloworldCmd)

	return operationGroupExampleCmd, nil
}
