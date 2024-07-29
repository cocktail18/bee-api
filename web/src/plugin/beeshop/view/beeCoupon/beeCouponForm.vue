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
        <el-form-item label="批量:" prop="batchSendUid">
          <el-input v-model.number="formData.batchSendUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第n天失效:" prop="dateEndDays">
          <el-input v-model.number="formData.dateEndDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="失效日期:" prop="dateEnd">
          <el-date-picker v-model="formData.dateEnd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="失效类型，1n天后失效，0指定:" prop="dateEndType">
          <el-input v-model.number="formData.dateEndType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生效日期:" prop="dateStart">
          <el-date-picker v-model="formData.dateStart" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="第n天生效:" prop="dateStartDays">
          <el-input v-model.number="formData.dateStartDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生效时间，1立即，2次日，0指定时间:" prop="dateStartType">
          <el-input v-model.number="formData.dateStartType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="满xx可用:" prop="moneyHreshold">
          <el-input-number v-model="formData.moneyHreshold" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="优惠券最大金额:" prop="moneyMax">
          <el-input-number v-model="formData.moneyMax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="优惠券最小金额:" prop="moneyMin">
          <el-input-number v-model="formData.moneyMin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="类型，1满xx，2折扣:" prop="moneyType">
          <el-input v-model.number="formData.moneyType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入优惠券名字" />
       </el-form-item>
        <el-form-item label="需要支付:" prop="needAmount">
          <el-input-number v-model="formData.needAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="需要积分:" prop="needScore">
          <el-input-number v-model="formData.needScore" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="连续签到多少天可得:" prop="needSignedContinuous">
          <el-input v-model.number="formData.needSignedContinuous" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="剩余数量:" prop="numberLeft">
          <el-input v-model.number="formData.numberLeft" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="每人最多领取多少:" prop="numberPersonMax">
          <el-input v-model.number="formData.numberPersonMax" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总数量:" prop="numberTotle">
          <el-input v-model.number="formData.numberTotle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用数量:" prop="numberUsed">
          <el-input v-model.number="formData.numberUsed" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="仅抵扣运费:" prop="onlyFreight">
          <el-switch v-model="formData.onlyFreight" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="生日赠送:" prop="sendBirthday">
          <el-switch v-model="formData.sendBirthday" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="邀请新人注册赠送:" prop="sendInviteM">
          <el-switch v-model="formData.sendInviteM" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="被邀请赠送:" prop="sendInviteS">
          <el-switch v-model="formData.sendInviteS" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="注册赠送:" prop="sendRegister">
          <el-switch v-model="formData.sendRegister" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="店铺id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="展示:" prop="showInFront">
          <el-switch v-model="formData.showInFront" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="状态，0是正常:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券类型描述:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入优惠券类型描述" />
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
  createBeeCoupon,
  updateBeeCoupon,
  findBeeCoupon
} from '@/plugin/beeshop/api/beeCoupon'

defineOptions({
    name: 'BeeCouponForm'
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
            batchSendUid: undefined,
            dateEndDays: undefined,
            dateEnd: new Date(),
            dateEndType: undefined,
            dateStart: new Date(),
            dateStartDays: undefined,
            dateStartType: undefined,
            moneyHreshold: 0,
            moneyMax: 0,
            moneyMin: 0,
            moneyType: undefined,
            name: '',
            needAmount: 0,
            needScore: 0,
            needSignedContinuous: undefined,
            numberLeft: undefined,
            numberPersonMax: undefined,
            numberTotle: undefined,
            numberUsed: undefined,
            onlyFreight: false,
            sendBirthday: false,
            sendInviteM: false,
            sendInviteS: false,
            sendRegister: false,
            shopId: undefined,
            showInFront: false,
            status: undefined,
            type: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeCoupon({ ID: route.query.id })
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
               res = await createBeeCoupon(formData.value)
               break
             case 'update':
               res = await updateBeeCoupon(formData.value)
               break
             default:
               res = await createBeeCoupon(formData.value)
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
