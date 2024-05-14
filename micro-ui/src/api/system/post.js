import request from '@/utils/request'

// 查询岗位列表
export function listPost(data) {
  return request({
    url: '/v1/post/list',
    method: 'post',
    data: data
  })
}

// 查询岗位详细
export function getPost(postId) {
  return request({
    url: '/v1/post/info/' + postId,
    method: 'get'
  })
}

// 新增岗位
export function addPost(data) {
  return request({
    url: '/v1/post/save',
    method: 'post',
    data: data
  })
}

// 修改岗位
export function updatePost(data) {
  return request({
    url: '/v1/post/update',
    method: 'put',
    data: data
  })
}

// 删除岗位
export function delPost(postId) {
  return request({
    url: '/v1/post/delete/' + postId,
    method: 'delete'
  })
}
