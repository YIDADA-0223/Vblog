import axios from 'axios'
import { Message } from '@arco-design/web-vue'

var client = axios.create({
  // 因为开启了代理，直接使用前端页面URL地址，由vite进行代理到后端
  baseURL: '',
  timeout: 5000
})
// 添加一个响应拦截器
client.interceptors.response.use(
  // 请求成功
  (value) => {
    // 返回成功请求后的数据
    return value.data
  },
  //请求失败
  (err) => {
    console.log(err)
    var msg = err.message
    var code = 0
    if (err.response && err.response.data) {
      msg = err.response.data.message
      code = err.response.data.code
    }
    // 业务异常特殊处理
    switch (code) {
      case 5000:
        // 这不是一个vue组件，也不是一个vue hook，当前就是一个和vue没有关系的js库
        location.assign('/login')
        break
      case 5002:
        location.assign('/login')
        break
      case 5003:
        location.assign('/login')
        break
    }
    // 用户的Token失效，或者被撤销

    // 异常采用消息提醒
    Message.error(msg)
    return Promise.reject(err)
  }
)

export default client
