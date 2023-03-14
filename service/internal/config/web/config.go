package admin

func NewAppConfig() IAppConfig {
	return newAppConfig()
}

func NewOpsConfig() IOpsConfig {
	return newOpsConfig()
}
