package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/dcalsky/kong_service_demo/internal/common/logs"
)

func getClusterEnv() string {
	envName := os.Getenv("CLUSTER_ENV")
	if envName == "" {
		return "-"
	}
	return envName
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func unmarshalConfDir(confDir string, target interface{}) error {
	if configDir == "" {
		return errors.New("require a config dir")
	}
	paths := []string{
		filepath.Join(confDir, "base.yaml"),
	}
	clusterEnv := getClusterEnv()
	if clusterEnv != "-" {
		paths = append(paths, filepath.Join(confDir, fmt.Sprintf("%s.yaml", clusterEnv)))
	}
	ctx := context.Background()
	for _, path := range paths {
		if !fileExist(path) {
			logs.Infof(ctx, "[Conf] %s doesn't exist", path)
			continue
		}
		logs.Infof(ctx, "[Conf] use %s", path)
		fileBody, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(fileBody, target); err != nil {
			return err
		}
	}
	return nil
}
