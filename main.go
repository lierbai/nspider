package main

import "fmt"

type A interface {
}

// B 1
type B interface {
}

// C 1
type C interface {
	GetB(a A) B
}

// RA R
type RA struct {
	data map[string]string
}

// RB R
type RB struct {
	rdata map[string]string
}

// RC R
type RC struct {
}

// GetB r
func (rc *RC) GetB(data map[string]string, ra *RA) *RB {
	return &RB{ra.data}
}

func main() {
	data := make(map[string]string)
	data["num"] = "ten"
	rc := &RC{}
	rb := rc.GetB(data, &RA{data})
	fmt.Println(rb.rdata)
	// &RA{data}
}
