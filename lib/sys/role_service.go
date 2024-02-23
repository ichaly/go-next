package sys

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"gorm.io/gorm"
)

var roleService *RoleService

type RoleService struct {
	db       *gorm.DB
	enforcer *casbin.SyncedEnforcer
}

func NewUserRoleService(db *gorm.DB, e *casbin.SyncedEnforcer) *RoleService {
	roleService = &RoleService{db: db, enforcer: e}
	return roleService
}

// 拼接角色ID，为了防止角色与用户名冲突
func (my *RoleService) makeRoleName(roleId base.Id) string {
	return fmt.Sprintf("role_%d", roleId)
}

// AddRolePolicy 添加角色权限策略
func (my *RoleService) AddRolePolicy(roleId base.Id, params ...string) (bool, error) {
	return my.enforcer.AddPolicy(my.makeRoleName(roleId), params)
}

// AddUserPolicy 添加用户角色策略
func (my *RoleService) AddUserPolicy(userId base.Id, roleId base.Id) (bool, error) {
	user := util.FormatLong(int64(userId))
	return my.enforcer.AddGroupingPolicy(user, my.makeRoleName(roleId))
}

// DeleteRole 删除角色对应的用户和权限
func (my *RoleService) DeleteRole(roleId base.Id) (bool, error) {
	return my.enforcer.DeleteRole(my.makeRoleName(roleId))
}

// DeleteUser 删除用户
func (my *RoleService) DeleteUser(userId base.Id) (bool, error) {
	user := util.FormatLong(int64(userId))
	return my.enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, user)
}

// DeleteRule 删除角色下的权限
func (my *RoleService) DeleteRule(params ...string) (bool, error) {
	return my.enforcer.RemoveFilteredNamedPolicy("p", 1, params...)
}
