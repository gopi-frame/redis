package redis

// ConnectionConfig connection config
type ConnectionConfig struct {
	MasterName       string   `json:"masterName" toml:"masterName" yaml:"masterName" mapstructure:"masterName"`
	Addresses        []string `json:"addresses" toml:"addresses" yaml:"addresses" mapstructure:"addresses"`
	ClientName       string   `json:"clientName" toml:"clientName" yaml:"clientName" mapstructure:"clientName"`
	Database         int      `json:"database" toml:"database" yaml:"database" mapstructure:"database"`
	Username         string   `json:"username" toml:"username" yaml:"username" mapstructure:"username"`
	Password         string   `json:"password" toml:"password" yaml:"password" mapstructure:"password"`
	SentinelUsername string   `json:"sentinelUsername" toml:"sentinelUsername" yaml:"sentinelUsername" mapstructure:"sentinelUsername"`
	SentinelPassword string   `json:"sentinelPassword" toml:"sentinelPassword" yaml:"sentinelPassword" mapstructure:"sentinelPassword"`
}
