package vsphere

// Properties describes know relation to properties to related objects and properties
var Properties = map[string]map[string][]string{
	"datastore": {
		"Datastore":      {"name"},
		"VirtualMachine": {"datastore"},
	},
	"urls": {
		"Datastore": {"summary.url"},
	},
	"host": {
		"HostSystem":     {"name", "parent"},
		"VirtualMachine": {"name", "runtime.host"},
	},
	"cluster": {
		"ClusterComputeResource": {"name"},
	},
	"network": {
		"DistributedVirtualPortgroup": {"name"},
		"Network":                     {"name"},
		"VirtualMachine":              {"network"},
	},
	"resourcepool": {
		"ResourcePool": {"name", "parent", "vm"},
	},
	"folder": {
		"Folder":         {"name", "parent"},
		"VirtualMachine": {"parent"},
	},
	"tags": {
		"VirtualMachine": {"tag"},
		"HostSystem":     {"tag"},
	},
	"numcpu": {
		"VirtualMachine": {"summary.config.numCpu"},
	},
	"memorysizemb": {
		"VirtualMachine": {"summary.config.memorySizeMB"},
	},
	"disks": {
		"VirtualMachine": {"guest.disk"},
	},
}

// PropertiesSections represent the mapping of attributes to sections in the cache
var PropertiesSections = map[string]string{
	"summary.url":                 "urls",
	"name":                        "names",
	"datastore":                   "datastores",
	"network":                     "networks",
	"runtime.host":                "hosts",
	"parent":                      "parents",
	"vm":                          "vms",
	"tag":                         "tags",
	"summary.config.numCpu":       "cpus",
	"summary.config.memorySizeMB": "memories",
	"guest.disk":                  "disks",
}
