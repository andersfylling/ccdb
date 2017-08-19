package config

import (
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
)

// Load loads the configuration from the given file path.
func Load(path string, conf *Config) error {
	path = expandUserHomePath(path)

	_, err := toml.DecodeFile(path, &conf)

	return err
}

func expandUserHomePath(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	if path[:2] == "~/" {
		path = strings.Replace(path, "~/", dir, 1)
	}
	return path
}
