package encode

// Meta ...
type Meta interface {
	Key() []byte
	Value() []byte
	Wrap() bool
	IsNil() bool
}
