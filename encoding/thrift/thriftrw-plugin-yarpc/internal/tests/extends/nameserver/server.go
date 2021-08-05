// Code generated by thriftrw-plugin-yarpc
// @generated

package nameserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	extends "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Name service.
type Interface interface {
	Name(
		ctx context.Context,
	) (string, error)
}

// New prepares an implementation of the Name service for
// registration.
//
// 	handler := NameHandler{}
// 	dispatcher.Register(nameserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Name",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "name",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Name),

					NoWire: Name_NoWireHandler{impl},
				},
				Signature:    "Name() (string)",
				ThriftModule: extends.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) Name(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args extends.Name_Name_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Name' procedure 'Name': %w", err)
	}

	success, appErr := h.impl.Name(ctx)

	hadError := appErr != nil
	result, err := extends.Name_Name_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

type Name_NoWireHandler struct{ impl Interface }

func (h Name_NoWireHandler) Handle(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args extends.Name_Name_Args

		rw stream.ResponseWriter

		err error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)

	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'Name' procedure 'Name': %w", err)
	}

	success, appErr := h.impl.Name(ctx)

	hadError := appErr != nil
	result, err := extends.Name_Name_Helper.WrapResponse(success, appErr)
	var response thrift.NoWireResponse
	response.ResponseWriter = rw
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}
