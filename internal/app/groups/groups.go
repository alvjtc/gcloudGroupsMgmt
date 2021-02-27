package groups

import (
	"context"
	"fmt"
	"google.golang.org/api/admin/directory/v1"
)

func GetAllGroups(googleDSrv *admin.Service, domain string) (groupList []*admin.Group, err error) {
	ctx := context.Background()

	err = googleDSrv.Groups.List().Domain(domain).Pages(ctx, func(groups *admin.Groups) (err error) {
		groupList = append(groupList, groups.Groups...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("googleDSrv.Groups.List: %w", err)
	}

	return groupList, nil
}
