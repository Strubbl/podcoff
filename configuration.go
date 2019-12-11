package podcoff

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const defaultConfigPath = "config.json"

const defaultDatabasePath = "podcasts.json"
const defaultDownloadHandler = "wget"
const defaultDownloadsPath = "downloads"
const defaultMetadataPath = "metadata"

// Configuration file for podcoff
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

func (p *Podcoff) loadConfig(configPath string) error {
	if configPath == "" {
		configPath = defaultConfigPath
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
