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
