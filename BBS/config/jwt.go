package config

type JWT struct {
	SigningKey   string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`       // jwt签名
	RExpiresTime int    `mapstructure:"rexpires-time" json:"rexpires-time" yaml:"rexpires-time"` // R过期时间
	AExpiresTime int    `mapstructure:"aexpires-time" json:"aexpires-time" yaml:"aexpires-time"` // A过期时间
	Issuer       string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                      // 签发者
}
