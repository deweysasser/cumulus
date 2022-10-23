package cumulus
import "context"

// WARNING:  this file is generated. DO NOT EDIT.  Edit the source instead


// AccountInfoer marks a type that can provide a channel of AccountInfo
type AccountInfoer interface {
  AccountInfos(context.Context) chan AccountInfo
}

// Instancer marks a type that can provide a channel of Instance
type Instancer interface {
  Instances(context.Context) chan Instance
}

// Zoner marks a type that can provide a channel of Zone
type Zoner interface {
  Zones(context.Context) chan Zone
}

