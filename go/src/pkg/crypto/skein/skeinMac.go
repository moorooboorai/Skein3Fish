package skein

import (
    "os"
    )
type SkeinMac struct {
    skein *Skein
    stateSave []uint64
}

// Initializes a Skein MAC context.
// 
// Initializes the context with this data and saves the resulting Skein 
// state variables for further use.
//
// Applications call the normal Skein functions to update the MAC and
// get the final result.
//
// stateSize
//     Which Skein state size to use. Supported values 
//     are 256, 512, and 1024
// outputSize
//     Number of MAC hash bits to compute
// key
//     The key bytes
//
func NewMac(stateSize, outputSize int, key []byte) (s *SkeinMac, err os.Error) {
    s = new(SkeinMac)
    s.skein, err = NewExtended(stateSize, outputSize, 0, key)
    if err != nil {
        return nil, err
    }
    s.stateSave = make([]uint64, s.skein.cipherStateWords)
    copy(s.stateSave, s.skein.state)
    return s, nil
}

// Update Skein MAC with the next part of the message.
//
// input
//      Byte slice that contains data to hash.
//
func (s *SkeinMac) Update(input []byte) {
    s.skein.Update(input)
}

// Update the MAC with a message bit string.
//
// Skein can handle data not only as bytes but also as bit strings of
// arbitrary length (up to its maximum design size).
//
// input
//      Byte slice that contains data to hash. The length of the byte slice
//      must match the formula: (numBits + 7) / 8.
// numBits
//      Number of bits to hash.
//
func (s *SkeinMac) UpdateBits(input []byte, numBits int) os.Error {
    return s.skein.UpdateBits(input, numBits)
}

// Finalize Skein MAC and return the hash.
// 
// This method resets the Skein MAC after it computed the MAC. An
// application may reuse this Skein MAC context to compute another
// MAC with the same key and sizes.
//
func (s *SkeinMac) DoFinal() (hash []byte) {
    hash = s.skein.DoFinal()
    s.Reset()
    return
}

// Resets a Skein context for further use.
// 
// Restores the saved chaining variables to reset the Skein context. 
// Thus applications can reuse the same setup to  process several 
// messages. This saves a complete Skein initialization cycle.
// 
func (s *SkeinMac) Reset() {
    s.skein.initializeWithState(s.stateSave)
}
