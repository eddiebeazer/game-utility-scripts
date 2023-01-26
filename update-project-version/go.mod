module bitbucket.org/thedigitalsages/update-project-version

go 1.16

replace bitbucket.org/thedigitalsages/playfab-rest-api/playfabApi => ../playfab-rest-api

require (
	bitbucket.org/thedigitalsages/playfab-rest-api/playfabApi v0.0.0-00010101000000-000000000000
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.67.0
)
