package instructions_service

type PredefinedConditions struct {
	AppExists            AppExistsCondition            `json:"app_exists"`
	FileExists           FileExistenceCondition        `json:"file_exists"`
	FileNotExists        FileExistenceCondition        `json:"file_not_exists"`
	OsVersion            OsVersionCondition            `json:"os_version"`
	ProcessNotRunning    ProcessNotRunningCondition    `json:"process_not_running"`
	ProfileExists        ProfileExistsCondition        `json:"profile_exists"`
	RequiredArchitecture RequiredArchitectureCondition `json:"required_architecture"`
}
