package podcoff_test

import (
	"testing"

	"github.com/strubbl/podcoff"
)

func TestInit(t *testing.T) {
	p := &podcoff.Podcoff{}
	err := p.Init("")
	if err != nil {
		t.Errorf("Error while Init")
	}
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
