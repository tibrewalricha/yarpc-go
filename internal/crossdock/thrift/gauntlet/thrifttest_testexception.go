// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ThriftTest_TestException_Args represents the arguments for the ThriftTest.testException function.
//
// The arguments for testException are sent and received over the wire as this struct.
type ThriftTest_TestException_Args struct {
	Arg *string `json:"arg,omitempty"`
}

// ToWire translates a ThriftTest_TestException_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestException_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg != nil {
		w, err = wire.NewValueString(*(v.Arg)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestException_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestException_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestException_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestException_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestException_Args
// struct.
func (v *ThriftTest_TestException_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}

	return fmt.Sprintf("ThriftTest_TestException_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestException_Args match the
// provided ThriftTest_TestException_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestException_Args) Equals(rhs *ThriftTest_TestException_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Arg, rhs.Arg) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestException_Args.
func (v *ThriftTest_TestException_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Arg != nil {
		enc.AddString("arg", *v.Arg)
	}
	return err
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestException_Args) GetArg() (o string) {
	if v.Arg != nil {
		return *v.Arg
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testException" for this struct.
func (v *ThriftTest_TestException_Args) MethodName() string {
	return "testException"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestException_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestException_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testException
// function.
var ThriftTest_TestException_Helper = struct {
	// Args accepts the parameters of testException in-order and returns
	// the arguments struct for the function.
	Args func(
		arg *string,
	) *ThriftTest_TestException_Args

	// IsException returns true if the given error can be thrown
	// by testException.
	//
	// An error can be thrown by testException only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testException
	// given the error returned by it. The provided error may
	// be nil if testException did not fail.
	//
	// This allows mapping errors returned by testException into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// testException
	//
	//   err := testException(args)
	//   result, err := ThriftTest_TestException_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testException: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*ThriftTest_TestException_Result, error)

	// UnwrapResponse takes the result struct for testException
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if testException threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := ThriftTest_TestException_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestException_Result) error
}{}

func init() {
	ThriftTest_TestException_Helper.Args = func(
		arg *string,
	) *ThriftTest_TestException_Args {
		return &ThriftTest_TestException_Args{
			Arg: arg,
		}
	}

	ThriftTest_TestException_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *Xception:
			return true
		default:
			return false
		}
	}

	ThriftTest_TestException_Helper.WrapResponse = func(err error) (*ThriftTest_TestException_Result, error) {
		if err == nil {
			return &ThriftTest_TestException_Result{}, nil
		}

		switch e := err.(type) {
		case *Xception:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestException_Result.Err1")
			}
			return &ThriftTest_TestException_Result{Err1: e}, nil
		}

		return nil, err
	}
	ThriftTest_TestException_Helper.UnwrapResponse = func(result *ThriftTest_TestException_Result) (err error) {
		if result.Err1 != nil {
			err = result.Err1
			return
		}
		return
	}

}

// ThriftTest_TestException_Result represents the result of a ThriftTest.testException function call.
//
// The result of a testException execution is sent and received over the wire as this struct.
type ThriftTest_TestException_Result struct {
	Err1 *Xception `json:"err1,omitempty"`
}

// ToWire translates a ThriftTest_TestException_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestException_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Err1 != nil {
		w, err = v.Err1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestException_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Xception_Read(w wire.Value) (*Xception, error) {
	var v Xception
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a ThriftTest_TestException_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestException_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestException_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestException_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err1, err = _Xception_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Err1 != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("ThriftTest_TestException_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestException_Result
// struct.
func (v *ThriftTest_TestException_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Err1 != nil {
		fields[i] = fmt.Sprintf("Err1: %v", v.Err1)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestException_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestException_Result match the
// provided ThriftTest_TestException_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestException_Result) Equals(rhs *ThriftTest_TestException_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Err1 == nil && rhs.Err1 == nil) || (v.Err1 != nil && rhs.Err1 != nil && v.Err1.Equals(rhs.Err1))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestException_Result.
func (v *ThriftTest_TestException_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Err1 != nil {
		err = multierr.Append(err, enc.AddObject("err1", v.Err1))
	}
	return err
}

// GetErr1 returns the value of Err1 if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestException_Result) GetErr1() (o *Xception) {
	if v.Err1 != nil {
		return v.Err1
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testException" for this struct.
func (v *ThriftTest_TestException_Result) MethodName() string {
	return "testException"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestException_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
