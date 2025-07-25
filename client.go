// Copyright 2025- The sacloud/nosql-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nosql

import (
	"fmt"
	"net/http"
	"runtime"

	client "github.com/sacloud/api-client-go"
	saht "github.com/sacloud/go-http"
	v1 "github.com/sacloud/nosql-api-go/apis/v1"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/tk1b/api/cloud/1.1"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"nosql-api-go/%s (%s/%s; +https://github.com/sacloud/nosql-api-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	client.DefaultUserAgent,
)

func NewClient(params ...client.ClientParam) (*v1.Client, error) {
	return NewClientWithApiUrl(DefaultAPIRootURL, params...)
}

func NewClientWithApiUrl(apiUrl string, params ...client.ClientParam) (*v1.Client, error) {
	params = append(params, client.WithUserAgent(UserAgent), client.WithOptions(&client.Options{
		RequestCustomizers: []saht.RequestCustomizer{
			func(req *http.Request) error {
				// 文字列を勝手に数値に変換しないようヘッダーで指定
				req.Header.Set("X-Sakura-Bigint-As-Int", "0")
				return nil
			}}}))
	c, err := client.NewClient(apiUrl, params...)
	if err != nil {
		return nil, NewError("NewClientWithApiUrl", err)
	}

	v1Client, err := v1.NewClient(c.ServerURL(), v1.WithClient(c.NewHttpRequestDoer()))
	if err != nil {
		return nil, NewError("NewClientWithApiUrl", err)
	}

	return v1Client, nil
}
