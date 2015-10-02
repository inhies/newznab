package newznab

import "time"

type Capabilities struct {
	Server struct {
		AppVersion string
		Version    string
		Title      string
		Strapline  string
		Email      string
		URL        string
		Image      string
	}

	Limits struct {
		Default int
		Max     int
	}

	Retention int

	Registration struct {
		Available bool
		Open      bool
	}

	Searching struct {
		Search bool
		Tv     bool
		Movie  bool
		Audio  bool
	}

	Categories []struct {
		Info
		Sub []struct {
			Info
		}
	}

	Groups []struct {
		Info
		Updated time.Time
	}

	Genres []struct {
		Info
	}
}

type Info struct {
	Name        string
	ID          int
	Description string
}
