package caws

import (
	"context"
	"github.com/deweysasser/golang-program/cumulus/stats"
)

var CallTimer = stats.NewTimer(context.Background(), "AWS API calls")
