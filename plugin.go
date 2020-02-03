package proxy

import (
	"plugin"
)

// Plugin interface create a proxy to opening a plugin (for mocking)
type Plugin interface {
	Open(string) (*plugin.Plugin, error)
}

// PluginProxy is an empty struct which implements Plugin interface
type PluginProxy struct {
}

// Open opens a plugin using standard go library
func (o *PluginProxy) Open(path string) (*plugin.Plugin, error) {
	return plugin.Open(path)
}
