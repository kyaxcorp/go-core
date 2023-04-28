package version

import "time"

var appVersion AppVersion // This is the defined version

var (
	ProjectName string // what's the project name
	Version     string // This is the version of the app:  1.0.0
	ServiceName string // this is the service name (it's optional)
	Stage       string // Dev, Staging, PreRelease, Prod etc...

	BuildAt        time.Time // WHen it has been built
	CompanyName    string
	CompanyWebsite string
	BuiltBy        string // By what user has been built

	// Git Details
	GitCommitId   string    // This is the latest commit id
	GitCommitDate time.Time // This is the latest commit date
	GitAuthorName string    // This is the author of the latest commit
)

type AppVersion struct {
	ProjectName string // what's the project name
	Version     string // This is the version of the app:  1.0.0
	ServiceName string // this is the service name (it's optional)
	Stage       string // Dev, Staging, PreRelease, Prod etc...

	BuildAt        time.Time // WHen it has been built
	CompanyName    string
	CompanyWebsite string
	BuiltBy        string // By what user has been built

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
