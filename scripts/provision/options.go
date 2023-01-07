package provision

// The struct with all possible bootstrap options
type ProvisionOptions struct {
	SkipAll            bool   // Autoapprove all prompts
	ThinPool           string // Name of thinpool (default: flintlock[-dev])
	Disk               string // Disk to use for DirectLVM thinpool (ignored for dev)
	GRPCEndpoint       string // Address for gRPC server (default: 0.0.0.0:9090)
	ParentInterface    string // Interface of the default route of the host
	Development        bool   // Set up development environment (uses Loop thinpools)
	DryRun             bool   // Set up development environment (uses Loop thinpools)
	FirecrackerVersion string // Firecracker version to use
	FlintlockVersion   string // Flintlock Version to use
	ContainerdVersion  string // ContainerD Version to use
}
