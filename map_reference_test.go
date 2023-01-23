package sync_test

// mapInterface is the interface Map implements.
type mapInterface interface {
	Load(any) (any, bool)
	Store(key, value any)
	LoadOrStore(key, value any) (actual any, loaded bool)
	LoadAndDelete(key any) (value any, loaded bool)
	Delete(any)
	Range(func(key, value any) (shouldContinue bool))
}
