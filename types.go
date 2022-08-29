package ramda

// Numeric Define Numeric types for Generics
type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Ordered Define Ordered types for Generics
type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uintptr | string | float32 | float64
}
