package tracker

import "context"

// tracker is the primary interface for the various tracker options
type Tracker interface {
	//Exec will be run as go function
	//the client to use for tracking ops.
	Exec(ctx context.Context) error
	String() string
}
