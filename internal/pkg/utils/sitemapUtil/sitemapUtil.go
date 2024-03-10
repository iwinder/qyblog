package sitemapUtil

import (
	"github.com/douyacun/gositemap"
)

func NewSiteMapCreator(list []string) {
	st := gositemap.NewSiteMap()
	st.SetPretty(true)
	st.ToXml()
}
