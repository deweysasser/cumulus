package cumulus

import "context"

type Zone interface {
	Id() ID
	JSON() string
	Text() string
}

type ZoneVisitor func(ctx context.Context, cancel context.CancelFunc, zone Zone) error

type VisitZoneser interface {
	VisitZones(ctx context.Context, cancel context.CancelFunc, visitor ZoneVisitor) error
}
