// Package config stores logic related to configurations like database endpoints, email server and so on

package config

// Config aggregates the config necessary for the application
type Config struct {
	Database struct {
		URI  string
		Name string
	}
	Server struct {
		Address string
		Timeout int
	}
}
