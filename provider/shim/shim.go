package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/squadcast/terraform-provider-squadcast/internal/provider"
)

func New(version string) func() *schema.Provider {
	return provider.New(version)
}
