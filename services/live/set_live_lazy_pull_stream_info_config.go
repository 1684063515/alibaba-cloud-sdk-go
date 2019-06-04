package live

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
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetLiveLazyPullStreamInfoConfig invokes the live.SetLiveLazyPullStreamInfoConfig API synchronously
// api document: https://help.aliyun.com/api/live/setlivelazypullstreaminfoconfig.html
func (client *Client) SetLiveLazyPullStreamInfoConfig(request *SetLiveLazyPullStreamInfoConfigRequest) (response *SetLiveLazyPullStreamInfoConfigResponse, err error) {
	response = CreateSetLiveLazyPullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetLiveLazyPullStreamInfoConfigWithChan invokes the live.SetLiveLazyPullStreamInfoConfig API asynchronously
// api document: https://help.aliyun.com/api/live/setlivelazypullstreaminfoconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLiveLazyPullStreamInfoConfigWithChan(request *SetLiveLazyPullStreamInfoConfigRequest) (<-chan *SetLiveLazyPullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *SetLiveLazyPullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLiveLazyPullStreamInfoConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetLiveLazyPullStreamInfoConfigWithCallback invokes the live.SetLiveLazyPullStreamInfoConfig API asynchronously
// api document: https://help.aliyun.com/api/live/setlivelazypullstreaminfoconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLiveLazyPullStreamInfoConfigWithCallback(request *SetLiveLazyPullStreamInfoConfigRequest, callback func(response *SetLiveLazyPullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLiveLazyPullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.SetLiveLazyPullStreamInfoConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetLiveLazyPullStreamInfoConfigRequest is the request struct for api SetLiveLazyPullStreamInfoConfig
type SetLiveLazyPullStreamInfoConfigRequest struct {
	*requests.RpcRequest
	PullArgs       string           `position:"Query" name:"PullArgs"`
	AppName        string           `position:"Query" name:"AppName"`
	PullAuthKey    string           `position:"Query" name:"PullAuthKey"`
	PullAuthType   string           `position:"Query" name:"PullAuthType"`
	DomainName     string           `position:"Query" name:"DomainName"`
	PullDomainName string           `position:"Query" name:"PullDomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	PullAppName    string           `position:"Query" name:"PullAppName"`
	PullProtocol   string           `position:"Query" name:"PullProtocol"`
}

// SetLiveLazyPullStreamInfoConfigResponse is the response struct for api SetLiveLazyPullStreamInfoConfig
type SetLiveLazyPullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLiveLazyPullStreamInfoConfigRequest creates a request to invoke SetLiveLazyPullStreamInfoConfig API
func CreateSetLiveLazyPullStreamInfoConfigRequest() (request *SetLiveLazyPullStreamInfoConfigRequest) {
	request = &SetLiveLazyPullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetLiveLazyPullStreamInfoConfig", "live", "openAPI")
	return
}

// CreateSetLiveLazyPullStreamInfoConfigResponse creates a response to parse from SetLiveLazyPullStreamInfoConfig response
func CreateSetLiveLazyPullStreamInfoConfigResponse() (response *SetLiveLazyPullStreamInfoConfigResponse) {
	response = &SetLiveLazyPullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
