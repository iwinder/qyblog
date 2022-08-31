package db

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"strings"
	"sync"
)

type CasbinData struct {
	Enf *casbin.SyncedEnforcer
	log *log.Helper
}

var (
	casbinData *CasbinData
	conce      sync.Once
)

// NewCasbinData .
func NewCasbinData(data *Data, conf *conf.Auth, logger log.Logger) (*CasbinData, error) {
	if strings.EqualFold(conf.Casbin.ModelPath, "") && data.Db == nil {
		log.Fatal("Casbin Model配置信息缺失")
	}
	var err error
	var m model.Model
	var e *casbin.SyncedEnforcer
	l := log.NewHelper(log.With(logger, "module", "mysql/casbin"))
	conce.Do(func() {
		Apter, _ := gormadapter.NewAdapterByDBWithCustomTable(data.Db, &po.CasbinRulePO{})
		// 从 .CONF 文件中加载 model
		m, err = model.NewModelFromFile(conf.Casbin.ModelPath)
		if err != nil {
			log.Fatalf("Casbin Model 配置失败" + err.Error())
		}
		e, err = casbin.NewSyncedEnforcer(m, Apter)
		err = e.LoadPolicy()
		if err != nil {
			log.Fatalf("Casbin 执行者创建失败" + err.Error())
		}
		casbinData = &CasbinData{
			Enf: e,
			log: l,
		}

	})

	if casbinData.Enf == nil {
		log.Fatalf("Casbin 执行者初始化失败: %v", err)
	}
	return casbinData, nil
}
