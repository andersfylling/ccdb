package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/andersfylling/ccdb/bot"
	"github.com/andersfylling/ccdb/config"
	"github.com/andersfylling/ccdb/version"
)

const (
	// EnvVarPrefix is the prefix for environment variables
	EnvVarPrefix = "CCDB"
)

var (
	// Version is the current package version
	Version string
)

// conf is a local package variable for access to the config from all cli commands
var conf config.Config

func main() {
	// Set package version for it to be accessible by subpackages
	version.PackageVersion = Version

	// Initialize command-line application
	app := &cli.App{
		Name:    "ccdb",
		Usage:   "Discord bot that displays a crypto currency",
		Version: version.PackageVersion,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				EnvVars: envVarNames("DEBUG"),
				Usage:   "debug mode",
			},
		},
		Before: initApplication,
		Action: runApplication,
	}
	app.Run(os.Args)
}

func envVarName(name string) string {
	return fmt.Sprintf("%s_%s", EnvVarPrefix, strings.ToUpper(name))
}

func envVarNames(names ...string) []string {
	res := make([]string, len(names))
	for i, name := range names {
		res[i] = envVarName(name)
	}
	return res
}

var logFormatter = logrus.TextFormatter{
	FullTimestamp:   true,
	TimestampFormat: "2006-01-02 15:04:05",
}

func initApplication(c *cli.Context) error {
	debug := c.Bool("debug")

	// Configure logger.
	logrus.SetFormatter(&logFormatter)
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Use environment variables as main config
	token := os.Getenv("CCDB_TOKEN")
	if token != "" {
		commandPrefix := os.Getenv("CCDB_COMMANDPREFIX")
		if commandPrefix == "" {
			commandPrefix = "$"
		}

		conf = config.Config{
			Application: config.ApplicationConfig{},
			Discord: config.DiscordConfig{
				Token: token,
			},
			Bot: config.BotConfig{
				CommandPrefix: commandPrefix,
				Status:        "",
			},
		}
	} else {
		logrus.Fatalf("Error could not find environment variable for discord token CCDB_TOKEN")
		return nil
	}

	return nil
}

func runApplication(c *cli.Context) error {
	bot.RunBot(&conf)

	return nil
}
