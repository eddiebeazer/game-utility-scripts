package main

import (
	"flag"
	"strings"

	"bitbucket.org/thedigitalsages/playfab-rest-api/playfabApi"
)

func main() {
	// parsing command flags
	playfabSecretKey := flag.String("secretKey", "", "API key for PlayFab")
	playfabTitleId := flag.String("titleId", "", "Title Id for PlayFab")
	version := flag.String("version", "", "Version to update to")
	imageName := flag.String("imageName", "", "Name of the new builds docker image")
	imageTag := flag.String("imageTag", "", "tag of the new builds docker image")
	minimumServerCount := flag.Int("minimumServerCount", 0, "Number of servers that will be ran at all times")
	flag.Parse()

	playfabApi.InitClient(*playfabTitleId, *playfabSecretKey, true)

	// setting up keys and value pairs for json
	versionKeyValue := ""
	queueKeyName := ""
	queueKeyValue := ""
	standbyServers := *minimumServerCount
	buildIdKey := ""
	if strings.Contains(*version, "dev") {
		versionKeyValue = strings.ReplaceAll(strings.ReplaceAll(*version, "_", "-"), ".", "_")
		queueKeyName = "DevQueueName"
		buildIdKey = "DevBuildId"
	}
	if strings.Contains(*version, "rel") {
		versionKeyValue = strings.ReplaceAll(strings.ReplaceAll(strings.TrimPrefix(*version, "rel_"), "_", "-"), ".", "_")
		queueKeyName = "BetaQueueName"
		buildIdKey = "BetaBuildId"
	}
	queueKeyValue = "QuickMatch_" + versionKeyValue
	if versionKeyValue == "" || queueKeyValue == "" {
		panic("Post body key blank")
	}

	createBuildJson := &playfabApi.CreateMultiplayerBuildRequest{
		BuildName:                   versionKeyValue,
		MultiplayerServerCountPerVm: 20,
		ContainerFlavor:             "CustomLinux",
		VmSize:                      "Standard_D2as_v4",
		Ports: []playfabApi.Port{
			{Name: "Unreal UDP", Num: "7777", Protocol: "UDP"},
		},
		RegionConfigurations: []playfabApi.BuildRegion{
			{
				MaxServers:     20,
				StandbyServers: standbyServers,
				Region:         "EastUs",
				DynamicStandbySettings: playfabApi.DynamicStandbySettings{
					IsEnabled: true,
				},
			},
		},
		ContainerImageReference: playfabApi.ContainerImageReference{
			ImageName: *imageName,
			ImageTag:  *imageTag,
		},
	}

	buildId := playfabApi.CreateBuildWithCustomContainer(*createBuildJson)

	newBuildTitleData := &playfabApi.TitleData{
		Key:   buildIdKey,
		Value: buildId,
	}
	playfabApi.SetTitleData(*newBuildTitleData)

	createMatchmakeJson := &playfabApi.CreateMatchMakingQueueRequest{
		MatchmakingQueue: playfabApi.MatchmakingQueue{
			BuildId:                 buildId,
			MaxMatchSize:            6,
			MinMatchSize:            2,
			Name:                    queueKeyValue,
			ServerAllocationEnabled: true,
			RegionSelectionRule: playfabApi.RegionSelectionRule{
				Name:       "Region",
				MaxLatency: 250,
				Path:       "Latency",
				Weight:     1.0,
			},
			Teams: []playfabApi.MatchmakingTeam{
				{
					Name:        "team1",
					MinTeamSize: 1,
					MaxTeamSize: 3,
				},
				{
					Name:        "team2",
					MinTeamSize: 1,
					MaxTeamSize: 3,
				},
			},
		},
	}

	playfabApi.CreateMatchmakingQueue(*createMatchmakeJson)

	gameQueueTitleData := &playfabApi.TitleData{
		Key:   queueKeyName,
		Value: queueKeyValue,
	}

	// setting the title data
	playfabApi.SetTitleData(*gameQueueTitleData)
}
