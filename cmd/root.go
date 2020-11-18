// Copyright 2020 huangh https://github.com/h2cone

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var (
	username  string
	language  string
	page      int
	perPage   int
	sort      string
	direction string

	rootCmd = &cobra.Command{
		Use:   "lstars",
		Short: "Lists repositories a user has starred",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				stars := ListStars(&username, page, perPage, &sort, &direction)
				if len(stars) == 0 {
					os.Exit(0)
				}
				printURL(stars, filterByLanguage)
				page++
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "username")
	rootCmd.MarkPersistentFlagRequired("user")
	rootCmd.PersistentFlags().StringVarP(&language, "lang", "l", "", "language")
	rootCmd.PersistentFlags().IntVar(&page, "num", 1, "page num")
	rootCmd.PersistentFlags().IntVar(&perPage, "size", 30, "page size")
	rootCmd.PersistentFlags().StringVar(&sort, "sort", "created", "created or updated")
	rootCmd.PersistentFlags().StringVar(&direction, "direction", "desc", "asc or desc")
}

// Execute exeute root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printURL(stars []map[string]interface{}, filter func(map[string]interface{}) bool) {
	for _, star := range stars {
		if filter(star) {
			fmt.Println(star["html_url"])
		}
	}
}

func filterByLanguage(star map[string]interface{}) bool {
	if len(language) == 0 {
		return true
	}
	return star != nil && ((language == "null" && star["language"] == nil) || language == star["language"])
}

var (
	client = resty.New()
	url    = "https://api.github.com/users/{username}/starred"
)

// ListStars Lists repositories a user has starred"
func ListStars(username *string, page, perPage int, sort, direction *string) []map[string]interface{} {
	res := make([]map[string]interface{}, 0, perPage)

	resp, err := client.R().EnableTrace().
		SetResult(&res).
		SetPathParams(map[string]string{
			"username": *username,
		}).
		SetQueryParams(map[string]string{
			"page":      strconv.Itoa(page),
			"per_page":  strconv.Itoa(perPage),
			"sort":      *sort,
			"direction": *direction,
		}).
		Get(url)
	if err != nil || resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		fmt.Printf("failed to list repositories, err: %v, statusCode: %d\n", err, resp.StatusCode())
		os.Exit(1)
	}
	return res
}
