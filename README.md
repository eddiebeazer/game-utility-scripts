# Get Matchmaking queue

https://docs.microsoft.com/en-us/rest/api/playfab/multiplayer/matchmaking-admin/set-matchmaking-queue?view=playfab-rest#matchmakingqueueconfig

Content-Type  application/json

X-EntityToken

{
  "MatchmakingQueue": {
    "Name": "ExampleQueueName",
    "MinMatchSize": 7,
    "MaxMatchSize": 8,
    "MaxTicketSize": 2,
    "ServerAllocationEnabled": true,
    "BuildId": "065a3208-39af-4691-8794-5f774c367ac2",
    "DifferenceRules": [
      {
        "Difference": 10,
        "MergeFunction": "Average",
        "DefaultAttributeValue": 0,
        "LinearExpansion": {
          "Delta": 0.5,
          "Limit": 20,
          "SecondsBetweenExpansions": 5
        },
        "Attribute": {
          "Path": "ExampleAttributeDifference",
          "Source": "User"
        },
        "AttributeNotSpecifiedBehavior": "UseDefault",
        "Weight": 1,
        "Name": "ExampleNameDifference",
        "SecondsUntilOptional": 50
      }
    ],
    "StringEqualityRules": [
      {
        "DefaultAttributeValue": "ExampleDefault",
        "Expansion": {
          "EnabledOverrides": [
            false,
            true,
            true,
            false
          ],
          "SecondsBetweenExpansions": 5
        },
        "Attribute": {
          "Path": "ExampleAttributeStringEquality",
          "Source": "User"
        },
        "AttributeNotSpecifiedBehavior": "UseDefault",
        "Weight": 1,
        "Name": "ExampleNameStringEquality",
        "SecondsUntilOptional": 50
      }
    ],
    "MatchTotalRules": [
      {
        "Attribute": {
          "Path": "ExampleAttribute",
          "Source": "User"
        },
        "Min": 2,
        "Max": 0,
        "Weight": 1,
        "Expansion": {
          "MaxOverrides": [
            {
              "Value": 4
            },
            {
              "Value": 4
            },
            null
          ],
          "SecondsBetweenExpansions": 5
        },
        "Name": "MatchTotalRule",
        "SecondsUntilOptional": 50
      }
    ],
    "SetIntersectionRules": [
      {
        "DefaultAttributeValue": [
          "a",
          "b",
          "c"
        ],
        "MinIntersectionSize": 2,
        "LinearExpansion": {
          "Delta": 2,
          "SecondsBetweenExpansions": 5
        },
        "Attribute": {
          "Path": "ExampleAttributeSetIntersectionRule",
          "Source": "User"
        },
        "AttributeNotSpecifiedBehavior": "UseDefault",
        "Weight": 1,
        "Name": "ExampleNameSetIntersectionRule",
        "SecondsUntilOptional": 50
      }
    ],
    "RegionSelectionRule": {
      "MaxLatency": 250,
      "Path": "Latencies",
      "LinearExpansion": {
        "Delta": 10,
        "Limit": 300,
        "SecondsBetweenExpansions": 5
      },
      "Weight": 1,
      "Name": "RegionSelectionRule",
      "SecondsUntilOptional": 50
    },
    "TeamSizeBalanceRule": {
      "Difference": 1,
      "LinearExpansion": {
        "Delta": 1,
        "Limit": 5,
        "SecondsBetweenExpansions": 5
      },
      "Name": "TeamSizeBalanceRule",
      "SecondsUntilOptional": 50
    },
    "TeamDifferenceRules": [
      {
        "Attribute": {
          "Path": "ExampleAttribute",
          "Source": "User"
        },
        "Difference": 2,
        "DefaultAttributeValue": 0,
        "LinearExpansion": {
          "Delta": 1,
          "Limit": 5,
          "SecondsBetweenExpansions": 5
        },
        "Name": "TeamDifferenceRule",
        "SecondsUntilOptional": 50
      }
    ],
    "TeamTicketSizeSimilarityRule": {
      "Name": "TeamTicketSizeSimilarityRule",
      "SecondsUntilOptional": 180
    },
    "Teams": [
      {
        "Name": "monster",
        "MinTeamSize": 1,
        "MaxTeamSize": 1
      },
      {
        "Name": "hunters",
        "MinTeamSize": 4,
        "MaxTeamSize": 8
      }
    ],
    "StatisticsVisibilityToPlayers": {
      "ShowNumberOfPlayersMatching": true,
      "ShowTimeToMatch": true
    }
  }
}

