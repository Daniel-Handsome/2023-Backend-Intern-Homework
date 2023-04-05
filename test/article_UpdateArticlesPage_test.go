package test

import (
	"context"
	"testing"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	testDB "github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (t *UpdateArticlesPage) TestIDColumnGetSuccess() {
	t.req = &pb.UpdateArticlesPageReq{
		OrderColumns: 0,
	}
	t.Request()

	if assert.NoError(t.T(), t.err) {
		testDB.Exist(t.T(), "page_nodes", []testDB.Condition{
			{Col: `id`, Operator: "=", Val: "1"},
			{Col: `article_ids`, Operator: "&&", Val: "{1,2}"},
		})
	}
}

func (t *UpdateArticlesPage) TestCreateColumnGetSuccess() {
	t.req = &pb.UpdateArticlesPageReq{
		OrderColumns: 1,
	}
	t.Request()

	if assert.NoError(t.T(), t.err) {
		testDB.Exist(t.T(), "page_nodes", []testDB.Condition{
			{Col: `id`, Operator: "=", Val: "1"},
			{Col: `article_ids`, Operator: "&&", Val: "{4,1}"},
		})
	}
}

//func (t *UpdateArticlesPage) TestInternalError() {
//	t.req = &pb.UpdateArticlesPageReq{
//		HeadKey: "11ab119f-3434-46e4-bac2-a1d0698c40ec",
//	}
//	t.Request()
//
//	errorAssert.IsGrpcInternalError(t.T(), t.err)
//}

type UpdateArticlesPage struct {
	suite.Suite
	req *pb.UpdateArticlesPageReq
	res *pb.UpdateArticlesPageRes
	err error
}

func TestUpdateArticlesPageInit(t *testing.T) {
	suite.Run(t, new(UpdateArticlesPage))
}

func (t *UpdateArticlesPage) Request() {
	t.res, t.err = ArticleClient.UpdateArticlesPage(context.Background(), t.req)
}

func (t *UpdateArticlesPage) SetupSuite() {
	ResetDB()

	t.err = nil
	t.req = &pb.UpdateArticlesPageReq{}
	t.res = &pb.UpdateArticlesPageRes{}
}
