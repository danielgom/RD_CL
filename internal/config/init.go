package config

var environment = struct {
	isTest             bool
	dBConnectionString string
}{}

func InitialiseTest(pgTestConnString, dbName string) {
	initialiseTestEnvironment(pgTestConnString)
	makeTestDB(dbName)
}

func initialiseTestEnvironment(pgTestConnString string) {
	environment.isTest = true
	environment.dBConnectionString = pgTestConnString
}
