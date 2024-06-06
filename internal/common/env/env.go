package env

import "os"

const (
	envName = "KONG_CLUSTER_ENV"
)

func GetEnv() string {
	r := os.Getenv(envName)
	return r
}

func IsProduction() bool {
	return GetEnv() == "prod"
}

func InCloud() bool {
	return GetEnv() != ""
}

func InLocal() bool {
	return GetEnv() == ""
}
