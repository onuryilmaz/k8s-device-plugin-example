package main

import (
	"flag"

	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	"github.com/onuryilmaz/k8s-device-plugin-example/pkg"
)

func main() {
	flag.Parse()
	manager := dpm.NewManager(pkg.Lister{})
	manager.Run()
}
