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
	"io/ioutil"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/groupssettings/v1"
	"google.golang.org/api/option"
)

type Googler struct {
	GoogleDirectorySrv *admin.Service
	GoogleGroupsSrv    *groupssettings.Service
}

func createDirectoryService(serviceAccountFilePath string, userEmail string) (*admin.Service, error) {
	ctx := context.Background()

	jsonCredentials, err := ioutil.ReadFile(serviceAccountFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(jsonCredentials, admin.AdminDirectoryGroupScope)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	config.Subject = userEmail

	ts := config.TokenSource(ctx)

	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return srv, nil
}

func createGroupsSettingsService(serviceAccountFilePath string) (*groupssettings.Service, error) {
	ctx := context.Background()

	clientOpts := []option.ClientOption{
		option.WithCredentialsFile(serviceAccountFilePath),
		option.WithScopes(groupssettings.AppsGroupsSettingsScope),
	}

	srv, err := groupssettings.NewService(ctx, clientOpts...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return srv, nil
}

func Connect(serviceAccountFilePath string, userEmail string) (*admin.Service, *groupssettings.Service, error) {
	var (
		googleDSrv *admin.Service
		googleGSrv *groupssettings.Service
		err        error
	)

	googleDSrv, err = createDirectoryService(serviceAccountFilePath, userEmail)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	googleGSrv, err = createGroupsSettingsService(serviceAccountFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	return googleDSrv, googleGSrv, nil
}
