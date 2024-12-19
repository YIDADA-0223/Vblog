<template>
  <a-layout>
    <a-layout-header class="header">
      <div class="header-left">
        <img width="40x" src="@/views/backend/access/k8s2024.png" />
        <p>K8S多集群管理平台</p>
      </div>
      <div class="header-right">
        <a-space>
          <a-button type="text" @click="router.push({ name: 'FrontendBlogList' })">前台</a-button>
          <!-- 1.调用logout接口 -->
          <!-- 2.跳转到登录界面 -->
          <a-button :loading="logoutLoadding" type="text" @click="handleLogout"
            ><span style="margin-right: 12px">退出</span><icon-export
          /></a-button>
        </a-space>
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider collapsible :width="260" class="sider-bar" breakpoint="xl">
        <a-menu
          @menu-item-click="handleMenuItemClick"
          :style="{ width: '100%', height: '100%' }"
          :default-open-keys="['BackendManagement']"
          :default-selected-keys="['BackendBlogList']"
          breakpoint="xl"
          @collapse="onCollapse"
        >
          <a-sub-menu key="BackendManagement">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>文章管理</template>
            <a-menu-item key="BackendBlogList">文章列表</a-menu-item>
            <a-menu-item key="BackendTagList">标签列表</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="BackendCommentsList">
            <template #icon><icon-book /></template>
            <template #title>评论管理</template>
            <a-menu-item key="BackendCommentsList">评论列表</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </a-layout-sider>
      <a-layout-content class="page"><router-view /></a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { LOGOUT } from '@/api/vblogs'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import app from '@/stores/app'
// 路由
const router = useRouter()
// 退出状态
const logoutLoadding = ref(false)
// 处理退出登录方法
const handleLogout = async () => {
  // 1. 调用logout接口
  try {
    // 设置退出状态
    console.log(app.value.token.refresh_token)
    logoutLoadding.value = true
    const accessToken = app.value.token.access_token
    const refreshToken = app.value.token.refresh_token
    await LOGOUT(accessToken, refreshToken)
    // console.log(resp)
    // 用户退出销毁数据
    // app.value.token = undefined
  } catch (err) {
    console.log(err)
  } finally {
    logoutLoadding.value = false
  }
  router.push({ name: 'LoginPage' })
}
// 侧边栏
const onCollapse = (v) => {
  console.log(v)
}
// 导航条联动绑定
const handleMenuItemClick = (v) => {
  // 联动直接路由过去
  router.push({ name: v })
  // console.log(v)
}
</script>

<style lang="css" scoped>
.header {
  height: 59px;
  border-bottom: 1px solid var(--color-border);
  padding: 0px 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: var(--color-neutral-8);
}
.sider-bar {
  height: calc(100vh - 60px);
}
.header-left {
  font-size: 25px;
  margin: 10px;
  display: flex;
  align-items: center;
  /* margin-right: 10px; */
}
.header-right {
}
/* 文章内布局 */
.page {
  padding: 12px;
  height: calc(100vh - 60px);
  overflow: auto;
  /* 背景颜色 */
  /* background-color: var(--color-neutral-2); */
  /* 滚动条消失但保留功能 */
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.page::-webkit-scrollbar {
  display: none;
}
</style>
