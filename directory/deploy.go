package directory

import (
	"github.com/hashicorp/otto/helper/uuid"
)

// DeployState is used to track the state of a deploy.
//
// This is required because a deploy is entered in the directory
// prior to the deploy actually happening so that we can always look
// up any binary blobs stored with a deploy even if it fails.
type DeployState byte

const (
	DeployStateInvalid DeployState = 0
	DeployStateNew     DeployState = iota
	DeployStateFail
	DeployStateSuccess
)

// Deploy represents a deploy of an App.
type Deploy struct {
	// Lookup information for the Deploy. AppID, Infra, and InfraFlavor
	// are required.
	Lookup

	// These fields should be set for Put and will be populated on Get
	State  DeployState       // State of the deploy
	Deploy map[string]string // Deploy information

	// Private fields. These are usually set on Get or Put.
	//
	// DO NOT MODIFY THESE.
	ID string
}

// IsNew reports if this deploy is freshly created and not yet run
func (d *Deploy) IsNew() bool {
	return d != nil && d.State == DeployStateNew
}

// MarkFailed sets a deploy's state to failed
func (d *Deploy) MarkFailed() {
	d.State = DeployStateFail
}

// MarkSuccess sets a deploy's state to success
func (d *Deploy) MarkSuccessful() {
	d.State = DeployStateSuccess
}

// MarkGone resets a deploy's state to the "new" state
func (d *Deploy) MarkGone() {
	d.State = DeployStateNew
}

func (d *Deploy) setId() {
	d.ID = uuid.GenerateUUID()
}
