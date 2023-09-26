package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserRoleUsecase, NewRoleMenusUsecase, NewRoleApiUsecase,
	NewRoleUsecase, NewUserUsecase, NewSiteConfigUsecase,
	NewApiUsecase, NewMenusAdminUsecase, NewApiGroupUsecase, NewCasbinRuleUsecase,
	NewFileLibTypeUsecase, NewFileLibConfigUsecase, NewFileLibUsecase,
	NewArticleUsecase, NewArticleContentUsecase,
	NewCommentAgentUsecase, NewCommentIndexUsecase, NewCommentContentUsecase,
	NewLinkUsecase, NewShortLinkUsecase, NewMenusUsecase, NewMenusAgentUsecase,
	NewTagsUsecase, NewCategoryUsecase, NewArticleTagsUsecase, NewCommentUsecase,
	NewSiteMapUsecase, NewArticleVisitorUsecase, NewArticleResourceUsecase,
)
