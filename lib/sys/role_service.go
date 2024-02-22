package sys

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

type RoleService struct {
	db       *gorm.DB
	enforcer *casbin.SyncedEnforcer
}

func NewUserRoleService(db *gorm.DB, e *casbin.SyncedEnforcer) *RoleService {
	return &RoleService{db: db, enforcer: e}
}

// 拼接角色ID，为了防止角色与用户名冲突
func (my *RoleService) makeRoleName(roleId base.Id) string {
	return fmt.Sprintf("role_%d", roleId)
}

// AddPolicy 添加策略
func (my *RoleService) AddPolicy(roleId base.Id, uri, method string) (bool, error) {
	return my.enforcer.AddPolicy(my.makeRoleName(roleId), uri, method)
}

// AddPolicies 批量添加策略
func (my *RoleService) AddPolicies(rules [][]string) (bool, error) {
	return my.enforcer.AddPolicies(rules)
}

// DeleteRole 删除角色对应的用户和权限
func (my *RoleService) DeleteRole(roleId base.Id) (bool, error) {
	return my.enforcer.DeleteRole(my.makeRoleName(roleId))
}

// DeleteRolePolicy 删除角色下的权限
func (my *RoleService) DeleteRolePolicy(roleId base.Id) (bool, error) {
	return my.enforcer.RemoveFilteredNamedPolicy("p", 0, my.makeRoleName(roleId))
}

// DeleteRoleUser 删除添加用户
func (my *RoleService) DeleteRoleUser(roleId base.Id) (bool, error) {
	return my.enforcer.RemoveFilteredNamedGroupingPolicy("g", 1, my.makeRoleName(roleId))
}

// DeleteUserRole 删除用户的角色信息
func (my *RoleService) DeleteUserRole(user string) (bool, error) {
	return my.enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, user)
}

// AddUserRole 添加角色和用户对应关系
func (my *RoleService) AddUserRole(user string, roleId base.Id) (bool, error) {
	return my.enforcer.AddGroupingPolicy(user, my.makeRoleName(roleId))
}

// AddUserRoles 批量添加角色和用户对应关联
func (my *RoleService) AddUserRoles(usernames []string, roleIds []base.Id) (bool, error) {
	rules := make([][]string, 0)
	for _, u := range usernames {
		for _, r := range roleIds {
			rules = append(rules, []string{u, my.makeRoleName(r)})
		}
	}
	return my.enforcer.AddGroupingPolicies(rules)
}
