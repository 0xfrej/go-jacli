// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-02-16 15:30:14.682291912 +0100 CET m=+0.004885438
package flag

import (
	"fmt"
	"strconv"
)
type Float32Flag struct {
	ValueFlag[float32]

	Names       []string
	Required    bool
	Description string
	value       float32
	wasSet      bool
}

func (f *Float32Flag) NameList() []string {
	return f.Names
}

func (f *Float32Flag) Value() float32 {
	return f.value
}

func (f *Float32Flag) IsSet() bool {
	return f.wasSet
}

func (f *Float32Flag) IsRequired() bool {
	return f.Required
}

func (f *Float32Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseFloat(a.String(), 32)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], a.String())
	}
	f.value = float32(v)
	f.wasSet = true
	return nil
}

func (f *Float32Flag) HelpDescription() string {
	return f.Description
}
type Float64Flag struct {
	ValueFlag[float64]

	Names       []string
	Required    bool
	Description string
	value       float64
	wasSet      bool
}

func (f *Float64Flag) NameList() []string {
	return f.Names
}

func (f *Float64Flag) Value() float64 {
	return f.value
}

func (f *Float64Flag) IsSet() bool {
	return f.wasSet
}

func (f *Float64Flag) IsRequired() bool {
	return f.Required
}

func (f *Float64Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseFloat(a.String(), 64)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], a.String())
	}
	f.value = float64(v)
	f.wasSet = true
	return nil
}

func (f *Float64Flag) HelpDescription() string {
	return f.Description
}
type IntFlag struct {
	ValueFlag[int]

	Names       []string
	Required    bool
	Description string
	value       int
	wasSet      bool
}

func (f *IntFlag) NameList() []string {
	return f.Names
}

func (f *IntFlag) Value() int {
	return f.value
}

func (f *IntFlag) IsSet() bool {
	return f.wasSet
}

func (f *IntFlag) IsRequired() bool {
	return f.Required
}

func (f *IntFlag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, 0)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = int(v)
	f.wasSet = true
	return nil
}

func (f *IntFlag) HelpDescription() string {
	return f.Description
}

type UintFlag struct {
	ValueFlag[uint]

	Names       []string
	Required    bool
	Description string
	value       uint
	wasSet      bool
}

func (f *UintFlag) NameList() []string {
	return f.Names
}

func (f *UintFlag) Value() uint {
	return f.value
}

func (f *UintFlag) IsSet() bool {
	return f.wasSet
}

func (f *UintFlag) IsRequired() bool {
	return f.Required
}

func (f *UintFlag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, 0)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = uint(v)
	f.wasSet = true
	return nil
}

func (f *UintFlag) HelpDescription() string {
	return f.Description
}
type Int8Flag struct {
	ValueFlag[int8]

	Names       []string
	Required    bool
	Description string
	value       int8
	wasSet      bool
}

func (f *Int8Flag) NameList() []string {
	return f.Names
}

func (f *Int8Flag) Value() int8 {
	return f.value
}

func (f *Int8Flag) IsSet() bool {
	return f.wasSet
}

func (f *Int8Flag) IsRequired() bool {
	return f.Required
}

func (f *Int8Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, 8)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = int8(v)
	f.wasSet = true
	return nil
}

func (f *Int8Flag) HelpDescription() string {
	return f.Description
}

type Uint8Flag struct {
	ValueFlag[uint8]

	Names       []string
	Required    bool
	Description string
	value       uint8
	wasSet      bool
}

func (f *Uint8Flag) NameList() []string {
	return f.Names
}

func (f *Uint8Flag) Value() uint8 {
	return f.value
}

func (f *Uint8Flag) IsSet() bool {
	return f.wasSet
}

func (f *Uint8Flag) IsRequired() bool {
	return f.Required
}

func (f *Uint8Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, 8)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = uint8(v)
	f.wasSet = true
	return nil
}

func (f *Uint8Flag) HelpDescription() string {
	return f.Description
}
type Int16Flag struct {
	ValueFlag[int16]

	Names       []string
	Required    bool
	Description string
	value       int16
	wasSet      bool
}

func (f *Int16Flag) NameList() []string {
	return f.Names
}

func (f *Int16Flag) Value() int16 {
	return f.value
}

func (f *Int16Flag) IsSet() bool {
	return f.wasSet
}

func (f *Int16Flag) IsRequired() bool {
	return f.Required
}

func (f *Int16Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, 16)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = int16(v)
	f.wasSet = true
	return nil
}

func (f *Int16Flag) HelpDescription() string {
	return f.Description
}

type Uint16Flag struct {
	ValueFlag[uint16]

	Names       []string
	Required    bool
	Description string
	value       uint16
	wasSet      bool
}

func (f *Uint16Flag) NameList() []string {
	return f.Names
}

func (f *Uint16Flag) Value() uint16 {
	return f.value
}

func (f *Uint16Flag) IsSet() bool {
	return f.wasSet
}

func (f *Uint16Flag) IsRequired() bool {
	return f.Required
}

func (f *Uint16Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, 16)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = uint16(v)
	f.wasSet = true
	return nil
}

func (f *Uint16Flag) HelpDescription() string {
	return f.Description
}
type Int32Flag struct {
	ValueFlag[int32]

	Names       []string
	Required    bool
	Description string
	value       int32
	wasSet      bool
}

func (f *Int32Flag) NameList() []string {
	return f.Names
}

func (f *Int32Flag) Value() int32 {
	return f.value
}

func (f *Int32Flag) IsSet() bool {
	return f.wasSet
}

func (f *Int32Flag) IsRequired() bool {
	return f.Required
}

func (f *Int32Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, 32)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = int32(v)
	f.wasSet = true
	return nil
}

func (f *Int32Flag) HelpDescription() string {
	return f.Description
}

type Uint32Flag struct {
	ValueFlag[uint32]

	Names       []string
	Required    bool
	Description string
	value       uint32
	wasSet      bool
}

func (f *Uint32Flag) NameList() []string {
	return f.Names
}

func (f *Uint32Flag) Value() uint32 {
	return f.value
}

func (f *Uint32Flag) IsSet() bool {
	return f.wasSet
}

func (f *Uint32Flag) IsRequired() bool {
	return f.Required
}

func (f *Uint32Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, 32)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = uint32(v)
	f.wasSet = true
	return nil
}

func (f *Uint32Flag) HelpDescription() string {
	return f.Description
}
type Int64Flag struct {
	ValueFlag[int64]

	Names       []string
	Required    bool
	Description string
	value       int64
	wasSet      bool
}

func (f *Int64Flag) NameList() []string {
	return f.Names
}

func (f *Int64Flag) Value() int64 {
	return f.value
}

func (f *Int64Flag) IsSet() bool {
	return f.wasSet
}

func (f *Int64Flag) IsRequired() bool {
	return f.Required
}

func (f *Int64Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, 64)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = int64(v)
	f.wasSet = true
	return nil
}

func (f *Int64Flag) HelpDescription() string {
	return f.Description
}

type Uint64Flag struct {
	ValueFlag[uint64]

	Names       []string
	Required    bool
	Description string
	value       uint64
	wasSet      bool
}

func (f *Uint64Flag) NameList() []string {
	return f.Names
}

func (f *Uint64Flag) Value() uint64 {
	return f.value
}

func (f *Uint64Flag) IsSet() bool {
	return f.wasSet
}

func (f *Uint64Flag) IsRequired() bool {
	return f.Required
}

func (f *Uint64Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, 64)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = uint64(v)
	f.wasSet = true
	return nil
}

func (f *Uint64Flag) HelpDescription() string {
	return f.Description
}