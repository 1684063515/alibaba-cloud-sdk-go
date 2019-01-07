package airec

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

// ResultItem is a nested struct in airec response
type ResultItem struct {
	Name           string    `json:"Name" xml:"Name"`
	PerUvBhv       float64   `json:"PerUvBhv" xml:"PerUvBhv"`
	ChargeType     string    `json:"ChargeType" xml:"ChargeType"`
	ErrorLevel     string    `json:"ErrorLevel" xml:"ErrorLevel"`
	ItemType       string    `json:"ItemType" xml:"ItemType"`
	Uv             int       `json:"Uv" xml:"Uv"`
	ClickUser      int       `json:"ClickUser" xml:"ClickUser"`
	BizDate        int       `json:"BizDate" xml:"BizDate"`
	GmtCreate      string    `json:"GmtCreate" xml:"GmtCreate"`
	DataSetVersion string    `json:"DataSetVersion" xml:"DataSetVersion"`
	RegionId       string    `json:"RegionId" xml:"RegionId"`
	ActiveItem     int       `json:"ActiveItem" xml:"ActiveItem"`
	UvCtr          float64   `json:"UvCtr" xml:"UvCtr"`
	InstanceId     string    `json:"InstanceId" xml:"InstanceId"`
	TraceId        string    `json:"TraceId" xml:"TraceId"`
	CommodityCode  string    `json:"CommodityCode" xml:"CommodityCode"`
	MatchInfo      string    `json:"MatchInfo" xml:"MatchInfo"`
	Type           string    `json:"Type" xml:"Type"`
	Click          int       `json:"Click" xml:"Click"`
	Message        string    `json:"Message" xml:"Message"`
	ExpiredTime    string    `json:"ExpiredTime" xml:"ExpiredTime"`
	Timestamp      string    `json:"Timestamp" xml:"Timestamp"`
	Pv             int       `json:"Pv" xml:"Pv"`
	ErrorType      string    `json:"ErrorType" xml:"ErrorType"`
	GmtModified    string    `json:"GmtModified" xml:"GmtModified"`
	Industry       string    `json:"Industry" xml:"Industry"`
	TableName      string    `json:"TableName" xml:"TableName"`
	Scene          string    `json:"Scene" xml:"Scene"`
	TraceInfo      string    `json:"TraceInfo" xml:"TraceInfo"`
	Ctr            float64   `json:"Ctr" xml:"Ctr"`
	PerUvClick     float64   `json:"PerUvClick" xml:"PerUvClick"`
	Position       int       `json:"Position" xml:"Position"`
	Status         string    `json:"Status" xml:"Status"`
	ItemId         int       `json:"ItemId" xml:"ItemId"`
	LockMode       string    `json:"LockMode" xml:"LockMode"`
	Weight         float64   `json:"Weight" xml:"Weight"`
	Meta           Meta      `json:"Meta" xml:"Meta"`
	Parameter      Parameter `json:"Parameter" xml:"Parameter"`
}
