package proxy

// Bool interface create a proxy for bool type (used for mocking)
type Bool interface {
	Get() bool
	Set(bool)
}

// BoolProxy wraps bool type
type BoolProxy struct {
	val bool
}

// Get returns the underling bool type
func (o *BoolProxy) Get() bool {
	return o.val
}

// Set sets the underling bool type
func (o *BoolProxy) Set(val bool) {
	o.val = val
}
