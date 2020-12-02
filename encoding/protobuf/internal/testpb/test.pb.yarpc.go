// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: test.proto

package testpb

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// TestYARPCClient is the YARPC client-side interface for the Test service.
type TestYARPCClient interface {
	Unary(context.Context, *TestMessage, ...yarpc.CallOption) (*TestMessage, error)
	Duplex(context.Context, ...yarpc.CallOption) (TestServiceDuplexYARPCClient, error)
}

// TestServiceDuplexYARPCClient sends TestMessages and receives TestMessages, returning io.EOF when the stream is complete.
type TestServiceDuplexYARPCClient interface {
	Context() context.Context
	Send(*TestMessage, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*TestMessage, error)
	CloseSend(...yarpc.StreamOption) error
}

func newTestYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) TestYARPCClient {
	return &_TestYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.encoding.protobuf.Test",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewTestYARPCClient builds a new YARPC client for the Test service.
func NewTestYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) TestYARPCClient {
	return newTestYARPCClient(clientConfig, nil, options...)
}

// TestYARPCServer is the YARPC server-side interface for the Test service.
type TestYARPCServer interface {
	Unary(context.Context, *TestMessage) (*TestMessage, error)
	Duplex(TestServiceDuplexYARPCServer) error
}

// TestServiceDuplexYARPCServer receives TestMessages and sends TestMessage.
type TestServiceDuplexYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*TestMessage, error)
	Send(*TestMessage, ...yarpc.StreamOption) error
}

type buildTestYARPCProceduresParams struct {
	Server      TestYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildTestYARPCProcedures(params buildTestYARPCProceduresParams) []transport.Procedure {
	handler := &_TestYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.encoding.protobuf.Test",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Unary",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Unary,
							NewRequest:  newTestServiceUnaryYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "Duplex",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.Duplex,
						},
					),
				},
			},
		},
	)
}

// BuildTestYARPCProcedures prepares an implementation of the Test service for YARPC registration.
func BuildTestYARPCProcedures(server TestYARPCServer) []transport.Procedure {
	return buildTestYARPCProcedures(buildTestYARPCProceduresParams{Server: server})
}

