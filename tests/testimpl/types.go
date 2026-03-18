package testimpl

import "github.com/launchbynttdata/lcaf-component-terratest/types"

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	// Event bus policy module has no additional test config beyond GenericTFModuleConfig.
}
