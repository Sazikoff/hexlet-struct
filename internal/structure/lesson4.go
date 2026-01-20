package structure

type Rack struct {
	ID          string
	Temperature float64
	PowerOff    bool
}

// BEGIN (write your solution here)
func Adjust(r *Rack, delta float64) bool {
	if r==nil {
		return false
	}
	r.Temperature += delta
	return true
}

func EmergencyShutdown(r *Rack) {
	if r != nil {
		r.PowerOff = true
	}
}

func Snapshot(r *Rack) Rack {
	var Ra Rack
	if r == nil {
		return Ra
	}
	Ra = *r
	return Ra
}
// END