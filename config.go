package emailist

import (
	"fmt"
	"os"
	"strconv"
)

// Config is the struct used to configure email sending
type Config struct {
	Host     string // SMTP Hostname (default is localhost)
	Port     int    // SMTP port (deafult is 25)
	User     string // SMTP username
	Password string // SMTP password
	TLS      bool   // Set to use TLS
	SSL      bool   // Set to use SSL
}

// NewConfigWithPrefix makes use of the same environment variables
// but prefixed with <prefix>_.
// eg: A prefix of foo, means varaibles would be named foo_MAIL_HOST, for example.
func NewConfigWithPrefix(prefix string) *Config {
	return newconfig(prefix + "_")
}

// NewConfig returns a new configuration object, initializd from environment
// variables.
func NewConfig() *Config {
	return newconfig("")
}

func newconfig(p string) *Config {
	c := Config{}

	c.Host = os.Getenv(fmt.Sprintf("%sMAIL_HOST", p))
	if c.Host == "" {
		c.Host = "localhost"
	}
	_port := os.Getenv(fmt.Sprintf("%sMAIL_PORT", p))
	if _port == "" {
		_port = "25"
	}
	c.Password = os.Getenv(fmt.Sprintf("%sMAIL_PASSWORD", p))
	c.User = os.Getenv(fmt.Sprintf("%sMAIL_USER", p))
	_ssl := os.Getenv(fmt.Sprintf("%sMAIL_SSL", p))
	_tls := os.Getenv(fmt.Sprintf("%sMAIL_TLS", p))

	if _ssl != "" {
		c.SSL = true
	}

	if _tls != "" {
		c.TLS = true
	}

	// explicit is better than implicit
	i, err := strconv.Atoi(_port)
	if err == nil {
		c.Port = i
	} else {
		c.Port = 0
	}

	return &c
}
