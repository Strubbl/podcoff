package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const defaultConfigPath = "config.json"

const defaultDatabasePath = "podcasts.json"
const defaultDownloadHandler = "wget"
const defaultDownloadsPath = "downloads"
const defaultMetadataPath = "metadata"

// Configuration holds the basic settings for the wallabag-offline application
type Configuration struct {
	CachePath       string
	DatabasePath    string
	DownloadHandler string
	DownloadsPath   string
	MetadataPath    string
}

func getDefaultConfiguration() Configuration {
	var c Configuration
	c.DatabasePath = defaultDatabasePath
	c.DownloadHandler = defaultDownloadHandler
	c.DownloadsPath = defaultDownloadsPath
	c.MetadataPath = defaultMetadataPath
	return c
}

func loadConfig(configPath string) (Configuration, error) {
	var config Configuration

	if configPath == "" {
		configPath = defaultConfigPath
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// config does not exist, create default config and save that
		c := getDefaultConfiguration()
		b, err := json.MarshalIndent(c, "", "	")
		if err != nil {
			return config, err
		}
		err = ioutil.WriteFile(configPath, b, 0644)
		if err != nil {
			return config, err
		}
		return c, nil
	}
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	createDirIfNotExists(config.DownloadsPath)
	createDirIfNotExists(config.MetadataPath)
	return config, nil
}

func createDirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}