# Remove matchmaking queue

https://docs.microsoft.com/en-us/rest/api/playfab/multiplayer/matchmaking-admin/remove-matchmaking-queue?view=playfab-rest

{
  "QueueName": "custom123"
}

# Multiplayer Server - Create Build With Custom Container

https://docs.microsoft.com/en-us/rest/api/playfab/multiplayer/multiplayer-server/create-build-with-custom-container?view=playfab-rest

{
  "ContainerRunCommand": "/data/Assets -startserver",
  "GameAssetReferences": [
    {
      "FileName": "gameserver.zip",
      "MountPath": "/data/Assets"
    }
  ],
  "ContainerImageReference": {
    "ImageName": "ContainerImageName",
    "Tag": "ContainerTag"
  },
  "LinuxInstrumentationConfiguration": {
    "IsEnabled": false
  },
  "ContainerFlavor": "CustomLinux",
  "BuildName": "GameBuildName",
  "Metadata": {
    "MetadataKey": "MetadataValue"
  },
  "VmSize": "Standard_D2_v2",
  "MultiplayerServerCountPerVm": 10,
  "Ports": [
    {
      "Name": "PortName",
      "Num": 1243,
      "Protocol": "TCP"
    }
  ],
  "RegionConfigurations": [
    {
      "Region": "EastUs",
      "MaxServers": 10,
      "StandbyServers": 5,
      "ScheduledStandbySettings": {
        "IsEnabled": true,
        "ScheduleList": [
          {
            "StartTime": "2020-08-21T17:00:00Z",
            "EndTime": "2020-08-24T09:00:00Z",
            "IsRecurringWeekly": true,
            "IsDisabled": false,
            "Description": "Weekend Schedule",
            "TargetStandby": 8
          },
          {
            "StartTime": "2020-08-24T09:00:00Z",
            "EndTime": "2020-08-28T17:00:00Z",
            "IsRecurringWeekly": true,
            "IsDisabled": false,
            "Description": "Weekday Schedule",
            "TargetStandby": 3
          }
        ]
      }
    },
    {
      "Region": "WestUs",
      "MaxServers": 50,
      "StandbyServers": 8,
      "ScheduledStandbySettings": {
        "IsEnabled": true,
        "ScheduleList": [
          {
            "StartTime": "2020-08-21T09:00:00Z",
            "EndTime": "2020-08-21T23:00:00Z",
            "IsRecurringWeekly": false,
            "IsDisabled": false,
            "Description": "Game Launch",
            "TargetStandby": 30
          }
        ]
      }
    },
    {
      "Region": "NorthEurope",
      "MaxServers": 7,
      "StandbyServers": 3
    }
  ],
  "GameCertificateReferences": [
    {
      "Name": "CertName",
      "GsdkAlias": "CertGsdkAlias"
    }
  ]
}

# Get entity Token

https://docs.microsoft.com/en-us/rest/api/playfab/authentication/authentication/get-entity-token?view=playfab-rest

X-SecretKey

# Delete Build

https://docs.microsoft.com/en-us/rest/api/playfab/multiplayer/multiplayer-server/delete-build?view=playfab-rest

{
    buildId
}

# Set Title Data
https://docs.microsoft.com/en-us/rest/api/playfab/admin/title-wide-data-management/set-title-data?view=playfab-rest
X-SecretKey

{
    Key,
    Value
}

# Get Title Data

https://docs.microsoft.com/en-us/rest/api/playfab/admin/title-wide-data-management/get-title-data?view=playfab-rest

x secret key

{
    Keys: [""]
}