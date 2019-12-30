package domain

import (
	"crypto/md5"
	"encoding/binary"
)

// Creates a seed for the random generator when
// assigning variants of a given experiment to a user.
//
// This guarantees that the seed is always the same for
// each user and experiment pair.
func SeedFor(experiment Experiment, user User) int64 {
	// create an md5 from the experiment name and user id
	h := md5.New()
	h.Write([]byte(user.Id + experiment.Name))
	hash := h.Sum(nil)
	
	// Convert the hash to an int64
	// md5 has 128 bits so we'll loose half of them
	return int64(binary.BigEndian.Uint64(hash))
}
