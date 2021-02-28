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
	"reflect"
	"strings"

	"google.golang.org/api/admin/directory/v1"
)

func GetGroupValidArgs() []string {
	return []string{"AdminCreated", "Aliases", "Description", "DirectMembersCount", "Etag", "Id", "Kind", "Name", "NonEditableAliases"}
}

func GetHeaders(args []string) []string {
	var ret []string

	ret = append(ret, "email")
	ret = append(ret, args...)

	return ret
}

func (g *Group) ToSlice(args []string) ([]string, error) {
	ret := make([]string, 0, len(g.Email))

	ret = append(ret, g.Email)

	for _, arg := range args {
		field, err := GetGroupField(g, arg)
		if err != nil {
			return nil, err
		}

		ret = append(ret, field)
	}

	return ret, nil
}

func GetGroupField(item interface{}, fieldName string) (value string, err error) {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return "", fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	findJSONName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}

		return "", fmt.Errorf("provided tag does not define a json tag")
	}

	fieldNames := map[string]int{}

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		name, _ := findJSONName(tag)
		fieldNames[name] = i
	}

	fieldNum, ok := fieldNames[fieldName]
	if !ok {
		return "", fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}

	return v.Field(fieldNum).String(), nil
}

func GetAllGroups(googleDSrv *admin.Service, domain string) (groupList []*Group, err error) {
	ctx := context.Background()

	err = googleDSrv.Groups.List().Domain(domain).Pages(ctx, func(groups *admin.Groups) (err error) {
		for _, g := range groups.Groups {
			groupList = append(groupList, (*Group)(g))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return groupList, nil
}
