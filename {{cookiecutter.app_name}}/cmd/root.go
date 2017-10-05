{% if cookiecutter.license == "Apache" -%}
// Copyright © {{cookiecutter.year}} {{cookiecutter.full_name}} <{{cookiecutter.email}}>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

{% endif -%}
package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
{%- if cookiecutter.use_viper == "y" %}
	"github.com/spf13/viper"
{%- endif %}
)

var (
{%- if cookiecutter.use_viper == "y" %}
	cfgFile     string
{%- endif %}
	showVersion bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "{{cookiecutter.app_name}}",
	Short: "{{cookiecutter.short_description}}",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("{{cookiecutter.app_name}} %s%s\n", Version, VersionPrerelease)
			return
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
{%- if cookiecutter.use_viper == "y" %}
	cobra.OnInitialize(initConfig)
	{% endif -%}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
{%- if cookiecutter.use_viper == "y" %}
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.{{cookiecutter.app_name}}.yaml)")
{%- endif %}
}
{%- if cookiecutter.use_viper == "y" %}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(home)
			os.Exit(1)
		}

		// Search config in home directory with name ".{{cookiecutter.app_name}}" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".{{cookiecutter.app_name}}")
	}

	viper.SetEnvPrefix("{{cookiecutter.app_name}}")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
{%- endif %}
