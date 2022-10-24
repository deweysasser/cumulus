package cumulus

import "context"

// WARNING:  this file is generated. DO NOT EDIT.  Edit the source instead

// Instancer marks a type that can provide a channel of Instance
type Instancer interface {
	Instances(context.Context) chan Instance
}

// Snapshoter marks a type that can provide a channel of Snapshot
type Snapshoter interface {
	Snapshots(context.Context) chan Snapshot
}

// MachineImager marks a type that can provide a channel of MachineImage
type MachineImager interface {
	MachineImages(context.Context) chan MachineImage
}

// AccountInfoer marks a type that can provide a channel of AccountInfo
type AccountInfoer interface {
	AccountInfos(context.Context) chan AccountInfo
}

// Zoner marks a type that can provide a channel of Zone
type Zoner interface {
	Zones(context.Context) chan Zone
}

// NameRecorder marks a type that can provide a channel of NameRecord
type NameRecorder interface {
	NameRecords(context.Context) chan NameRecord
}
