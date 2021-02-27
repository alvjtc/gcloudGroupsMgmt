package main

import (
	"fmt"
	"os"

	"gcloudGroupsMgmt/internal/app/groups"
)

var serviceAccountFilePath = "credential.json"
var domain = "ipartnerga4.net"

func main() {
	fmt.Println("Google Workspace Groups Management Tool")

	googleDSrv, googleGSrv, err := groups.Connect(serviceAccountFilePath, "adminip@ipartnerga4.net")
	if err != nil {
		fmt.Println("connect: ", err)
		os.Exit(1)
	}

	groupList, err := groups.GetAllGroups(googleDSrv, domain)
	if err != nil {
		fmt.Println("GetAllGroups: ", err)
		os.Exit(1)
	}

	for _, g := range groupList {
		fmt.Println(g.Email)
	}

	_, err = googleGSrv.Groups.Get("aaa@example.com").Do()
	if err != nil {
		fmt.Println("Groups.Get: ", err)
		os.Exit(1)
	}
}
