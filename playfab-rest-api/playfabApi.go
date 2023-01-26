package playfabApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type TitleData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type EntityCredential struct {
	Data struct {
		EntityToken string `json:"EntityToken"`
	} `json:"data"`
	Code int `json:"code"`
}

type GetTitleDataResponse struct {
	Data struct {
		Data struct {
			BetaQueueName   string `json:"BetaQueueName"`
			BetaBuildId     string `json:"BetaBuildId"`
			BetaGameVersion string `json:"BetaGameVersion"`
			DevBuildId      string `json:"DevBuildId"`
			DevGameVersion  string `json:"DevGameVersion"`
			DevQueueName    string `json:"DevQueueName"`
			LiveQueueName   string `json:"LiveQueueName"`
			LiveBuildId     string `json:"LiveBuildId"`
			LiveGameVersion string `json:"LiveGameVersion"`
		} `json:"Data"`
	} `json:"data"`
	Code int `json:"code"`
}

type CreateMultiplayerBuildResponse struct {
	Data struct {
		BuildId string `json:"BuildId"`
	} `json:"data"`
	Code int `json:"code"`
}

type CreateMatchMakingQueueResponse struct {
	Code int `json:"code"`
}

type ListMatchmakingQueuesResult struct {
	Data struct {
		MatchmakingQueues []MatchmakingQueue `json:"MatchMakingQueues"`
	} `json:"data"`
	Code int `json:"code"`
}

type CreateMatchMakingQueueRequest struct {
	MatchmakingQueue MatchmakingQueue `json:"MatchmakingQueue"`
}

type GetTitleDataRequest struct {
	Keys []string `json:"Keys"`
}

type MatchmakingQueue struct {
	BuildId                 string              `json:"BuildId"`
	MaxMatchSize            int                 `json:"MaxMatchSize"`
	MinMatchSize            int                 `json:"MinMatchSize"`
	Name                    string              `json:"Name"`
	ServerAllocationEnabled bool                `json:"ServerAllocationEnabled"`
	RegionSelectionRule     RegionSelectionRule `json:"RegionSelectionRule"`
	Teams                   []MatchmakingTeam   `json:"Teams"`
}

type MatchmakingTeam struct {
	Name        string `json:"Name"`
	MinTeamSize int    `json:"MinTeamSize"`
	MaxTeamSize int    `json:"MaxTeamSize"`
}

type RegionSelectionRule struct {
	Name       string  `json:"Name"`
	MaxLatency int     `json:"MaxLatency"`
	Path       string  `json:"Path"`
	Weight     float32 `json:"Weight"`
}

type RemoveMatchmakingQueueRequest struct {
	QueueName string `json:"QueueName"`
}

type CreateMultiplayerBuildRequest struct {
	BuildName                   string                  `json:"BuildName"`
	MultiplayerServerCountPerVm int                     `json:"MultiplayerServerCountPerVm"`
	Ports                       []Port                  `json:"Ports"`
	RegionConfigurations        []BuildRegion           `json:"RegionConfigurations"`
	ContainerFlavor             string                  `json:"ContainerFlavor"`
	ContainerImageReference     ContainerImageReference `json:"ContainerImageReference"`
	VmSize                      string                  `json:"VmSize"`
}

type ContainerImageReference struct {
	ImageName string `json:"ImageName"`
	ImageTag  string `json:"Tag"`
}

type Port struct {
	Name     string `json:"Name"`
	Num      string `json:"Num"`
	Protocol string `json:"Protocol"`
}

type BuildRegion struct {
	MaxServers             int                    `json:"MaxServers"`
	DynamicStandbySettings DynamicStandbySettings `json:"DynamicStandbySettings"`
	Region                 string                 `json:"Region"`
	StandbyServers         int                    `json:"StandbyServers"`
}

type DynamicStandbySettings struct {
	IsEnabled bool `json:"IsEnabled"`
}

type ListMultiplayerBuildsRequest struct {
	PageSize int `json:"PageSize"`
}

