package biz

import (
	"context"
	"github.com/douyacun/gositemap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/sitemapUtil"
	"strings"
	"time"
)

type SiteMapUsecase struct {
	log  *log.Helper
	au   *ArticleUsecase
	site *SiteConfigUsecase
}

func NewSiteMapUsecase(logger log.Logger, site *SiteConfigUsecase,
	au *ArticleUsecase,
) *SiteMapUsecase {
	return &SiteMapUsecase{log: log.NewHelper(logger),
		au: au, site: site}
}

// GeneratorMap 生成地图 TODO: 后期考虑文件的问题
func (uc *SiteMapUsecase) GeneratorMap(ctx context.Context, path string) error {
	mapIndex := gositemap.NewSiteMapIndex() // sitemap_index.xml
	siteUrl := uc.site.FindValueByKey(ctx, "site_url")
	if !strings.HasSuffix(siteUrl, "/") {
		siteUrl += "/"
	}
	postFile := uc.generatorPostXml(ctx, 1, siteUrl, path)
	pageFile := uc.generatorPostXml(ctx, 2, siteUrl, path)
	mapIndex.Append(siteUrl + postFile)
	mapIndex.Append(siteUrl + pageFile)

	_, err := mapIndex.Storage(path + "sitemap_index.xml")
	if err != nil {
		return err
	}
	return nil
}

func (uc *SiteMapUsecase) generatorPostXml(ctx context.Context, atype int, siteUrl string, path string) string {
	fileName := sitemapUtil.SITE_MAP_FILE_PAGE_NAME
	if atype == 1 {
		fileName = sitemapUtil.SITE_MAP_FILE_POST_NAME
	}
	st1 := gositemap.NewSiteMap()
	st1.SetPretty(true)
	st1.SetPublicPath(path)
	st1.SetFilename(fileName)
	st1.SetCompress(false)
	opts := ArticleDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = 1
	opts.ListOptions.PageSize = 100
	opts.StatusFlag = 2
	opts.Atype = atype // TODO: 暂且指定死，后期有其它需要前端展示列表的再放开
	opts.ListOptions.Init()
	dataDOs, err := uc.au.GeneratorMapListAll(ctx, opts)
	if err != nil {
		uc.log.WithContext(ctx).Error("更新站点地图失败：%w", err)
	}

	url := gositemap.NewUrl()
	if atype == 1 {
		url.SetLoc(siteUrl)
		url.SetLastmod(time.Now())
		st1.AppendUrl(url)
	}
	for _, dataDo := range dataDOs.Items {
		url = gositemap.NewUrl()
		url.SetLoc(siteUrl + dataDo.PermaLink)
		url.SetLastmod(dataDo.UpdatedAt.Local())
		st1.AppendUrl(url)

	}
	opts.TotalCount = dataDOs.TotalCount
	opts.ListOptions.IsLast()
	totalPage := opts.Pages
	var page int64
	for page = 2; page <= totalPage; page++ {
		opts.Current = page
		opts.ListOptions.Init()
		dataDOs, err = uc.au.GeneratorMapListAll(ctx, opts)
		if err != nil {
			uc.log.WithContext(ctx).Error("更新站点地图失败：%w", err)
		}
		for _, dataDo := range dataDOs.Items {
			url = gositemap.NewUrl()
			url.SetLoc(siteUrl + dataDo.PermaLink)
			url.SetLastmod(dataDo.UpdatedAt.Local())
			st1.AppendUrl(url)
		}
	}

	st1Filename, aerr := st1.Storage()
	if aerr != nil {
		uc.log.WithContext(ctx).Error("更新站点地图失败：%w", err)
		return ""
	}
	return st1Filename
}
