package containerd

import "github.com/urfave/cli/v2"

type ContainerdCommandOptions struct {
	Version  string
	ThinPool string
}

func BuildContainerdOptions(cCtx *cli.Context) ContainerdCommandOptions {

	var version = cCtx.String("version")
	var thinPool = cCtx.String("thinpool")

	var newOptions = ContainerdCommandOptions{
		Version:  version,
		ThinPool: thinPool,
	}

	return newOptions
}
