package serde

import (
	"errors"
)

type Int8Visitor struct {
	v *int8
}

func NewInt8Visitor(v *int8) Int8Visitor {
	return Int8Visitor{v: v}
}

func (vi Int8Visitor) String() string {
	return "Int8"
}

func (vi Int8Visitor) VisitInt8(v int8) (err error) {
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitUint8(v uint8) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitInt16(v int16) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitInt32(v int32) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitInt(v int) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

func (vi Int8Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*vi.v = int8(v)
	return nil
}

type Uint8Visitor struct {
	v *uint8
}

func NewUint8Visitor(v *uint8) Uint8Visitor {
	return Uint8Visitor{v: v}
}

func (vi Uint8Visitor) String() string {
	return "Uint8"
}

func (vi Uint8Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitInt16(v int16) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitInt32(v int32) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitInt(v int) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

func (vi Uint8Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*vi.v = uint8(v)
	return nil
}

type Int16Visitor struct {
	v *int16
}

func NewInt16Visitor(v *int16) Int16Visitor {
	return Int16Visitor{v: v}
}

func (vi Int16Visitor) String() string {
	return "Int16"
}

func (vi Int16Visitor) VisitInt8(v int8) (err error) {
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitInt16(v int16) (err error) {
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitInt32(v int32) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitInt(v int) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

func (vi Int16Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*vi.v = int16(v)
	return nil
}

type Uint16Visitor struct {
	v *uint16
}

func NewUint16Visitor(v *uint16) Uint16Visitor {
	return Uint16Visitor{v: v}
}

func (vi Uint16Visitor) String() string {
	return "Uint16"
}

func (vi Uint16Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitUint16(v uint16) (err error) {
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitInt32(v int32) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitInt(v int) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

func (vi Uint16Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*vi.v = uint16(v)
	return nil
}

type Int32Visitor struct {
	v *int32
}

func NewInt32Visitor(v *int32) Int32Visitor {
	return Int32Visitor{v: v}
}

func (vi Int32Visitor) String() string {
	return "Int32"
}

func (vi Int32Visitor) VisitInt8(v int8) (err error) {
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitInt16(v int16) (err error) {
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitUint16(v uint16) (err error) {
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitInt32(v int32) (err error) {
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitInt(v int) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	if v < MinInt32 {
		return errors.New("overflow")
	}
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	if v < MinInt32 {
		return errors.New("overflow")
	}
	*vi.v = int32(v)
	return nil
}

func (vi Int32Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*vi.v = int32(v)
	return nil
}

type IntVisitor struct {
	v *int
}

func NewIntVisitor(v *int) IntVisitor {
	return IntVisitor{v: v}
}

func (vi IntVisitor) String() string {
	return "Int"
}

func (vi IntVisitor) VisitInt8(v int8) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitUint8(v uint8) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitInt16(v int16) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitUint16(v uint16) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitInt32(v int32) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitInt(v int) (err error) {
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitUint32(v uint32) (err error) {
	if UintSize == 32 && v > MaxInt32 {
		return errors.New("overflow")
	}
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitUint(v uint) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitInt64(v int64) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	if v < MinInt {
		return errors.New("overflow")
	}
	*vi.v = int(v)
	return nil
}

func (vi IntVisitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	*vi.v = int(v)
	return nil
}

type Uint32Visitor struct {
	v *uint32
}

func NewUint32Visitor(v *uint32) Uint32Visitor {
	return Uint32Visitor{v: v}
}

func (vi Uint32Visitor) String() string {
	return "Uint32"
}

func (vi Uint32Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitUint16(v uint16) (err error) {
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitInt32(v int32) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitInt(v int) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitUint32(v uint32) (err error) {
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

func (vi Uint32Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	*vi.v = uint32(v)
	return nil
}

type UintVisitor struct {
	v *uint
}

func NewUintVisitor(v *uint) UintVisitor {
	return UintVisitor{v: v}
}

func (vi UintVisitor) String() string {
	return "Uint"
}

func (vi UintVisitor) VisitInt8(v int8) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitUint8(v uint8) (err error) {
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitInt16(v int16) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitUint16(v uint16) (err error) {
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitInt32(v int32) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitInt(v int) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitUint32(v uint32) (err error) {
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitUint(v uint) (err error) {
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitInt64(v int64) (err error) {
	if UintSize == 32 && v > MaxUint32 {
		return errors.New("overflow")
	}
	if v < MinUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

func (vi UintVisitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint {
		return errors.New("overflow")
	}
	*vi.v = uint(v)
	return nil
}

type Int64Visitor struct {
	v *int64
}

func NewInt64Visitor(v *int64) Int64Visitor {
	return Int64Visitor{v: v}
}

func (vi Int64Visitor) String() string {
	return "Int64"
}

func (vi Int64Visitor) VisitInt8(v int8) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitInt16(v int16) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitUint16(v uint16) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitInt32(v int32) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitInt(v int) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitUint32(v uint32) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitUint(v uint) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitInt64(v int64) (err error) {
	*vi.v = int64(v)
	return nil
}

func (vi Int64Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt64 {
		return errors.New("overflow")
	}
	*vi.v = int64(v)
	return nil
}

type Uint64Visitor struct {
	v *uint64
}

func NewUint64Visitor(v *uint64) Uint64Visitor {
	return Uint64Visitor{v: v}
}

func (vi Uint64Visitor) String() string {
	return "Uint64"
}

func (vi Uint64Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitUint8(v uint8) (err error) {
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitUint16(v uint16) (err error) {
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitInt32(v int32) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitInt(v int) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitUint32(v uint32) (err error) {
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitUint(v uint) (err error) {
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitInt64(v int64) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*vi.v = uint64(v)
	return nil
}

func (vi Uint64Visitor) VisitUint64(v uint64) (err error) {
	*vi.v = uint64(v)
	return nil
}
