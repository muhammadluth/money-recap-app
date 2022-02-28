package dto

type (
	PropertiesService struct {
		AppModeDebug bool           `json:"app_mode_debug"`
		Service      ServiceConfig  `json:"service_config" validate:"dive"`
		DbPostgres   DbConfig       `json:"db_postgres" validate:"dive"`
		Firebase     FirebaseConfig `json:"firebase" validate:"dive"`
		Cors         CorsConfig     `json:"cors" validate:"dive"`
	}

	ServiceConfig struct {
		ServicePort int `json:"service_port" validate:"required,number"`
	}

	DbConfig struct {
		DBHost     string `json:"db_host" validate:"required,min=3,max=50"`
		DBPort     int    `json:"db_port" validate:"required,number"`
		DBName     string `json:"db_name" validate:"required,min=3,max=50"`
		DBUser     string `json:"db_user" validate:"required,min=3,max=50"`
		DBPassword string `json:"db_password" validate:"required,min=3,max=100"`
	}

	FirebaseConfig struct {
		FirebaseAdminSDKConfigName string `json:"firebase_admin_sdk_config_name"  validate:"required,min=3,max=50"`
	}

	CorsConfig struct {
		AllowOrigins string `json:"allow_origins"  validate:"required,min=3,max=50"`
	}
)
