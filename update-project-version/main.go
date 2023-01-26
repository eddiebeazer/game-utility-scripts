package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

const DevVersionFile = "E:\\Perforce\\elg_dev_version.txt"
const RelVersionFile = "E:\\Perforce\\elg_rel_version.txt"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCurrentVersion(versionFile string) string {
	data, err := ioutil.ReadFile(versionFile)
	CheckError(err)

	return string(data)
}

func UpdateVersionFile(newVersion string, versionFile string) {
	err := ioutil.WriteFile(versionFile, []byte(newVersion), 0777)
	CheckError(err)
}

func main() {
	// parsing command flags
	defaultGameIniPath := flag.String("iniPath", "", "file path to the DefaultGame ini file we want to edit")
	branch := flag.String("branch", "", "Branch of the build")
	flag.Parse()

	versionFile := DevVersionFile
	if strings.Contains(*branch, "patch") {
		versionFile = RelVersionFile
	}

	// if the new rel version has a different major/minor version than the current rel version, copy number from dev
	if strings.Contains(*branch, "rel") {
		relVerSplit := strings.Split(strings.TrimPrefix(*branch, "rel"), ".")
		curDevVerSplit := strings.Split(GetCurrentVersion(DevVersionFile), ".")
		majorDev, err := strconv.Atoi(curDevVerSplit[0])
		CheckError(err)
		minorDev, err := strconv.Atoi(curDevVerSplit[1])
		CheckError(err)
		majorRel, err := strconv.Atoi(relVerSplit[0])
		CheckError(err)
		minorRel, err := strconv.Atoi(relVerSplit[1])
		CheckError(err)

		if majorDev != majorRel || minorRel != minorDev {
			versionFile = DevVersionFile
		} else {
			versionFile = RelVersionFile
		}
	}

	currentVersion := GetCurrentVersion(versionFile)

	fmt.Printf("Version pulled from file: %s\n", currentVersion)

	semVerSplit := strings.Split(currentVersion, ".")
	major, err := strconv.Atoi(semVerSplit[0])
	CheckError(err)
	minor, err := strconv.Atoi(semVerSplit[1])
	CheckError(err)
	patch, err := strconv.Atoi(semVerSplit[2])
	CheckError(err)
	build, err := strconv.Atoi(semVerSplit[3])
	CheckError(err)

	// hard setting the version tag if the branch is dev
	if strings.Contains(*branch, "dev") {
		devVerSplit := strings.Split(strings.TrimPrefix(*branch, "dev"), ".")
		newMajor, err := strconv.Atoi(devVerSplit[0])
		CheckError(err)
		newMinor, err := strconv.Atoi(devVerSplit[1])
		CheckError(err)
		// Overriding the major/minor version if it's not currently set
		if newMajor != major || minor != newMinor {
			major = newMajor
			minor = newMinor
			patch = 0
			build = 1
		} else {
			build += 1
		}
	}

	// we should never be doing any dev on main branch.  So we just return the current version at the end

	// patch branches are for hotfixes applied directly to rel branch
	if strings.Contains(*branch, "patch") {
		patchVerSplit := strings.Split(strings.TrimPrefix(*branch, "patch"), ".")
		newMajor, err := strconv.Atoi(patchVerSplit[0])
		CheckError(err)
		newMinor, err := strconv.Atoi(patchVerSplit[1])
		CheckError(err)
		newPatch, err := strconv.Atoi(patchVerSplit[2])
		CheckError(err)

		// Overriding the major/minor/patch version if it's not currently set to this branch
		if newMajor != major || minor != newMinor || newPatch != patch {
			major = newMajor
			minor = newMinor
			patch = newPatch
			build = 1
		} else {
			build += 1
		}
	}

	// for task branches we want to increment the build number
	if strings.Contains(*branch, "task") {
		build += 1
	}

	newVersionNumber := fmt.Sprintf("%d.%d.%d.%d", major, minor, patch, build)

	// loading ini file
	ini.PrettyFormat = false
	cfg, err := ini.ShadowLoad(*defaultGameIniPath)
	CheckError(err)

	// updating game version in ini
	cfg.Section("/Script/EngineSettings.GeneralProjectSettings").Key("ProjectVersion").SetValue(newVersionNumber)
	err = cfg.SaveTo(*defaultGameIniPath)
	CheckError(err)

	// updating the global version file
	UpdateVersionFile(newVersionNumber, versionFile)

	// copying dev version to rel version if this is a new major/minor rel release
	if strings.Contains(*branch, "rel") && versionFile == DevVersionFile {
		UpdateVersionFile(newVersionNumber, RelVersionFile)
	}

	// printing the string to standard output
	fmt.Printf("New version written to file: %s\n", newVersionNumber)
}
