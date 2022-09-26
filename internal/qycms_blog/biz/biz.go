package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase,
	NewUserRoleUsecase,
	NewRoleUsecase, NewUserUsecase, NewApiUsecase, NewMenusAdminUsecase, NewCasbinRuleUsecase,
	NewArticleUsecase, NewArticleContentUsecase,
	NewCommentAgentUsecase, NewCommentIndexUsecase, NewCommentContentUsecase,
)
