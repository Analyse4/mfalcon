// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"log"
	"os/exec"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer [file] [dst]",
	Short: "Transfer any files to any node",
	Long:  `transfer simplify command scp and it transfer any files to any node`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		pkPath := viper.GetString("private_key")
		var ip, username, dir string
		para := make([]string, 0)
		para = append(para, "-i", pkPath)
		for _, v := range viper.Get("node").(map[int]interface{}) {
			if v.(Node).Nodename == args[len(args)-1] {
				ip = v.(Node).IP
				username = v.(Node).Username
				dir = v.(Node).Dir
			}
		}
		for _, v := range args[:len(args)-1] {
			para = append(para, v)
		}
		para = append(para, username+"@"+ip+":"+dir)
		//fmt.Println(para)
		c := exec.Command("scp", para...)
		if err := c.Run(); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