// FxTestYARPCClientParams defines the input
// for NewFxTestYARPCClient. It provides the
// paramaters to get a TestYARPCClient in an
// Fx application.
type FxTestYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxTestYARPCClientResult defines the output
// of NewFxTestYARPCClient. It provides a
// TestYARPCClient to an Fx application.
type FxTestYARPCClientResult struct {
	fx.Out

	Client TestYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxTestYARPCClient provides a TestYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    testpb.NewFxTestYARPCClient("service-name"),
//    ...
//  )
func NewFxTestYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxTestYARPCClientParams) FxTestYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxTestYARPCClientResult{
			Client: newTestYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxTestYARPCProceduresParams defines the input
// for NewFxTestYARPCProcedures. It provides the
// paramaters to get TestYARPCServer procedures in an
// Fx application.
type FxTestYARPCProceduresParams struct {
	fx.In

	Server      TestYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxTestYARPCProceduresResult defines the output
// of NewFxTestYARPCProcedures. It provides
// TestYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxTestYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxTestYARPCProcedures provides TestYARPCServer procedures to an Fx application.
// It expects a TestYARPCServer to be present in the container.
//
//  fx.Provide(
//    testpb.NewFxTestYARPCProcedures(),
//    ...
//  )
func NewFxTestYARPCProcedures() interface{} {
	return func(params FxTestYARPCProceduresParams) FxTestYARPCProceduresResult {
		return FxTestYARPCProceduresResult{
			Procedures: buildTestYARPCProcedures(buildTestYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.yarpc.encoding.protobuf.Test",
				FileDescriptors: yarpcFileDescriptorClosurec161fcfdc0c3ff1e,
			},
		}
	}
}

type _TestYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_TestYARPCCaller) Unary(ctx context.Context, request *TestMessage, options ...yarpc.CallOption) (*TestMessage, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Unary", request, newTestServiceUnaryYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceUnaryYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_TestYARPCCaller) Duplex(ctx context.Context, options ...yarpc.CallOption) (TestServiceDuplexYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "Duplex", options...)
	if err != nil {
		return nil, err
	}
	return &_TestServiceDuplexYARPCClient{stream: stream}, nil
}

type _TestYARPCHandler struct {
	server TestYARPCServer
}

func (h *_TestYARPCHandler) Unary(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *TestMessage
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*TestMessage)
		if !ok {
			return nil, protobuf.CastError(emptyTestServiceUnaryYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Unary(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_TestYARPCHandler) Duplex(serverStream *protobuf.ServerStream) error {
	return h.server.Duplex(&_TestServiceDuplexYARPCServer{serverStream: serverStream})
}

type _TestServiceDuplexYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_TestServiceDuplexYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_TestServiceDuplexYARPCClient) Send(request *TestMessage, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_TestServiceDuplexYARPCClient) Recv(options ...yarpc.StreamOption) (*TestMessage, error) {
	responseMessage, err := c.stream.Receive(newTestServiceDuplexYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceDuplexYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_TestServiceDuplexYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _TestServiceDuplexYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_TestServiceDuplexYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_TestServiceDuplexYARPCServer) Recv(options ...yarpc.StreamOption) (*TestMessage, error) {
	requestMessage, err := s.serverStream.Receive(newTestServiceDuplexYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*TestMessage)
	if !ok {
		return nil, protobuf.CastError(emptyTestServiceDuplexYARPCRequest, requestMessage)
	}
	return request, err
}

func (s *_TestServiceDuplexYARPCServer) Send(response *TestMessage, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

func newTestServiceUnaryYARPCRequest() proto.Message {
	return &TestMessage{}
}

func newTestServiceUnaryYARPCResponse() proto.Message {
	return &TestMessage{}
}

func newTestServiceDuplexYARPCRequest() proto.Message {
	return &TestMessage{}
}

func newTestServiceDuplexYARPCResponse() proto.Message {
	return &TestMessage{}
}

var (
	emptyTestServiceUnaryYARPCRequest   = &TestMessage{}
	emptyTestServiceUnaryYARPCResponse  = &TestMessage{}
	emptyTestServiceDuplexYARPCRequest  = &TestMessage{}
	emptyTestServiceDuplexYARPCResponse = &TestMessage{}
)

var yarpcFileDescriptorClosurec161fcfdc0c3ff1e = [][]byte{
	// test.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
		0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0x4d, 0x4a, 0x2d, 0xd2, 0xab, 0x4c, 0x2c,
		0x2a, 0x48, 0xd6, 0x4b, 0xcd, 0x4b, 0xce, 0x4f, 0xc9, 0xcc, 0x4b, 0x87, 0x48, 0x25, 0x95, 0xa6,
		0x29, 0x29, 0x73, 0x71, 0x87, 0xa4, 0x16, 0x97, 0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a,
		0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41,
		0x38, 0x46, 0x27, 0x19, 0xb9, 0x58, 0x40, 0xaa, 0x84, 0x62, 0xb9, 0x58, 0x43, 0xf3, 0x12, 0x8b,
		0x2a, 0x85, 0x34, 0xf5, 0xf0, 0x99, 0xaa, 0x87, 0x64, 0xa4, 0x14, 0xf1, 0x4a, 0x85, 0x92, 0xb8,
		0xd8, 0x5c, 0x4a, 0x0b, 0x72, 0x52, 0x2b, 0x68, 0x63, 0xbe, 0x06, 0xa3, 0x01, 0xa3, 0x13, 0x47,
		0x14, 0x1b, 0x28, 0x70, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x6a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
		0xff, 0x6f, 0x29, 0xe9, 0xdd, 0x2d, 0x01, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) TestYARPCClient {
			return NewTestYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}

func init() {

	proto.RegisterType((*TestMessage)(nil), "uber.yarpc.encoding.protobuf.TestMessage")

}
