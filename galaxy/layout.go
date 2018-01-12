package galaxy

type Layout interface {
	saveToDataStore()
	getFromDataStore()
	cacheDirectoryName()
}
