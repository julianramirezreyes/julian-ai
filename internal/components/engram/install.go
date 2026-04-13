package engram

import (
	"github.com/julianramirezreyes/julian-ai/internal/installcmd"
	"github.com/julianramirezreyes/julian-ai/internal/model"
	"github.com/julianramirezreyes/julian-ai/internal/system"
)

func InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, model.ComponentEngram)
}
