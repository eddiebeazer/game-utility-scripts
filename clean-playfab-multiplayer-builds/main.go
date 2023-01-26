package main

import (
	"flag"
	"reflect"

	"bitbucket.org/thedigitalsages/playfab-rest-api/playfabApi"
)

/*
1. get current version of queue and build
2. loop through all match queues and buiid ids. If they are not in the list for current queue or build, delete queue
3. Foir builds, if delete if servers = 0, if servers dont equal zero set standby to 0
4.
*/
func main() {
	// parsing command flags
	playfabSecretKey := flag.String("secretKey", "", "API key for PlayFab")
	playfabTitleId := flag.String("titleId", "", "Title Id for PlayFab")
	flag.Parse()

	playfabApi.InitClient(*playfabTitleId, *playfabSecretKey, true)

	var itemsToSave []string

	getTitleDataJson := &playfabApi.GetTitleDataRequest{
		Keys: []string{
			"BetaQueueName",
			"BetaBuildId",
			"DevBuildId",
			"DevQueueName",
			"LiveQueueName",
			"LiveBuildId",
		},
	}

	titleDataResp := playfabApi.GetTitleData(*getTitleDataJson)

	// loop through object without knowing keys
	v := reflect.ValueOf(titleDataResp.Data.Data)
	for i := 0; i < v.NumField(); i++ {
		// fmt.Println(v.Type().Field(i).Name)  GIVES KEY NAME v.Interface().(float64)
		itemsToSave = append(itemsToSave, v.Field(i).Interface().(string))
	}

	queueList := playfabApi.ListMatchmakingQueues()

	for _, queue := range queueList {
		shouldSaveQueue := false
		for _, queueName := range itemsToSave {
			if queue.Name == queueName {
				shouldSaveQueue = true
			}
		}
		if shouldSaveQueue == false {
			removeQueueJson := &playfabApi.RemoveMatchmakingQueueRequest{
				QueueName: queue.Name,
			}
			playfabApi.RemoveMatchmakingQueue(*removeQueueJson)
		}
	}

	buildList := playfabApi.ListMultiplayerBuilds()

	for _, build := range buildList {
		shouldSaveBuild := false
		activeRegion := false
		for _, buildName := range itemsToSave {
			if build.BuildId == buildName {
				shouldSaveBuild = true
			}
		}
		if shouldSaveBuild == true {
			continue
		}
		var regions []playfabApi.BuildRegion
		for _, region := range build.RegionConfigurations {
			if region.StandbyServers > 0 {
				activeRegion = true
				region.StandbyServers = 0
				region.DynamicStandbySettings.IsEnabled = false
			}
			regions = append(regions, region)
		}
		if activeRegion {
			updateReq := &playfabApi.UpdateMultiplayerBuildRegionsRequest{
				BuildId:      build.BuildId,
				BuildRegions: regions,
			}
			playfabApi.UpdateMultiplayerBuildRegions(*updateReq)
		} else {
			deleteReq := &playfabApi.DeleteMultiplayerBuildRequest{
				BuildId: build.BuildId,
			}
			playfabApi.DeleteMultiplayerBuild(*deleteReq)
		}
	}
}
