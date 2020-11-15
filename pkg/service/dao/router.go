package dao

import (
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"github.com/jinzhu/gorm"
)

type RouteDao struct {
	Dao
}

func NewRoute(db *gorm.DB) *RouteDao {
	return &RouteDao{Dao{db: db}}
}

func (dao RouteDao) Save(route entity.Routers) entity.Routers {
	dao.db.Save(&route)
	return route
}

func (dao RouteDao) GetOne(path string, method string) entity.Routers {
	var route entity.Routers
	dao.db.Table(entity.TABLE_ROUTERS).
		Where(entity.Routers{Path: path, Method: method}).
		First(&route)
	return route
}

func (dao RouteDao) GetById(ids []int) []entity.Routers {
	var route []entity.Routers
	dao.db.Table(entity.TABLE_ROUTERS).Debug().Where("id in (?)", ids).Scan(&route)
	return route
}
