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

// CreateDirectoryService Build and returns an Admin SDK Directory service object authorized with
// the service accounts that act on behalf of the given user.
// Args:
//    user_email: The email of the user. Needs permissions to access the Admin APIs.
// Returns:
//    Admin SDK directory service object.
func createDirectoryService(serviceAccountFilePath string, userEmail string) (*admin.Service, error) {
	ctx := context.Background()

	jsonCredentials, err := ioutil.ReadFile(serviceAccountFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(jsonCredentials, admin.AdminDirectoryGroupScope)
	if err != nil {
		return nil, fmt.Errorf("JWTConfigFromJSON: %w", err)
	}
	config.Subject = userEmail

	ts := config.TokenSource(ctx)

	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, fmt.Errorf("NewService: %w", err)
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
		return nil, fmt.Errorf("NewService: %w", err)
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
		return nil, nil, fmt.Errorf("createDirectoryService: %w", err)
	}

	googleGSrv, err = createGroupsSettingsService(serviceAccountFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("createGroupsSettingsService: %w", err)
	}

	return googleDSrv, googleGSrv, nil
}
