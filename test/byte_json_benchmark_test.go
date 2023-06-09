package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"graph_robot/neure"
	"log"
	"testing"
)

var Neure1 = neure.Neure{
	Synapses:               make(map[string]*neure.Synapse),
	NowLinkedDendritesIds:  make(map[string]struct{}),
	ElectricalConductivity: 4423423,
}

func BenchmarkJson(*testing.B) {
	for i := 0; i < 10000; i++ {
		// use json is more faster
		nb, _ := json.Marshal(Neure1)
		_ = string(nb)
		var neu neure.Neure
		_ = json.Unmarshal(nb, &neu)
	}
	// result BenchmarkJson
	// BenchmarkJson-8   	1000000000	         0.01597 ns/op	       0 B/op	       0 allocs/op
	// for now, this is the most efficient way.
}

func BenchmarkGob(*testing.B) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	for i := 0; i < 10000; i++ {

		// Encode (send) the value.
		err := enc.Encode(Neure1)
		if err != nil {
			log.Fatal("encode error:", err)
		}

		_ = network.Bytes()

		// Decode (receive) the value.
		var q neure.Neure
		err = dec.Decode(&q)
		if err != nil {
			log.Fatal("decode error:", err)
		}
	}
	// result BenchmarkGob
	// BenchmarkGob-8   	1000000000	         0.2523 ns/op	       0 B/op	       0 allocs/op
	// although this way is faster than json, but it's not thread safe.
}
