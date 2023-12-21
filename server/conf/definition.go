package conf

type ServerConfigStruct struct {
	HttpHost      string `mapstructure:"http_host"`
	HttpPort      int    `mapstructure:"http_port"`
	DebugLogFile  string `mapstructure:"debug_log"`
	AccessLogFile string `mapstructure:"access_log"`
}

type SecurityConfigStruct struct {
	JWTSecret string `mapstructure:"jwt_secret"`
	JWTExpire int    `mapstructure:"jwt_expire"`
	OTPUrl    string `mapstructure:"otp_url"`
}

type RedisConfigStruct struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	MainDB   int    `mapstructure:"main_db"`
}
