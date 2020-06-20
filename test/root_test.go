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

package test

import (
	"github.com/h2cone/lstars/cmd"
	"os/exec"
	"testing"
)

var (
	username = "h2cone"
	language = "null"
	page = 1
	perPage = 30
	sort = ""
	direction = ""
)

func TestListStars(t *testing.T) {
	res := cmd.ListStars(&username, page, perPage, &sort, &direction)
	if len(res) < perPage {
		t.Fail()
	}
}

func TestMain(t *testing.T) {
	cmd := exec.Command("/usr/local/go/bin/go", "run", "../main.go", "-u=" + username, "-l=" + language)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("cmd.Run() failed with %s\n", err)
	}
	t.Logf("combined out:\n%s\n", string(out))
}