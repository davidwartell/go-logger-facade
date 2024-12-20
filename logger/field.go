package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Field zapcore.Field

// Skip constructs a no-op field, which is often useful when handling invalid
// inputs in other Field constructors.
//
//goland:noinspection GoUnusedExportedFunction
func Skip() Field { return Field(zap.Skip()) }

// Binary constructs a field that carries an opaque binary blob.
//
// Binary data is serialized in an encoding-appropriate format. For example,
// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
// use ByteString.
//
//goland:noinspection GoUnusedExportedFunction
func Binary(key string, val []byte) Field { return Field(zap.Binary(key, val)) }

// Bool constructs a field that carries a bool.
//
//goland:noinspection GoUnusedExportedFunction
func Bool(key string, val bool) Field { return Field(zap.Bool(key, val)) }

// Boolp constructs a field that carries a *bool. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Boolp(key string, val *bool) Field { return Field(zap.Boolp(key, val)) }

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
// Binary.
//
//goland:noinspection GoUnusedExportedFunction
func ByteString(key string, val []byte) Field { return Field(zap.ByteString(key, val)) }

// Complex128 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex128 to
// interface{}).
//
//goland:noinspection GoUnusedExportedFunction
func Complex128(key string, val complex128) Field { return Field(zap.Complex128(key, val)) }

// Complex128p constructs a field that carries a *complex128. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Complex128p(key string, val *complex128) Field { return Field(zap.Complex128p(key, val)) }

// Complex64 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex64 to
// interface{}).
//
//goland:noinspection GoUnusedExportedFunction
func Complex64(key string, val complex64) Field { return Field(zap.Complex64(key, val)) }

// Complex64p constructs a field that carries a *complex64. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Complex64p(key string, val *complex64) Field { return Field(zap.Complex64p(key, val)) }

// Float64 constructs a field that carries a float64. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
//
//goland:noinspection GoUnusedExportedFunction
func Float64(key string, val float64) Field { return Field(zap.Float64(key, val)) }

// Float64p constructs a field that carries a *float64. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Float64p(key string, val *float64) Field { return Field(zap.Float64p(key, val)) }

// Float32 constructs a field that carries a float32. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
//
//goland:noinspection GoUnusedExportedFunction
func Float32(key string, val float32) Field { return Field(zap.Float32(key, val)) }

// Float32p constructs a field that carries a *float32. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Float32p(key string, val *float32) Field { return Field(zap.Float32p(key, val)) }

// Int constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Int(key string, val int) Field { return Field(zap.Int(key, val)) }

// Intp constructs a field that carries a *int. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Intp(key string, val *int) Field { return Field(zap.Intp(key, val)) }

// Int64 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Int64(key string, val int64) Field { return Field(zap.Int64(key, val)) }

// Int64p constructs a field that carries a *int64. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Int64p(key string, val *int64) Field { return Field(zap.Int64p(key, val)) }

// Int32 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Int32(key string, val int32) Field { return Field(zap.Int32(key, val)) }

// Int32p constructs a field that carries a *int32. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Int32p(key string, val *int32) Field { return Field(zap.Int32p(key, val)) }

// Int16 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Int16(key string, val int16) Field { return Field(zap.Int16(key, val)) }

// Int16p constructs a field that carries a *int16. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Int16p(key string, val *int16) Field { return Field(zap.Int16p(key, val)) }

// Int8 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Int8(key string, val int8) Field { return Field(zap.Int8(key, val)) }

// Int8p constructs a field that carries a *int8. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Int8p(key string, val *int8) Field { return Field(zap.Int8p(key, val)) }

// String constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func String(key string, val string) Field { return Field(zap.String(key, val)) }

// Stringp constructs a field that carries a *string. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Stringp(key string, val *string) Field { return Field(zap.Stringp(key, val)) }

// Uint constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uint(key string, val uint) Field { return Field(zap.Uint(key, val)) }

// Uintp constructs a field that carries a *uint. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uintp(key string, val *uint) Field { return Field(zap.Uintp(key, val)) }

// Uint64 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uint64(key string, val uint64) Field { return Field(zap.Uint64(key, val)) }

// Uint64p constructs a field that carries a *uint64. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uint64p(key string, val *uint64) Field { return Field(zap.Uint64p(key, val)) }

// Uint32 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uint32(key string, val uint32) Field { return Field(zap.Uint32(key, val)) }

// Uint32p constructs a field that carries a *uint32. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uint32p(key string, val *uint32) Field { return Field(zap.Uint32p(key, val)) }

// Uint16 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uint16(key string, val uint16) Field { return Field(zap.Uint16(key, val)) }

// Uint16p constructs a field that carries a *uint16. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uint16p(key string, val *uint16) Field { return Field(zap.Uint16p(key, val)) }

// Uint8 constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uint8(key string, val uint8) Field { return Field(zap.Uint8(key, val)) }

// Uint8p constructs a field that carries a *uint8. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uint8p(key string, val *uint8) Field { return Field(zap.Uint8p(key, val)) }

