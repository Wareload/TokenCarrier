package proxy

import (
	"github.com/go-playground/validator/v10"
	"os"
)

const prefix = "OIDC_PROXY_"
const defaultLoginPath = "/oidc/login"
const defaultLogoutPath = "/oidc/logout"
const defaultProfilePath = "/oidc/profile"
const defaultCallbackPath = "/oidc/callback"
const defaultBackChannelPath = "/oidc/backchannel/logout"
const defaultUpstreamServerSchema = "http"
const defaultScope = "openid"

type configuration struct {
	ClientId              string `validate:"required"`
	ClientSecret          string `validate:"required"`
	WellKnownURL          string `validate:"required"`
	UpstreamServer        string `validate:"required"`
	UpstreamServerSchema  string `validate:"required"`
	LoginPath             string `validate:"required"`
	LogoutPath            string `validate:"required"`
	ProfilePath           string `validate:"required"`
	CallbackPath          string `validate:"required"`
	BackChannelLogoutPath string `validate:"required"`
	Scope                 string `validate:"required"`
}

func getConfiguration() (configuration, error) {
	config := configuration{
		ClientId:              os.Getenv(prefix + "CLIENT_ID"),
		ClientSecret:          os.Getenv(prefix + "CLIENT_SECRET"),
		WellKnownURL:          os.Getenv(prefix + "WELL_KNOWN_URL"),
		UpstreamServer:        os.Getenv(prefix + "UPSTREAM_SERVER"),
		UpstreamServerSchema:  getEnvWithDefaultValue(os.Getenv(prefix+"UPSTREAM_SERVER_SCHEMA"), defaultUpstreamServerSchema),
		LoginPath:             getEnvWithDefaultValue(os.Getenv(prefix+"LOGIN_PATH"), defaultLoginPath),
		LogoutPath:            getEnvWithDefaultValue(os.Getenv(prefix+"LOGOUT_PATH"), defaultLogoutPath),
		ProfilePath:           getEnvWithDefaultValue(os.Getenv(prefix+"PROFILE_PATH"), defaultProfilePath),
		CallbackPath:          getEnvWithDefaultValue(os.Getenv(prefix+"CALLBACK_PATH"), defaultCallbackPath),
		BackChannelLogoutPath: getEnvWithDefaultValue(os.Getenv(prefix+"BACK_CHANNEL_LOGOUT_PATH"), defaultBackChannelPath),
		Scope:                 defaultScope + " " + os.Getenv(prefix+"SCOPE"),
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(config)
	return config, err
}

func getEnvWithDefaultValue(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		value = defaultValue
	}
	return value
}
