import request from '@/utils/request'

// 查询部门列表
export function listDept(data) {
  return request({
    url: '/v1/dept/list',
    method: 'post',
    data: data === undefined ? {} : data
  })
}

// 查询部门列表（排除节点）
export function listDeptExcludeChild(deptId) {
  return request({
    url: '/v1/dept/list/exclude/' + deptId,
    method: 'get'
  })
}

// 查询部门详细
export function getDept(deptId) {
  return request({
    url: '/v1/dept/info/' + deptId,
    method: 'get'
  })
}

// 新增部门
export function addDept(data) {
  return request({
    url: '/v1/dept/save',
    method: 'post',
    data: data
  })
}

// 修改部门
export function updateDept(data) {
  return request({
    url: '/v1/dept/update',
    method: 'put',
    data: data
  })
}

// 删除部门
export function delDept(deptId) {
  return request({
    url: '/v1/dept/delete/' + deptId,
    method: 'delete'
  })
}
