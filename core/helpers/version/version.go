package version

import "time"

var appVersion AppVersion // This is the defined version

type AppVersion struct {
	ProjectName string // what's the project name -> for example: connext-express
	Version     string // This is the version of the app:  1.0.0
	ServiceName string // this is the service name (it's optional)

	BuildTime time.Time // WHen it has being built
	BuildBy   string    // By what user has being built

	// Git Details
	GitCommitId   string    // This is the latest commit id
	GitCommitDate time.Time // This is the latest commit date
	GitAuthorName string    // This is the author of the latest commit
}

func Init(version AppVersion) {
	// Add here GIT Commit ID
	// Commit Date
	// Add here Go Build Version
	// Add here Build Time
	// Add here Latest GIT Author
	appVersion = version
}

func GetAppVersion() AppVersion {
	return appVersion
}
