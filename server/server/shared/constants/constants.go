package constants

// **** Vals **** //

const (
	sessionDataKey = "session-data"
)


/**** Functions ****/

// The key used to identify session data on the context
// as it's passed down through the middleware.
func SessionDataKey() string {
	return sessionDataKey
}
