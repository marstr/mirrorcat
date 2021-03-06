// Copyright © 2017 Microsoft Corp and Contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var commit string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints information about the instance of MirrorCat that is running.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Built from commit:", commit)
		fmt.Println(runtime.Version(), runtime.GOOS)
		if viper.GetBool("license") {
			fmt.Println()
			fmt.Println(`This Software was written by Microsoft Corp. and Contributors and is Licensed under the MIT license.
The text of which can be found here:
https://github.com/Azure/mirrorcat/blob/master/LICENSE

A portion of this Software was written by the github.com/go-redis/redis Authors and is licensed under the BSD-2 Clause.
The License for this software can be found here:
https://github.com/go-redis/redis/blob/master/LICENSE`)
		}
	},
}

func init() {
	if commit == "" {
		commit = "Unknown"
	}

	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	versionCmd.Flags().BoolP("license", "l", false, "Provide this flag to see license information about this project.")
	viper.BindPFlag("license", versionCmd.Flags().Lookup("license"))
}
