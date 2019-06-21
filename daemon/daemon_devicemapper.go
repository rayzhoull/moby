// +build !exclude_graphdriver_devicemapper

package daemon

import (
	_ "github.com/rayzhoull/moby/api/daemon/graphdriver/devmapper"
)
