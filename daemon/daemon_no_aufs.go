// +build exclude_graphdriver_aufs

package daemon

import (
	"github.com/rayzhoull/moby/api/daemon/graphdriver"
)

func migrateIfAufs(driver graphdriver.Driver, root string) error {
	return nil
}
