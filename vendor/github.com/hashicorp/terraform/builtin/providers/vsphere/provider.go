package vsphere

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_USER", nil),
				Description: "The user name for vSphere API operations.",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_PASSWORD", nil),
				Description: "The user password for vSphere API operations.",
			},

			"vsphere_server": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_SERVER", nil),
				Description: "The vSphere Server name for vSphere API operations.",
			},
			"allow_unverified_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_ALLOW_UNVERIFIED_SSL", false),
				Description: "If set, VMware vSphere client will permit unverifiable SSL certificates.",
			},
			"vcenter_server": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_VCENTER", nil),
				Deprecated:  "This field has been renamed to vsphere_server.",
			},
			"client_debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_CLIENT_DEBUG", false),
				Description: "govomomi debug",
			},
			"client_debug_path_run": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_CLIENT_DEBUG_PATH_RUN", ""),
				Description: "govomomi debug path for a single run",
			},
			"client_debug_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSPHERE_CLIENT_DEBUG_PATH", ""),
				Description: "govomomi debug path for debug",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"vsphere_file":            resourceVSphereFile(),
			"vsphere_folder":          resourceVSphereFolder(),
			"vsphere_virtual_disk":    resourceVSphereVirtualDisk(),
			"vsphere_virtual_machine": resourceVSphereVirtualMachine(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// Handle backcompat support for vcenter_server; once that is removed,
	// vsphere_server can just become a Required field that is referenced inline
	// in Config below.
	server := d.Get("vsphere_server").(string)

	if server == "" {
		server = d.Get("vcenter_server").(string)
	}

	if server == "" {
		return nil, fmt.Errorf(
			"One of vsphere_server or [deprecated] vcenter_server must be provided.")
	}

	config := Config{
		User:          d.Get("user").(string),
		Password:      d.Get("password").(string),
		InsecureFlag:  d.Get("allow_unverified_ssl").(bool),
		VSphereServer: server,
		Debug:         d.Get("client_debug").(bool),
		DebugPathRun:  d.Get("client_debug_path_run").(string),
		DebugPath:     d.Get("client_debug_path").(string),
	}

	return config.Client()
}
