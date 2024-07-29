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
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="可以使用时间:" prop="dateStart">
          <el-date-picker v-model="formData.dateStart" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="过期时间（毫秒）:" prop="expiryMillis">
          <el-input v-model.number="formData.expiryMillis" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="金额:" prop="money">
          <el-input-number v-model="formData.money" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="满xx可用:" prop="moneyHreshold">
          <el-input v-model.number="formData.moneyHreshold" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券面额范围:" prop="moneyMin">
          <el-input v-model.number="formData.moneyMin" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券面额范围:" prop="moneyMax">
          <el-input v-model.number="formData.moneyMax" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券类型:" prop="moneyType">
          <el-input v-model.number="formData.moneyType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入优惠券名字" />
       </el-form-item>
        <el-form-item label="仅抵扣运费:" prop="onlyFreight">
          <el-switch v-model="formData.onlyFreight" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="优惠券id:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理员批量发放:" prop="source">
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入管理员批量发放" />
       </el-form-item>
        <el-form-item label="使用状态，0未使用，1使用中，2已使用:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
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
  createBeeUserCoupon,
  updateBeeUserCoupon,
  findBeeUserCoupon
} from '@/plugin/beeshop/api/beeUserCoupon'

defineOptions({
    name: 'BeeUserCouponForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            userId: undefined,
            isDeleted: false,
            dateAdd: new Date(),
            dateUpdate: new Date(),
            dateDelete: new Date(),
            uid: undefined,
            dateStart: new Date(),
            expiryMillis: undefined,
            money: 0,
            moneyHreshold: undefined,
            moneyMin: undefined,
            moneyMax: undefined,
            moneyType: undefined,
            name: '',
            onlyFreight: false,
            pid: undefined,
            shopId: undefined,
            source: '',
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeUserCoupon({ ID: route.query.id })
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
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBeeUserCoupon(formData.value)
               break
             case 'update':
               res = await updateBeeUserCoupon(formData.value)
               break
             default:
               res = await createBeeUserCoupon(formData.value)
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
