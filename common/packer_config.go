package common

import "github.com/mitchellh/packer/packer"

// PackerConfig is a struct that contains the configuration keys that
// are sent by packer, properly tagged already so mapstructure can load
// them. Embed this structure into your configuration class to get it.
type PackerConfig struct {
	PackerBuildName   string            `mapstructure:"packer_build_name"`
	PackerBuilderType string            `mapstructure:"packer_builder_type"`
	PackerDebug       bool              `mapstructure:"packer_debug"`
	PackerDebugMode   packer.DebugMode  `mapstructure:"packer_debug_mode"`
	PackerForce       bool              `mapstructure:"packer_force"`
	PackerUserVars    map[string]string `mapstructure:"packer_user_variables"`
}
