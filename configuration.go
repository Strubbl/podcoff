package podcoff

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

const DefaultConfigPath = "config.json"

const DefaultDatabasePath = "podcasts.json"
const DefaultDownloadHandler = "wget"
const DefaultDownloadsPath = "downloads"
const DefaultMetadataPath = "metadata"

// Configuration file for podcoff
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
