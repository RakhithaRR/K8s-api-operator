// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// removeVersionTag removes version number in a url provided
func removeVersionTag(url string) string {
	regExpString := `\/v[\d.-]*\/?$`
	regExp := regexp.MustCompile(regExpString)
	return regExp.ReplaceAllString(url, "")
}

// isStringArrayContains checks the given text contains in the given arr
func isStringArrayContains(arr []string, text string) bool {
	for _, s := range arr {
		if s == text {
			return true
		}
	}
	return false
}

// getRandFileName returns a file name with suffixing a random number
func getRandFileName(filename string) string {
	fileSplits := strings.SplitN(filename, ".", 2)
	return fmt.Sprintf("%v-%v.%v", fileSplits[0], rand.Intn(10000), fileSplits[1])
}
