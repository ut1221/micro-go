import request from '@/utils/request'

// 查询字典类型列表
export function listType(data) {
  return request({
    url: '/v1/dict/type/list',
    method: 'post',
    data: data
  })
}

// 查询字典类型详细
export function getType(dictId) {
  return request({
    url: '/v1/dict/type/info/' + dictId,
    method: 'get'
  })
}

// 新增字典类型
export function addType(data) {
  return request({
    url: '/v1/dict/type/save',
    method: 'post',
    data: data
  })
}

// 修改字典类型
export function updateType(data) {
  return request({
    url: '/v1/dict/type/update',
    method: 'put',
    data: data
  })
}

// 删除字典类型
export function delType(dictId) {
  return request({
    url: '/v1/dict/type/delete/' + dictId,
    method: 'delete'
  })
}

// 刷新字典缓存
export function refreshCache() {
  return request({
    url: '/v1/dict/type/refresh',
    method: 'get'
  })
}

// 获取字典选择框列表
export function optionselect() {
  return request({
    url: '/v1/dict/type/option/select',
    method: 'get'
  })
}
