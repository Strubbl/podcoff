package podcoff

// Podcoff represents an instance of the Podcasts coffline application
type Podcoff struct {
	Config   Configuration
	Podcasts []Podcast
	Verbose  bool
	Debug    bool
}

// Init initializes the podcoff, which means it loads the config and the database
func (p *Podcoff) Init(configPath string) error {
	err := (*p).loadConfig(configPath)
	if err != nil {
		return err
	}
	err = (*p).loadPodcasts()
	if err != nil {
		return err
	}
	return nil
}