// Uintptr constructs a field with the given key and value.
//
//goland:noinspection GoUnusedExportedFunction
func Uintptr(key string, val uintptr) Field { return Field(zap.Uintptr(key, val)) }

// Uintptrp constructs a field that carries a *uintptr. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Uintptrp(key string, val *uintptr) Field { return Field(zap.Uintptrp(key, val)) }

// Reflect constructs a field with the given key and an arbitrary object. It uses
// an encoding-appropriate, reflection-based function to lazily serialize nearly
// any object into the logging context, but it's relatively slow and
// allocation-heavy. Outside tests, Any is always a better choice.
//
// If encoding fails (e.g., trying to serialize a map[int]string to JSON), Reflect
// includes the error message in the final log output.
//
//goland:noinspection GoUnusedExportedFunction
func Reflect(key string, val interface{}) Field { return Field(zap.Reflect(key, val)) }

// Namespace creates a named, isolated scope within the logger's context. All
// subsequent fields will be added to the new namespace.
//
// This helps prevent key collisions when injecting loggers into sub-components
// or third-party libraries.
//
//goland:noinspection GoUnusedExportedFunction
func Namespace(key string) Field { return Field(zap.Namespace(key)) }

// Stringer constructs a field with the given key and the output of the value's
// String method. The Stringer's String method is called lazily.
//
//goland:noinspection GoUnusedExportedFunction
func Stringer(key string, val fmt.Stringer) Field { return Field(zap.Stringer(key, val)) }

// Time constructs a Field with the given key and value. The encoder
// controls how the time is serialized.
//
//goland:noinspection GoUnusedExportedFunction
func Time(key string, val time.Time) Field { return Field(zap.Time(key, val)) }

// Timep constructs a field that carries a *time.Time. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Timep(key string, val *time.Time) Field { return Field(zap.Timep(key, val)) }

// Stack constructs a field that stores a stacktrace of the current goroutine
// under provided key. Keep in mind that taking a stacktrace is eager and
// expensive (relatively speaking); this function both makes an allocation and
// takes about two microseconds.
//
//goland:noinspection GoUnusedExportedFunction
func Stack(key string) Field { return Field(zap.Stack(key)) }

// StackSkip constructs a field similarly to Stack, but also skips the given
// number of frames from the top of the stacktrace.
//
//goland:noinspection GoUnusedExportedFunction
func StackSkip(key string, skip int) Field { return Field(zap.StackSkip(key, skip)) }

// Duration constructs a field with the given key and value. The encoder
// controls how the duration is serialized.
func Duration(key string, val time.Duration) Field { return Field(zap.Duration(key, val)) }

// Durationp constructs a field that carries a *time.Duration. The returned Field will safely
// and explicitly represent `nil` when appropriate.
//
//goland:noinspection GoUnusedExportedFunction
func Durationp(key string, val *time.Duration) Field { return Field(zap.Durationp(key, val)) }

// Object constructs a field with the given key and ObjectMarshaler. It
// provides a flexible, but still type-safe and efficient, way to add map- or
// struct-like user-defined types to the logging context. The struct's
// MarshalLogObject method is called lazily.
//
//goland:noinspection GoUnusedExportedFunction
func Object(key string, val zapcore.ObjectMarshaler) Field { return Field(zap.Object(key, val)) }

// Inline constructs a Field that is similar to Object, but it
// will add the elements of the provided ObjectMarshaler to the
// current namespace.
//
//goland:noinspection GoUnusedExportedFunction
func Inline(val zapcore.ObjectMarshaler) Field { return Field(zap.Inline(val)) }

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
//
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
// them. To minimize surprises, []byte values are treated as binary blobs, byte
// values are treated as uint8, and runes are always treated as integers
//
//goland:noinspection GoUnusedExportedFunction
func Any(key string, value interface{}) Field { return Field(zap.Any(key, value)) }

// Error is shorthand for the common idiom NamedError("error", err).
func Error(err error) Field {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	return Field(zap.String("error", errStr))
}

// NamedError constructs a field that lazily stores err.Error() under the
// provided key. Errors which also implement fmt.Formatter (like those produced
// by github.com/pkg/errors) will also have their verbose representation stored
// under key+"Verbose". If passed a nil error, the field is a no-op.
//
// For the common case in which the key is simply "error", the Error function
// is shorter and less repetitive.
//
//goland:noinspection GoUnusedExportedFunction
func NamedError(key string, err error) Field { return Field(zap.NamedError(key, err)) }

func fieldsContainContextCancelled(fields ...Field) bool {
	//for i := range fields {
	//	if fields[i].Type == zapcore.ErrorType {
	//		fieldErr, fieldIsError := fields[i].Interface.(error)
	//		if fieldIsError && fieldErr != nil && strings.Contains(fieldErr.Error(), "context canceled") {
	//			return true
	//		}
	//	}
	//	if fields[i].Type == zapcore.StringType && strings.Contains(fields[i].String, "context canceled") {
	//		return true
	//	}
	//}
	return false
}