type ListMultiplayerBuildsResponse struct {
	Data struct {
		BuildSummaries []BuildSummary `json:"BuildSummaries"`
	} `json:"data"`
	Code int `json:"code"`
}

type BuildSummary struct {
	BuildName            string        `json:"BuildName"`
	BuildId              string        `json:"BuildId"`
	RegionConfigurations []BuildRegion `json:"RegionConfigurations"`
}

type UpdateMultiplayerBuildRegionsRequest struct {
	BuildId      string        `json:BuildId`
	BuildRegions []BuildRegion `json:"BuildRegions"`
}

type DeleteMultiplayerBuildRequest struct {
	BuildId string `json:BuildId`
}

var baseApiUrl, titleId, secretKey, entityKey string
var client *http.Client

func InitClient(playFabTitleId string, playfabSecretKey string, fetchEntityToken bool) {
	// creating a http client and performing the post request to update title data
	titleId = playFabTitleId
	secretKey = playfabSecretKey
	baseApiUrl = "https://" + titleId + ".playfabapi.com"
	client = &http.Client{Timeout: 30 * time.Second}

	if fetchEntityToken {
		GetEntityToken()
	}
}

func SetTitleData(reqBody TitleData) {
	json_data, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	doPost("/Admin/SetTitleData", json_data, true, false)
}

func GetTitleData(reqBody GetTitleDataRequest) GetTitleDataResponse {
	json_data, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	data := doPost("/Admin/GetTitleData", json_data, true, false)

	var response GetTitleDataResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func GetEntityToken() {
	data := doPost("/Authentication/GetEntityToken", []byte{}, true, false)

	var response EntityCredential
	err := json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}
	entityKey = response.Data.EntityToken
}

func ListMatchmakingQueues() []MatchmakingQueue {
	data := doPost("/Match/ListMatchmakingQueues", []byte{}, false, true)
	var response ListMatchmakingQueuesResult
	err := json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}
	return response.Data.MatchmakingQueues
}

func RemoveMatchmakingQueue(reqBody RemoveMatchmakingQueueRequest) {
	json_data, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	doPost("/Match/RemoveMatchmakingQueue", json_data, false, true)
}

func CreateBuildWithCustomContainer(reqBody CreateMultiplayerBuildRequest) string {
	json_data, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	data := doPost("/MultiplayerServer/CreateBuildWithCustomContainer", json_data, false, true)

	var response CreateMultiplayerBuildResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response.Data.BuildId
}

func ListMultiplayerBuilds() []BuildSummary {
	req := &ListMultiplayerBuildsRequest{
		PageSize: 50,
	}
	json_data, err := json.MarshalIndent(*req, "", "\t")
	if err != nil {
		panic(err)
	}

	data := doPost("/MultiplayerServer/ListBuildSummariesV2", json_data, false, true)

	var response ListMultiplayerBuildsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response.Data.BuildSummaries
}

func UpdateMultiplayerBuildRegions(req UpdateMultiplayerBuildRegionsRequest) {
	json_data, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		panic(err)
	}

	doPost("/MultiplayerServer/UpdateBuildRegions", json_data, false, true)
}

func DeleteMultiplayerBuild(req DeleteMultiplayerBuildRequest) {
	json_data, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		panic(err)
	}

	doPost("/MultiplayerServer/DeleteBuild", json_data, false, true)
}

func CreateMatchmakingQueue(reqBody CreateMatchMakingQueueRequest) {
	json_data, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	doPost("/Match/SetMatchmakingQueue", json_data, false, true)
}

// executes a post request and returns raw json
func doPost(url string, json []byte, addSecretHeader bool, addEntityHeader bool) []byte {

	req, err := http.NewRequest("POST", baseApiUrl+url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	if addSecretHeader {
		req.Header.Add("X-SecretKey", secretKey)
	}
	if addEntityHeader {
		req.Header.Add("X-EntityToken", entityKey)
	}

	// executing request and reading the result into a []byte
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}

func debugJson(data []byte) {
	fmt.Println(string(data))
}
