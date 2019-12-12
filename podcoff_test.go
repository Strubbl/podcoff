package podcoff_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/strubbl/podcoff"
)

func TestInitNoConfigPath(t *testing.T) {
	p := &podcoff.Podcoff{}
	err := p.Init("")
	if err != nil {
		t.Errorf("Error while Init %v", err)
	}
	defer os.Remove(podcoff.DefaultConfigPath)
	if p.Config.DatabasePath != podcoff.DefaultDatabasePath {
		t.Errorf("config DatabasePath path: Expected %q got %q instead.", podcoff.DefaultDatabasePath, p.Config.DatabasePath)
	}
	if p.Config.DownloadHandler != podcoff.DefaultDownloadHandler {
		t.Errorf("config DownloadHandler path: Expected %q got %q instead.", podcoff.DefaultDownloadHandler, p.Config.DownloadHandler)
	}
	if p.Config.DownloadsPath != podcoff.DefaultDownloadsPath {
		t.Errorf("config DownloadsPath path: Expected %q got %q instead.", podcoff.DefaultDownloadsPath, p.Config.DownloadsPath)
	}
	if p.Config.MetadataPath != podcoff.DefaultMetadataPath {
		t.Errorf("config MetadataPath path: Expected %q got %q instead.", podcoff.DefaultMetadataPath, p.Config.MetadataPath)
	}
}

func TestInitEmptyConfigFile(t *testing.T) {
	emptyConfigFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Error while creating temp emtpy config %v", err)
	}
	defer os.Remove(emptyConfigFile.Name())
	p := &podcoff.Podcoff{}
	err = p.Init(emptyConfigFile.Name())
	if err == nil {
		t.Errorf("No error while Init although we gave an empty config file")
	}
}

func TestInitEmptyConfigStructInFile(t *testing.T) {
	emptyConfigFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Error while creating temp emtpy config %v", err)
	}
	defer os.Remove(emptyConfigFile.Name())
	b, err := json.MarshalIndent(podcoff.Configuration{}, "", "	")
	if err != nil {
		t.Errorf("Error while MarshalIndent of emtpy config struct %v", err)
	}
	err = ioutil.WriteFile(emptyConfigFile.Name(), b, 0644)
	if err != nil {
		t.Errorf("Error while writing temp file with emtpy json marshalled config struct %v", err)
	}

	p := &podcoff.Podcoff{}
	err = p.Init(emptyConfigFile.Name())
	if err == nil {
		t.Errorf("No error while Init although we gave an empty struct config")
	}
}
