package config

type AllConfig struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// Sqlx
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
