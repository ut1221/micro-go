import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/micro";

// 查询用户列表
export function listUser(data) {
  return request({
    url: '/v1/user/list',
    method: 'post',
    data: data
  })
}

// 查询用户详细
export function getUser(userId) {
  return request({
    url: '/v1/user/info/' + parseStrEmpty(userId),
    method: 'get'
  })
}

// 查询用户角色部门
export function getOtherUser() {
  return request({
    url: '/v1/user/otherInfo',
    method: 'get'
  })
}

// 新增用户
export function addUser(data) {
  return request({
    url: '/v1/user/save',
    method: 'post',
    data: data
  })
}

// 修改用户
export function updateUser(data) {
  return request({
    url: '/v1/user/update',
    method: 'put',
    data: data
  })
}

// 删除用户
export function delUser(userId) {
  return request({
    url: '/v1/user/delete/' + userId,
    method: 'delete'
  })
}

// 用户密码重置
export function resetUserPwd(userId, password) {
  const data = {
    userId,
    password
  }
  return request({
    url: '/v1/user/resetPwd',
    method: 'put',
    data: data
  })
}

// 用户状态修改
export function changeUserStatus(userId, status) {
  const data = {
    userId,
    status
  }
  return request({
    url: '/system/user/changeStatus',
    method: 'put',
    data: data
  })
}

// 查询用户个人信息
export function getUserProfile() {
  return request({
    url: '/v1/user/profile',
    method: 'get'
  })
}

// 修改用户个人信息
export function updateUserProfile(data) {
  return request({
    url: '/system/user/profile',
    method: 'put',
    data: data
  })
}

// 用户密码重置
export function updateUserPwd(oldPassword, newPassword) {
  const data = {
    oldPassword,
    newPassword
  }
  return request({
    url: '/system/user/profile/updatePwd',
    method: 'put',
    params: data
  })
}

// 用户头像上传
export function uploadAvatar(data) {
  return request({
    url: '/system/user/profile/avatar',
    method: 'post',
    data: data
  })
}

// 查询授权角色
export function getAuthRole(userId) {
  return request({
    url: '/v1/user/authRole/' + userId,
    method: 'get'
  })
}

// 保存授权角色
export function updateAuthRole(data) {
  return request({
    url: '/v1/user/authRole',
    method: 'put',
    data: data
  })
}

// 查询部门下拉树结构
export function deptTreeSelect() {
  return request({
    url: '/v1/dept/deptTree',
    method: 'get'
  })
}
