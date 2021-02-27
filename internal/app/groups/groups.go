//Copyright 2021 Álvaro José Teijido Carpente
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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
