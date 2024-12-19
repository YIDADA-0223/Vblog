<template>
  <div>
    <a-page-header :title="isEdit ? '编辑文字' : '创建文章'" @back="$router.go(-1)">
      <template #extra>
        <a-button :loading="createLoadding" type="outline" @click="handleSave">
          <template #icon>
            <icon-save />
          </template>
          保存</a-button
        >
      </template>
    </a-page-header>

    <!-- 提交form表单 -->
    <!-- 注意这里的ref不是做绑定而 -->
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-form-item
        field="title"
        label="请输入文章标题"
        :rules="{ required: true, message: '请输入文章标题' }"
      >
        <a-input v-model="form.title" placeholder="please enter your username..." />
      </a-form-item>
      <a-form-item
        field="summary"
        label="请输入文章概要"
        :rules="{ required: true, message: '请输入文章概要' }"
      >
        <a-input v-model="form.summary" placeholder="please enter your post..." />
      </a-form-item>
      <a-form-item
        field="content"
        label="文章内容"
        :rules="{ required: true, message: '请输入文章内容' }"
      >
        <MdEditor v-model="form.content"></MdEditor>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { CREATE_BLOG, UPDATE_BLOG } from '@/api/vblogs'
import { GET_BLOG } from '@/api/vblogs'
// 消息组件
import { Notification } from '@arco-design/web-vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const isEdit = ref(false)

// 判断功能模式
// var blogId = ''
// var blogId = ''
const getBlogLoadding = ref(false)
watch(
  () => router.currentRoute.value.query,
  async (v) => {
    if (v.id) {
      // blogId = v.id
      isEdit.value = true
      // 通过id获取文章内容
      try {
        getBlogLoadding.value = true
        const resp = await GET_BLOG(v.id)
        form.value = resp
        // console.log(resp)
      } finally {
        getBlogLoadding.value = false
      }
    }
    console.log(v)
  },
  { immediate: true }
)
const form = ref({
  title: '',
  author: '',
  content: '',
  summary: '',
  create_by: '',
  tags: {}
})
// form表单实例
// 定义响应式变量与表单进行映射,form表单实例
const formRef = ref(null)
const createLoadding = ref(false)
// 保存
const handleSave = async () => {
  // 校验通过提交value

  const resp = await formRef.value.validate()
  if (!resp) {
    try {
      createLoadding.value = true
      if (isEdit.value) {
        await UPDATE_BLOG(router.currentRoute.value.query.id, form.value)
        Notification.success('更新成功')
      } else {
        await CREATE_BLOG(form.value)
        Notification.success('保存成功')
      }
    } finally {
      createLoadding.value = false
    }
  }
}
</script>

<style lang="css" scoped>
.md-editor {
  height: 1000px;
}
</style>
