package test

import (
	"context"
	"testing"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	testDB "github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/db"
	errorAssert "github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (t *GetArticlesPage) TestGetSuccess() {
	t.req = &pb.GetArticlesPageReq{
		HeadKey: "7e0156f8-38e5-45b6-9801-844bda55f270",
	}
	t.Request()

	if assert.NoError(t.T(), t.err) {
		testDB.Exist(t.T(), "page_nodes", []testDB.Condition{
			{Col: `uuid`, Operator: "=", Val: "7e0156f8-38e5-45b6-9801-844bda55f270"},
			{Col: `next`, Operator: "=", Val: "39988b34-e7ff-4cad-99e1-5fe25dc5259c"},
		})

		testDB.CountByPostgresArray(t.T(), 2, "article_ids", "page_nodes", []testDB.Condition{
			{Col: `uuid`, Operator: "=", Val: "7e0156f8-38e5-45b6-9801-844bda55f270"},
		})

	}
}

func (t *GetArticlesPage) TestInternalError() {
	t.req = &pb.GetArticlesPageReq{
		HeadKey: "11ab119f-3434-46e4-bac2-a1d0698c40ec",
	}
	t.Request()

	errorAssert.IsGrpcInternalError(t.T(), t.err)
}

type GetArticlesPage struct {
	suite.Suite
	req *pb.GetArticlesPageReq
	res *pb.GetArticlesPageRes
	err error
}

func TestGetArticlesPageInit(t *testing.T) {
	suite.Run(t, new(GetArticlesPage))
}

func (t *GetArticlesPage) Request() {
	t.res, t.err = ArticleClient.GetArticlesPage(context.Background(), t.req)
}

func (t *GetArticlesPage) SetupSuite() {
	ResetDB()

	t.err = nil
	t.req = &pb.GetArticlesPageReq{}
	t.res = &pb.GetArticlesPageRes{}
}
