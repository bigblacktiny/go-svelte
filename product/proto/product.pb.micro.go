// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product/proto/product.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "goTemp/globalProtos"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ProductSrv service

func NewProductSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductSrv service

type ProductSrvService interface {
	GetProductById(ctx context.Context, in *SearchId, opts ...client.CallOption) (*Product, error)
	GetProducts(ctx context.Context, in *SearchParams, opts ...client.CallOption) (*Products, error)
	CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	DeleteProduct(ctx context.Context, in *SearchId, opts ...client.CallOption) (*Response, error)
	BeforeCreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error)
	BeforeUpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error)
	BeforeDeleteProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error)
	AfterCreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error)
	AfterUpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error)
	AfterDeleteProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error)
}

type productSrvService struct {
	c    client.Client
	name string
}

func NewProductSrvService(name string, c client.Client) ProductSrvService {
	return &productSrvService{
		c:    c,
		name: name,
	}
}

func (c *productSrvService) GetProductById(ctx context.Context, in *SearchId, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.GetProductById", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) GetProducts(ctx context.Context, in *SearchParams, opts ...client.CallOption) (*Products, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.GetProducts", in)
	out := new(Products)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.CreateProduct", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) UpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.UpdateProduct", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) DeleteProduct(ctx context.Context, in *SearchId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.DeleteProduct", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) BeforeCreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.BeforeCreateProduct", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) BeforeUpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.BeforeUpdateProduct", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) BeforeDeleteProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ValidationErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.BeforeDeleteProduct", in)
	out := new(ValidationErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) AfterCreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.AfterCreateProduct", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) AfterUpdateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.AfterUpdateProduct", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productSrvService) AfterDeleteProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*AfterFuncErr, error) {
	req := c.c.NewRequest(c.name, "ProductSrv.AfterDeleteProduct", in)
	out := new(AfterFuncErr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductSrv service

type ProductSrvHandler interface {
	GetProductById(context.Context, *SearchId, *Product) error
	GetProducts(context.Context, *SearchParams, *Products) error
	CreateProduct(context.Context, *Product, *Response) error
	UpdateProduct(context.Context, *Product, *Response) error
	DeleteProduct(context.Context, *SearchId, *Response) error
	BeforeCreateProduct(context.Context, *Product, *ValidationErr) error
	BeforeUpdateProduct(context.Context, *Product, *ValidationErr) error
	BeforeDeleteProduct(context.Context, *Product, *ValidationErr) error
	AfterCreateProduct(context.Context, *Product, *AfterFuncErr) error
	AfterUpdateProduct(context.Context, *Product, *AfterFuncErr) error
	AfterDeleteProduct(context.Context, *Product, *AfterFuncErr) error
}

func RegisterProductSrvHandler(s server.Server, hdlr ProductSrvHandler, opts ...server.HandlerOption) error {
	type productSrv interface {
		GetProductById(ctx context.Context, in *SearchId, out *Product) error
		GetProducts(ctx context.Context, in *SearchParams, out *Products) error
		CreateProduct(ctx context.Context, in *Product, out *Response) error
		UpdateProduct(ctx context.Context, in *Product, out *Response) error
		DeleteProduct(ctx context.Context, in *SearchId, out *Response) error
		BeforeCreateProduct(ctx context.Context, in *Product, out *ValidationErr) error
		BeforeUpdateProduct(ctx context.Context, in *Product, out *ValidationErr) error
		BeforeDeleteProduct(ctx context.Context, in *Product, out *ValidationErr) error
		AfterCreateProduct(ctx context.Context, in *Product, out *AfterFuncErr) error
		AfterUpdateProduct(ctx context.Context, in *Product, out *AfterFuncErr) error
		AfterDeleteProduct(ctx context.Context, in *Product, out *AfterFuncErr) error
	}
	type ProductSrv struct {
		productSrv
	}
	h := &productSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductSrv{h}, opts...))
}

type productSrvHandler struct {
	ProductSrvHandler
}

func (h *productSrvHandler) GetProductById(ctx context.Context, in *SearchId, out *Product) error {
	return h.ProductSrvHandler.GetProductById(ctx, in, out)
}

func (h *productSrvHandler) GetProducts(ctx context.Context, in *SearchParams, out *Products) error {
	return h.ProductSrvHandler.GetProducts(ctx, in, out)
}

func (h *productSrvHandler) CreateProduct(ctx context.Context, in *Product, out *Response) error {
	return h.ProductSrvHandler.CreateProduct(ctx, in, out)
}

func (h *productSrvHandler) UpdateProduct(ctx context.Context, in *Product, out *Response) error {
	return h.ProductSrvHandler.UpdateProduct(ctx, in, out)
}

func (h *productSrvHandler) DeleteProduct(ctx context.Context, in *SearchId, out *Response) error {
	return h.ProductSrvHandler.DeleteProduct(ctx, in, out)
}

func (h *productSrvHandler) BeforeCreateProduct(ctx context.Context, in *Product, out *ValidationErr) error {
	return h.ProductSrvHandler.BeforeCreateProduct(ctx, in, out)
}

func (h *productSrvHandler) BeforeUpdateProduct(ctx context.Context, in *Product, out *ValidationErr) error {
	return h.ProductSrvHandler.BeforeUpdateProduct(ctx, in, out)
}

func (h *productSrvHandler) BeforeDeleteProduct(ctx context.Context, in *Product, out *ValidationErr) error {
	return h.ProductSrvHandler.BeforeDeleteProduct(ctx, in, out)
}

func (h *productSrvHandler) AfterCreateProduct(ctx context.Context, in *Product, out *AfterFuncErr) error {
	return h.ProductSrvHandler.AfterCreateProduct(ctx, in, out)
}

func (h *productSrvHandler) AfterUpdateProduct(ctx context.Context, in *Product, out *AfterFuncErr) error {
	return h.ProductSrvHandler.AfterUpdateProduct(ctx, in, out)
}

func (h *productSrvHandler) AfterDeleteProduct(ctx context.Context, in *Product, out *AfterFuncErr) error {
	return h.ProductSrvHandler.AfterDeleteProduct(ctx, in, out)
}
