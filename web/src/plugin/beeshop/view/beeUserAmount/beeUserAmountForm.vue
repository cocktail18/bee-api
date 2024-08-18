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
        <el-form-item label="余额:" prop="balance">
          <el-input-number v-model="formData.balance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="冻结金额:" prop="freeze">
          <el-input-number v-model="formData.freeze" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="佣金:" prop="fxCommisionPaying">
          <el-input-number v-model="formData.fxCommisionPaying" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="成长值:" prop="growth">
          <el-input-number v-model="formData.growth" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="积分:" prop="score">
          <el-input-number v-model="formData.score" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="上一个周期的积分:" prop="scoreLastRound">
          <el-input-number v-model="formData.scoreLastRound" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="支付总额:" prop="totalPayAmount">
          <el-input-number v-model="formData.totalPayAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="支付订单数:" prop="totalPayNumber">
          <el-input-number v-model="formData.totalPayNumber" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="总积分:" prop="totalScore">
          <el-input-number v-model="formData.totalScore" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="撤回次数:" prop="totalWithdraw">
          <el-input-number v-model="formData.totalWithdraw" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="累计消费金额:" prop="totalConsumed">
          <el-input-number v-model="formData.totalConsumed" :precision="2" :clearable="true"></el-input-number>
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
  createBeeUserAmount,
  updateBeeUserAmount,
  findBeeUserAmount
} from '@/plugin/beeshop/api/beeUserAmount'

defineOptions({
    name: 'BeeUserAmountForm'
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
            dateDelete: undefined,
            uid: undefined,
            balance: 0,
            freeze: 0,
            fxCommisionPaying: 0,
            growth: 0,
            score: 0,
            scoreLastRound: 0,
            totalPayAmount: 0,
            totalPayNumber: 0,
            totalScore: 0,
            totalWithdraw: 0,
            totalConsumed: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeUserAmount({ ID: route.query.id })
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
               res = await createBeeUserAmount(formData.value)
               break
             case 'update':
               res = await updateBeeUserAmount(formData.value)
               break
             default:
               res = await createBeeUserAmount(formData.value)
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
