package jbiz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewCountCommentJobRepo, NewSiteMapJobRepo, NewPostVCountJobRepo, NewEmailSendJobRepo)
