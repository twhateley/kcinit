// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
    "strings"
    "fmt"
)

// tokenCmd represents the token command
var showCmd = &cobra.Command{
	Use:   "show [server|realm|realm-url|client]",
	Short: "Output configuration information",
	Long: "Output configuration information.  If no argument specified, complete config is dumped.",
	Args: cobra.MaximumNArgs(1),
	Run: show,
}

func show(cmd *cobra.Command, args []string) {

    what := "all"
    if (len(args) == 1) {
        what = args[0]
    }


    if (what == "server") {
        server := getServer()
        fmt.Print(server)

    } else if (what == "realm") {
        realm := getRealm()
        fmt.Print(realm)
    } else if (what == "realm-url") {
        fmt.Print(viper.GetString(REALM_URL))
    } else if (what == "client") {
        fmt.Print(viper.GetString(LOGIN_CLIENT))
    } else if (what == "all") {
        fmt.Println("Server:", getServer())
        fmt.Println("Realm:", getRealm())
        fmt.Println("client", viper.GetString(LOGIN_CLIENT))
    }

}

func getRealm() string {
    url := viper.GetString(REALM_URL)
    if (url != "") {
        idx := strings.LastIndex(url, "/realms")
        if (idx != -1) {
            url = url[idx+8:]
        }
    }
    return url
}

func getServer() string {
    url := viper.GetString(REALM_URL)
    idx := strings.LastIndex(url, "/realms")
    if (idx != -1) {
        url = url[:idx]
    }
    return url
}


func init() {
	rootCmd.AddCommand(showCmd)
}
