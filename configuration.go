package podcoff

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

// DefaultConfigPath contains the default file path for the config
const DefaultConfigPath = "config.json"

// DefaultDatabasePath contains the default file path for the database
const DefaultDatabasePath = "podcasts.json"

// DefaultDownloadHandler contains the default executable being used for the item download
const DefaultDownloadHandler = "wget"

// DefaultDownloadsPath contains the default file path where all the downloads are saved
const DefaultDownloadsPath = "downloads"

// DefaultMetadataPath contains the default file path where all feed metadata is saved
const DefaultMetadataPath = "metadata"

// Configuration represents the config file for podcoff
type Configuration struct {
	DatabasePath    string
	DownloadHandler string
	DownloadsPath   string
	MetadataPath    string
}

func getDefaultConfiguration() Configuration {
	var c Configuration
	c.DatabasePath = DefaultDatabasePath
	c.DownloadHandler = DefaultDownloadHandler
	c.DownloadsPath = DefaultDownloadsPath
	c.MetadataPath = DefaultMetadataPath
	return c
}

func (p *Podcoff) loadConfig(configPath string) error {
	if configPath == "" {
		configPath = DefaultConfigPath
		if p.Debug {
			log.Println("no configPath given, using default path:", configPath)
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// config does not exist, create default config and save that
		if p.Debug {
			log.Println("config does not exist, creating a default config")
		}
		(*p).Config = getDefaultConfiguration()
		if p.Debug {
			log.Println("config is set to", p.Config)
		}
		b, err := json.MarshalIndent(p.Config, "", "	")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(configPath, b, 0644)
		if err != nil {
			return err
		}
		return nil
	}
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&p.Config)
	if err != nil {
		return err
	}
	if (Configuration{}) == p.Config {
		return errors.New("Config file found, but marshalling json content of it returned empty config")
	}
	if p.Debug {
		log.Println("Config loaded:", (*p).Config)
	}
	createDirIfNotExists((*p).Config.DownloadsPath)
	createDirIfNotExists((*p).Config.MetadataPath)
	return nil
}

func createDirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}
