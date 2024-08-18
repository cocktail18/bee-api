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
        <el-form-item label="添加时间:" prop="dateAdd">
          <el-date-picker v-model="formData.dateAdd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="dateUpdate">
          <el-date-picker v-model="formData.dateUpdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="dateDelete">
          <el-date-picker v-model="formData.dateDelete" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="售后类型:" prop="afterSale">
          <el-input v-model="formData.afterSale" :clearable="true"  placeholder="请输入售后类型" />
       </el-form-item>
        <el-form-item label="总金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="优惠券金额:" prop="amountCoupon">
          <el-input-number v-model="formData.amountCoupon" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="单价:" prop="amountSingle">
          <el-input-number v-model="formData.amountSingle" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="基本单价:" prop="amountSingleBase">
          <el-input-number v-model="formData.amountSingleBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="奖励发放结束:" prop="buyRewardEnd">
          <el-switch v-model="formData.buyRewardEnd" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="分类:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="餐饮桌子状态:" prop="cyTableStatus">
          <el-input v-model.number="formData.cyTableStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="返现类型:" prop="fxType">
          <el-input v-model.number="formData.fxType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品id:" prop="goodsId">
          <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品名:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名" />
       </el-form-item>
        <el-form-item label="是否积分订单:" prop="isScoreOrder">
          <el-switch v-model="formData.isScoreOrder" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="数量:" prop="number">
          <el-input v-model.number="formData.number" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="未发货数量:" prop="numberNoFahuo">
          <el-input v-model.number="formData.numberNoFahuo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单id:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="人数:" prop="persion">
          <el-input v-model.number="formData.persion" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品图:" prop="pic">
          <el-input v-model="formData.pic" :clearable="true"  placeholder="请输入商品图" />
       </el-form-item>
        <el-form-item label="价格id:" prop="priceId">
          <el-input v-model.number="formData.priceId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="属性:" prop="property">
          <el-input v-model="formData.property" :clearable="true"  placeholder="请输入属性" />
       </el-form-item>
        <el-form-item label="属性id:" prop="propertyChildIds">
          <el-input v-model="formData.propertyChildIds" :clearable="true"  placeholder="请输入属性id" />
       </el-form-item>
        <el-form-item label="是否已支付:" prop="purchase">
          <el-input v-model="formData.purchase" :clearable="true"  placeholder="请输入是否已支付" />
       </el-form-item>
        <el-form-item label="退款状态:" prop="refundStatus">
          <el-input v-model.number="formData.refundStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商店id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位:" prop="unit">
          <el-input v-model="formData.unit" :clearable="true"  placeholder="请输入单位" />
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
  createBeeOrderGoods,
  updateBeeOrderGoods,
  findBeeOrderGoods
} from '@/plugin/beeshop/api/beeOrderGoods'

defineOptions({
    name: 'BeeOrderGoodsForm'
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
            afterSale: '',
            amount: 0,
            amountCoupon: 0,
            amountSingle: 0,
            amountSingleBase: 0,
            buyRewardEnd: false,
            categoryId: undefined,
            cyTableStatus: undefined,
            fxType: undefined,
            goodsId: undefined,
            goodsName: '',
            isScoreOrder: false,
            number: undefined,
            numberNoFahuo: undefined,
            orderId: undefined,
            persion: undefined,
            pic: '',
            priceId: undefined,
            property: '',
            propertyChildIds: '',
            purchase: '',
            refundStatus: undefined,
            score: undefined,
            shopId: undefined,
            status: undefined,
            type: undefined,
            uid: undefined,
            unit: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeOrderGoods({ ID: route.query.id })
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
               res = await createBeeOrderGoods(formData.value)
               break
             case 'update':
               res = await updateBeeOrderGoods(formData.value)
               break
             default:
               res = await createBeeOrderGoods(formData.value)
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
