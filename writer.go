package simplelog

// devNull...
type devNull int

// Write ...
func (d devNull) Write(b []byte) (int, error) {
	return len(b), nil
}

// Close ...
func (d devNull) Close() error {
	return nil
}

// Discard ...
var Discard = devNull(0)
