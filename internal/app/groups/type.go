package groups

import (
	"google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/groupssettings/v1"
)

type Googler struct {
	GoogleDirectorySrv *admin.Service
	GoogleGroupsSrv    *groupssettings.Service
}

type Group admin.Group
