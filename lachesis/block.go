package lachesis

import (
	"github.com/panoptisDev/lachesis-base/hash"
)

// Block is a part of an ordered chain of batches of events.
type Block struct {
	Electing hash.Event
	Atropos  hash.Event
	Cheaters Cheaters
}
