package test

import (
	"context"
	"testing"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	errorAssert "github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (t *GetUserArticlesHeadKey) TestGetSuccess() {
	t.req = &pb.GetUserArticlesHeadKeyReq{
		UserId: "01ab119f-5345-46e4-bac2-a1d0698c40ec",
	}
	t.Request()

	if assert.NoError(t.T(), t.err) {
		t.res.ArticlePageHeadKey = "7e0156f8-38e5-45b6-9801-844bda55f270"
	}
}

func (t *GetUserArticlesHeadKey) TestInternalError() {
	t.req = &pb.GetUserArticlesHeadKeyReq{
		UserId: "01ab119f-3434-46e4-bac2-a1d0698c40ec",
	}
	t.Request()

	errorAssert.IsGrpcInternalError(t.T(), t.err)
}

type GetUserArticlesHeadKey struct {
	suite.Suite
	req *pb.GetUserArticlesHeadKeyReq
	res *pb.GetUserArticlesHeadKeyRes
	err error
}

func TestGetUserArticlesHeadKeyInit(t *testing.T) {
	suite.Run(t, new(GetUserArticlesHeadKey))
}

func (t *GetUserArticlesHeadKey) Request() {
	t.res, t.err = UserClient.GetUserArticlesHeadKey(context.Background(), t.req)
}

func (t *GetUserArticlesHeadKey) SetupSuite() {
	ResetDB()

	t.err = nil
	t.req = &pb.GetUserArticlesHeadKeyReq{}
	t.res = &pb.GetUserArticlesHeadKeyRes{}
}
