// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"

	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type RealWorldHTTPServer interface {
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentResponse, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentResponse, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleResponse, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticleResponse, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileResponse, error)
	GetArticles(context.Context, *GetArticleRequest) (*SingleArticleResponse, error)
	GetComment(context.Context, *GetCommentRequest) (*MultipleCommentResponse, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileResponse, error)
	GetTags(context.Context, *GetTagsRequest) (*TagListResponse, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticleResponse, error)
	Login(context.Context, *LoginRequest) (*UserResponse, error)
	Register(context.Context, *RegisterRequest) (*UserResponse, error)
	UnfavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*SingleArticleResponse, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error)
}

func RegisterRealWorldHTTPServer(s *http.Server, srv RealWorldHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/login", _RealWorld_Login0_HTTP_Handler(srv))
	r.POST("/api/users", _RealWorld_Register0_HTTP_Handler(srv))
	r.GET("/api/user", _RealWorld_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/api/user", _RealWorld_UpdateUser0_HTTP_Handler(srv))
	r.GET("/api/profiles/{username}", _RealWorld_GetProfile0_HTTP_Handler(srv))
	r.POST("/api/profiles/{username}/follow", _RealWorld_FollowUser0_HTTP_Handler(srv))
	r.DELETE("/api/profiles/{username}/follow", _RealWorld_UnfollowUser0_HTTP_Handler(srv))
	r.GET("/api/articles", _RealWorld_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _RealWorld_FeedArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}", _RealWorld_GetArticles0_HTTP_Handler(srv))
	r.POST("/api/articles", _RealWorld_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/{slug}", _RealWorld_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}", _RealWorld_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/comments", _RealWorld_AddComment0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}/comments", _RealWorld_GetComment0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/comments/{id}", _RealWorld_DeleteComment0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/favorite", _RealWorld_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/favorite", _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv))
	r.GET("/api/tags", _RealWorld_GetTags0_HTTP_Handler(srv))
}

func _RealWorld_Login0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*UserResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_Register0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*UserResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_GetCurrentUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/GetCurrentUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*UserResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_UpdateUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*UserResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_GetProfile0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/GetProfile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*ProfileResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_FollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/FollowUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowUser(ctx, req.(*FollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*ProfileResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_UnfollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfollowUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/UnfollowUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfollowUser(ctx, req.(*UnfollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*ProfileResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_ListArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/ListArticles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*MultipleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_FeedArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/FeedArticles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*MultipleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_GetArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/GetArticles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticles(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_CreateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/CreateArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_UpdateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/UpdateArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_DeleteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/DeleteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_AddComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/AddComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComment(ctx, req.(*AddCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleCommentResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_GetComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/GetComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetComment(ctx, req.(*GetCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*MultipleCommentResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_DeleteComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/DeleteComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleCommentResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_FavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/FavoriteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfavoriteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/UnfavoriteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfavoriteArticle(ctx, req.(*UnfavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*SingleArticleResponse)
		return ctx.Result(200, Response)
	}
}

func _RealWorld_GetTags0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.RealWorld/GetTags")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		Response := out.(*TagListResponse)
		return ctx.Result(200, Response)
	}
}

type RealWorldHTTPClient interface {
	AddComment(ctx context.Context, req *AddCommentRequest, opts ...http.CallOption) (rsp *SingleCommentResponse, err error)
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *SingleCommentResponse, err error)
	FavoriteArticle(ctx context.Context, req *FavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	FeedArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticleResponse, err error)
	FollowUser(ctx context.Context, req *FollowUserRequest, opts ...http.CallOption) (rsp *ProfileResponse, err error)
	GetArticles(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	GetComment(ctx context.Context, req *GetCommentRequest, opts ...http.CallOption) (rsp *MultipleCommentResponse, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *UserResponse, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *ProfileResponse, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *TagListResponse, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticleResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserResponse, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserResponse, err error)
	UnfavoriteArticle(ctx context.Context, req *UnfavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	UnfollowUser(ctx context.Context, req *UnfollowUserRequest, opts ...http.CallOption) (rsp *ProfileResponse, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserResponse, err error)
}

type RealWorldHTTPClientImpl struct {
	cc *http.Client
}

func NewRealWorldHTTPClient(client *http.Client) RealWorldHTTPClient {
	return &RealWorldHTTPClientImpl{client}
}

func (c *RealWorldHTTPClientImpl) AddComment(ctx context.Context, in *AddCommentRequest, opts ...http.CallOption) (*SingleCommentResponse, error) {
	var out SingleCommentResponse
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/AddComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/CreateArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/DeleteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*SingleCommentResponse, error) {
	var out SingleCommentResponse
	pattern := "/api/articles/{slug}/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/DeleteComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/FavoriteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*MultipleArticleResponse, error) {
	var out MultipleArticleResponse
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/FeedArticles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...http.CallOption) (*ProfileResponse, error) {
	var out ProfileResponse
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/FollowUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetArticles(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/GetArticles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetComment(ctx context.Context, in *GetCommentRequest, opts ...http.CallOption) (*MultipleCommentResponse, error) {
	var out MultipleCommentResponse
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/GetComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*UserResponse, error) {
	var out UserResponse
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/GetCurrentUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*ProfileResponse, error) {
	var out ProfileResponse
	pattern := "/api/profiles/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/GetProfile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*TagListResponse, error) {
	var out TagListResponse
	pattern := "/api/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/GetTags"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*MultipleArticleResponse, error) {
	var out MultipleArticleResponse
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/ListArticles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserResponse, error) {
	var out UserResponse
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserResponse, error) {
	var out UserResponse
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/UnfavoriteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...http.CallOption) (*ProfileResponse, error) {
	var out ProfileResponse
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/UnfollowUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*SingleArticleResponse, error) {
	var out SingleArticleResponse
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/UpdateArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserResponse, error) {
	var out UserResponse
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.RealWorld/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
