package service

import (
	"context"

	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (Response *v1.ProfileResponse, err error) {
	Response = &v1.ProfileResponse{}

	return Response, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (Response *v1.ProfileResponse, err error) {
	Response = &v1.ProfileResponse{}

	return Response, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (Response *v1.ProfileResponse, err error) {
	Response = &v1.ProfileResponse{}

	return Response, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (Response *v1.SingleCommentResponse, err error) {
	Response = &v1.SingleCommentResponse{}

	return Response, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.AddCommentRequest) (Response *v1.MultipleCommentResponse, err error) {
	Response = &v1.MultipleCommentResponse{}

	return Response, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (Response *v1.SingleCommentResponse, err error) {
	Response = &v1.SingleCommentResponse{}

	return Response, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (Response *v1.MultipleArticleResponse, err error) {
	Response = &v1.MultipleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (Response *v1.MultipleArticleResponse, err error) {
	Response = &v1.MultipleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (Response *v1.TagListResponse, err error) {
	Response = &v1.TagListResponse{}

	return Response, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (Response *v1.SingleArticleResponse, err error) {
	Response = &v1.SingleArticleResponse{}

	return Response, nil
}
