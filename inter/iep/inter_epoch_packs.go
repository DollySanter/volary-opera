package iep

import (
	"github.com/DollySanter/volary-opera/inter"
	"github.com/DollySanter/volary-opera/inter/ier"
)

type LlrEpochPack struct {
	Votes  []inter.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
