
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商店id:" prop="shopId">
          <el-select v-model="formData.shopId" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in shopList" :key="key" :label="item.id + (item.isDeleted ? '(已删除)' :'') " :value="parseInt(item.id)" />
          </el-select>
        </el-form-item>
        <el-form-item label="state:" prop="state">
          <el-input v-model="formData.state" :clearable="true" placeholder="授权校验码"/>
        </el-form-item>
        <el-form-item label="开发者帐号ID:" prop="state">
          <el-input v-model="formData.source" :clearable="true" placeholder="云喇叭开放平台分配的开发者帐号ID"/>
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
  bindYunlabaShop
} from '@/plugin/beeshop/api/beeDelivery'

defineOptions({
  name: 'BeeDeliveryForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import {getBeeShopInfoList} from "@/plugin/beeshop/api/beeShopInfo";

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  shopId: undefined,
  state: '',
  source: '',
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

const shopList = ref([])
const getShopList = async() => {
  const table = await getBeeShopInfoList({ page: 1, pageSize: 10000 })
  if (table.code === 0) {
    shopList.value = table.data.list
  }
}
getShopList()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.state) {
    formData.value.state = route.query.state
    formData.value.source = route.query.source
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res = await bindYunlabaShop(formData.value)
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
