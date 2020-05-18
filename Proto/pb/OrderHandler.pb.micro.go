// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: OrderHandler.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderService service

type OrderService interface {
	// 创建订单
	Created(ctx context.Context, in *OrderCreatedRequest, opts ...client.CallOption) (*OrderCreatedResponse, error)
	// 更新订单
	Updated(ctx context.Context, in *OrderUpdatedRequest, opts ...client.CallOption) (*OrderUpdatedResponse, error)
	// 查询
	Find(ctx context.Context, in *OrderFindRequest, opts ...client.CallOption) (*OrderFindResponse, error)
	Select(ctx context.Context, in *OrderSelectRequest, opts ...client.CallOption) (*OrderSelectResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pb"
	}
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Created(ctx context.Context, in *OrderCreatedRequest, opts ...client.CallOption) (*OrderCreatedResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Created", in)
	out := new(OrderCreatedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Updated(ctx context.Context, in *OrderUpdatedRequest, opts ...client.CallOption) (*OrderUpdatedResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Updated", in)
	out := new(OrderUpdatedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Find(ctx context.Context, in *OrderFindRequest, opts ...client.CallOption) (*OrderFindResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Find", in)
	out := new(OrderFindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Select(ctx context.Context, in *OrderSelectRequest, opts ...client.CallOption) (*OrderSelectResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Select", in)
	out := new(OrderSelectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	// 创建订单
	Created(context.Context, *OrderCreatedRequest, *OrderCreatedResponse) error
	// 更新订单
	Updated(context.Context, *OrderUpdatedRequest, *OrderUpdatedResponse) error
	// 查询
	Find(context.Context, *OrderFindRequest, *OrderFindResponse) error
	Select(context.Context, *OrderSelectRequest, *OrderSelectResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		Created(ctx context.Context, in *OrderCreatedRequest, out *OrderCreatedResponse) error
		Updated(ctx context.Context, in *OrderUpdatedRequest, out *OrderUpdatedResponse) error
		Find(ctx context.Context, in *OrderFindRequest, out *OrderFindResponse) error
		Select(ctx context.Context, in *OrderSelectRequest, out *OrderSelectResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) Created(ctx context.Context, in *OrderCreatedRequest, out *OrderCreatedResponse) error {
	return h.OrderServiceHandler.Created(ctx, in, out)
}

func (h *orderServiceHandler) Updated(ctx context.Context, in *OrderUpdatedRequest, out *OrderUpdatedResponse) error {
	return h.OrderServiceHandler.Updated(ctx, in, out)
}

func (h *orderServiceHandler) Find(ctx context.Context, in *OrderFindRequest, out *OrderFindResponse) error {
	return h.OrderServiceHandler.Find(ctx, in, out)
}

func (h *orderServiceHandler) Select(ctx context.Context, in *OrderSelectRequest, out *OrderSelectResponse) error {
	return h.OrderServiceHandler.Select(ctx, in, out)
}
