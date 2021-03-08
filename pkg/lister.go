package pkg

import (
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

type Lister struct{}

func (Lister) GetResourceNamespace() string {
	return "extend-k8s.io"
}

func (Lister) Discover(pluginListCh chan dpm.PluginNameList) {
	var plugins = make(dpm.PluginNameList, 0)

	plugins = append(plugins, "example")

	pluginListCh <- plugins
}

func (Lister) NewPlugin(deviceID string) dpm.PluginInterface {
	return &ExamplePlugin{}
}
