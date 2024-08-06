
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商店用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="appid:" prop="appid">
          <el-input v-model="formData.appid" :clearable="true"  placeholder="请输入appid" />
        </el-form-item>
        <el-form-item label="app_secret:" prop="appSecret">
          <el-input v-model="formData.appSecret" :clearable="true"  placeholder="请输入app_secret" />
        </el-form-item>
        <el-form-item label="打印机名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入打印机名称" />
        </el-form-item>
        <el-form-item label="打印机密钥:" prop="key">
          <el-input v-model="formData.key" :clearable="true"  placeholder="请输入打印机密钥" />
        </el-form-item>
        <el-form-item label="品牌:" prop="brand">
          <el-input v-model.number="formData.brand" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="打印条件:" prop="condition">
          <el-input v-model.number="formData.condition" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="打印机编号:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入打印机编号" />
        </el-form-item>
        <el-form-item label="打印份数:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="打印模板:" prop="template">
          <el-input v-model="formData.template" :clearable="true"  placeholder="请输入打印模板" />
        </el-form-item>
        <el-form-item label="已删除:" prop="isDeleted">
          <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="创建时间:" prop="dateAdd">
          <el-date-picker v-model="formData.dateAdd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="更新时间:" prop="dateUpdate">
          <el-date-picker v-model="formData.dateUpdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="删除时间:" prop="dateDelete">
          <el-date-picker v-model="formData.dateDelete" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createBeePrinter,
  updateBeePrinter,
  findBeePrinter
} from '@/plugin/beeshop/api/beePrinter'

defineOptions({
  name: 'BeePrinterForm'
})

// 自动获取字典
import {getDictFunc} from '@/utils/format'
import {useRoute, useRouter} from "vue-router"
import {ElMessage} from 'element-plus'
import {ref, reactive} from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  id: undefined,
  userId: undefined,
  appid: '',
  appSecret: '',
  name: '',
  key: '',
  brand: undefined,
  condition: undefined,
  code: '',
  num: undefined,
  shopId: undefined,
  template: '',
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
})
// 验证规则
const rule = reactive({})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findBeePrinter({ID: route.query.id})
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBeePrinter(formData.value)
        break
      case 'update':
        res = await updateBeePrinter(formData.value)
        break
      default:
        res = await createBeePrinter(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
