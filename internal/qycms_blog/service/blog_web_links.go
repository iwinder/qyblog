package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/web/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

func (b *BlogWebApiService) ListQyWebIndexLinks(ctx context.Context, request *v1.ListQyWebLinksRequest) (*v1.ListQyWebLinksReply, error) {

	objList, err := b.lu.FindIndexLinkAllWitchCache(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]*v1.WebLinkInfo, 0, len(objList))
	for _, item := range objList {
		titem := bizToWebLinkResponse(item)
		objs = append(objs, titem)
	}
	return &v1.ListQyWebLinksReply{Items: objs}, nil
}

func (b *BlogWebApiService) ListQyWebLinks(ctx context.Context, request *v1.ListQyWebLinksRequest) (*v1.ListQyWebLinksReply, error) {
	objList, err := b.lu.FindAllWitchCache(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]*v1.WebLinkInfo, 0, len(objList))
	for _, item := range objList {
		titem := bizToWebLinkResponse(item)
		objs = append(objs, titem)
	}
	return &v1.ListQyWebLinksReply{Items: objs}, nil
}

func (b *BlogWebApiService) ListQyWebShortLinks(ctx context.Context, request *v1.ListQyWebLinksRequest) (*v1.ListQyWebShortLinksReply, error) {

	objList, err := b.slu.FindAllWitchCache(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]*v1.WebShortLinkInfo, 0, len(objList))
	for _, item := range objList {
		titem := bizToWebShortLinkResponse(item)
		objs = append(objs, titem)
	}
	return &v1.ListQyWebShortLinksReply{Items: objs}, nil
}

func bizToWebLinkResponse(obj *biz.LinkDO) *v1.WebLinkInfo {
	objInfoRsp := &v1.WebLinkInfo{
		Id:          obj.ID,
		Name:        obj.Name,
		Url:         obj.Url,
		Description: obj.Description,
	}
	return objInfoRsp
}

func bizToWebShortLinkResponse(obj *biz.ShortLinkDO) *v1.WebShortLinkInfo {
	objInfoRsp := &v1.WebShortLinkInfo{
		Url:        obj.Url,
		Identifier: obj.Identifier,
	}
	return objInfoRsp
}
