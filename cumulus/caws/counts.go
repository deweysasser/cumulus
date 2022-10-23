package caws

import (
	"context"
	"github.com/deweysasser/cumulus/cumulus/stats"
)

var CallTimer = stats.NewTimer(context.Background(), "AWS API calls")
