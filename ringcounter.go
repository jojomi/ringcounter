package ringcounter

// Ringcounter is a counter that is wrapped at a given size (operations mod n).
type Ringcounter struct {
	Value int
	Size  int
}

// New returns a new Ringcounter with given initial state.
func New(size, initValue int) *Ringcounter {
	r := &Ringcounter{
		Value: initValue,
		Size:  size,
	}
	return r
}

// Add increments the ringcounter by the given amount. Can be negative.
func (r *Ringcounter) Add(value int) {
	v := r.Value
	v += value
	for v < 0 {
		v += r.Size
	}
	r.Value = v - (v/r.Size)*r.Size
}

// Advance is a convenience method that adds one to the given ringcounter.
func (r *Ringcounter) Advance() {
	r.Add(1)
}
