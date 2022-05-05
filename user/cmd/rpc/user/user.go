// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmptyRequest          = user.EmptyRequest
	EmptyResponse         = user.EmptyResponse
	GetUsersResponse      = user.GetUsersResponse
	IdRequest             = user.IdRequest
	RegisterBatchResponse = user.RegisterBatchResponse
	RegisterRequest       = user.RegisterRequest
	RegisterResponse      = user.RegisterResponse
	SetSecretaryRequest   = user.SetSecretaryRequest
	TokenRequest          = user.TokenRequest
	TokenResponse         = user.TokenResponse
	UserResponse          = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		GetAllTeachers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
		GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		SetSecretary(ctx context.Context, in *SetSecretaryRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetAllTeachers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAllTeachers(ctx, in, opts...)
}

func (m *defaultUser) GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetToken(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) SetSecretary(ctx context.Context, in *SetSecretaryRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetSecretary(ctx, in, opts...)
}
