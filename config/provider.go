/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	blockdevice "github.com/micgoe/provider-vmware-vra/config/block_device"
	blueprint "github.com/micgoe/provider-vmware-vra/config/blueprint"
	catalogitementitlement "github.com/micgoe/provider-vmware-vra/config/catalog_item"
	catalogsource "github.com/micgoe/provider-vmware-vra/config/catalog_source"
	cloudaccount "github.com/micgoe/provider-vmware-vra/config/cloud_account"
	contentsource "github.com/micgoe/provider-vmware-vra/config/content_source"
	deployment "github.com/micgoe/provider-vmware-vra/config/deployment"
	fabric "github.com/micgoe/provider-vmware-vra/config/fabric"
	flavorprofile "github.com/micgoe/provider-vmware-vra/config/flavor_profile"
	imageprofile "github.com/micgoe/provider-vmware-vra/config/image_profile"
	integration "github.com/micgoe/provider-vmware-vra/config/integration"
	loadbalancer "github.com/micgoe/provider-vmware-vra/config/load_balancer"
	machine "github.com/micgoe/provider-vmware-vra/config/machine"
	network "github.com/micgoe/provider-vmware-vra/config/network"
	project "github.com/micgoe/provider-vmware-vra/config/project"
	storage "github.com/micgoe/provider-vmware-vra/config/storage"
	zone "github.com/micgoe/provider-vmware-vra/config/zone"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/micgoe/provider-vmware-vra"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("micgoe.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		blueprint.Configure,
		deployment.Configure,
		fabric.Configure,
		blockdevice.Configure,
		flavorprofile.Configure,
		imageprofile.Configure,
		storage.Configure,
		catalogsource.Configure,
		catalogitementitlement.Configure,
		cloudaccount.Configure,
		contentsource.Configure,
		integration.Configure,
		loadbalancer.Configure,
		machine.Configure,
		network.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
