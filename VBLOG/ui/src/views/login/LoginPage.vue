<template>
  <div class="login-container-wrapper">
    <div class="login-container">
      <div class="title">后台管理系统</div>
      <!-- 登录表单 -->
      <div style="margin: 20px">
        <a-form :model="form" size="large" @submit="handleSubmit" auto-label-width>
          <a-form-item
            required
            hide-asterisk
            field="username"
            label=""
            :rules="{ required: true, message: '请输入用户名' }"
          >
            <a-input v-model="form.username" placeholder="请输入你的用户名" allow-clear>
              <template #prefix>
                <icon-user />
              </template>
            </a-input>
          </a-form-item>
          <a-form-item
            required
            hide-asterisk
            field="password"
            label=""
            :rules="{ requied: true, message: '请输入密码' }"
          >
            <a-input-password v-model="form.password" placeholder="请输入你的密码" allow-clear>
              <template #prefix>
                <icon-lock />
              </template>
            </a-input-password>
          </a-form-item>
          <a-form-item field="is_member" label="">
            <a-checkbox value="true">记住密码</a-checkbox>
          </a-form-item>
          <a-form-item>
            <a-button :loading="LoginLoadding" html-type="submit" style="width: 100%"
              >登录</a-button
            >
          </a-form-item>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { LOGIN } from '@/api/vblogs'
import { useRouter } from 'vue-router'
import app from '@/stores/app'
// 获取router，使用Router来帮忙进行路由切换
const router = useRouter()
const LoginLoadding = ref(false)
// 响应式函数
const form = reactive({
  username: '',
  password: '',
  is_remember: false
})
// 表单事件
const handleSubmit = async (data) => {
  if (!data.errors) {
    console.log(data.values)
    // 把这个form表单数据提交给后代API Server
    // 1.登录login（是不是网络请求中）
    // 2.如果登录成功了，需要跳转到后台页面
    try {
      LoginLoadding.value = true
      const resp = await LOGIN(data.values)
      // 如果登录成功了需要跳转到后台页面
      // 使用push方法，指定需要跳转的路由
      // 直接把登录后到状态信息保存在LocalStores，并且这个LocalStores还是响应式的，通过Vueuse实现
      app.value.token = resp
      // console.log(resp)
      router.push({ name: 'BackendLayout' })
    } finally {
      LoginLoadding.value = false
    }
  }
}
</script>

<style lang="css" scoped>
.login-container-wrapper {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  padding-top: 15%;
  background: white;
}
.login-container {
  height: 350px;
  width: 400px;
  border: solid 1px var(--color-neutral-3);
  padding: 16px;
  box-shadow: 0 -2px 5px rgba(0, 0, 0, 0.1);
}
.title {
  font-size: 24px;
  font-weight: 600;
  display: flex;
  justify-content: center;
  color: var(--color-neutral-8);
  margin: 12px;
}
</style>
