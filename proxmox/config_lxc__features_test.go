package proxmox

import (
	"testing"

	"github.com/Telmate/proxmox-api-go/internal/util"
	"github.com/stretchr/testify/require"
)

func Test_RawConfigLXC_Features(t *testing.T) {
	require.Equal(t, &LxcFeatures{
		Unprivileged: &UnprivilegedFeatures{
			CreateDeviceNodes: util.Pointer(true),
			FUSE:              util.Pointer(false),
			KeyCtl:            util.Pointer(false),
			Nesting:           util.Pointer(true)},
	}, RawConfigLXC{
		"features":     string("keyctl=0,mknod=1,nesting=1,fuse=0"),
		"unprivileged": float64(1),
	}.Features())
}
