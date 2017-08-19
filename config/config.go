package config

// Config contains all configuration objects.
type Config struct {
	Application *ApplicationConfig
	Discord     *DiscordConfig
	Bot         *BotConfig
}

// ApplicationConfig is the general application configuration.
type ApplicationConfig struct {
}

// DiscordConfig is the discord configuration.
type DiscordConfig struct {
	Token string `toml:"token"`
}

// BotConfig is the bot behavior configuration.
type BotConfig struct {
	// Prefix for commands
	CommandPrefix string `toml:"command_prefix"`
	// Initial status of the bot
	Status string
}
