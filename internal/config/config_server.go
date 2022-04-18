package config

import (
	"time"

	"github.com/spf13/viper"

	"github.com/developer-profile/devmetr/internal/datastorage"
	"github.com/developer-profile/devmetr/internal/server"
)

func NewServerConfig(v *viper.Viper) *server.Config {
	v.SetDefault(envServer, DefaultServer)
	v.SetDefault(envStoreInterval, DefaultStoreInterval)
	v.SetDefault(envStoreFile, DefaultStoreFile)
	v.SetDefault(envRestore, DefaultRestore)

	return &server.Config{
		Server: v.GetString(envServer),
		StorageConfig: datastorage.StorageConfig{
			StoreInterval: v.GetDuration(envStoreInterval),
			StoreFile:     v.GetString(envStoreFile),
			Restore:       v.GetBool(envRestore),
			Store:         v.GetString(envStoreFile) != "",
			Synchronized:  v.GetDuration(envStoreInterval) == time.Duration(0),
		},
	}
}

func NewServerConfigWithDefaults(v *viper.Viper, adress string, stroreInterval time.Duration, storeFile string, restore bool) *server.Config {
	v.SetDefault(envServer, adress)
	v.SetDefault(envStoreInterval, stroreInterval)
	v.SetDefault(envStoreFile, storeFile)
	v.SetDefault(envRestore, restore)

	return &server.Config{
		Server: v.GetString(envServer),
		StorageConfig: datastorage.StorageConfig{
			StoreInterval: v.GetDuration(envStoreInterval),
			StoreFile:     v.GetString(envStoreFile),
			Restore:       v.GetBool(envRestore),
			Store:         v.GetString(envStoreFile) != "",
			Synchronized:  v.GetDuration(envStoreInterval) == time.Duration(0),
		},
	}
}
