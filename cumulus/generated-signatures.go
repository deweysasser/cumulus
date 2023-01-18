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

// Volumer marks a type that can provide a channel of Volume
type Volumer interface {
	Volumes(context.Context) chan Volume
}

// Subscriptioner marks a type that can provide a channel of Subscription
type Subscriptioner interface {
	Subscriptions(context.Context) chan Subscription
}

// Topicer marks a type that can provide a channel of Topic
type Topicer interface {
	Topics(context.Context) chan Topic
}

// DBClusterer marks a type that can provide a channel of DBCluster
type DBClusterer interface {
	DBClusters(context.Context) chan DBCluster
}

// DBInstancer marks a type that can provide a channel of DBInstance
type DBInstancer interface {
	DBInstances(context.Context) chan DBInstance
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
