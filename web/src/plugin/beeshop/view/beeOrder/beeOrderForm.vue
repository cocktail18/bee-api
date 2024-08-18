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
        <el-form-item label="商品金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="优惠卡抵扣:" prop="amountCard">
          <el-input-number v-model="formData.amountCard" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="优惠券抵扣:" prop="amountCoupons">
          <el-input-number v-model="formData.amountCoupons" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="运费:" prop="amountLogistics">
          <el-input-number v-model="formData.amountLogistics" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="付款金额:" prop="amountReal">
          <el-input-number v-model="formData.amountReal" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="退款总金额:" prop="amountRefundTotal">
          <el-input-number v-model="formData.amountRefundTotal" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="总税金额:" prop="amountTax">
          <el-input-number v-model="formData.amountTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="消费税:" prop="amountTaxGst">
          <el-input-number v-model="formData.amountTaxGst" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="税服务费用:" prop="amountTaxService">
          <el-input-number v-model="formData.amountTaxService" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="自动发货状态:" prop="autoDeliverStatus">
          <el-input v-model.number="formData.autoDeliverStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关闭订单时间:" prop="dateClose">
          <el-date-picker v-model="formData.dateClose" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="支付订单时间:" prop="datePay">
          <el-date-picker v-model="formData.datePay" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="商品总数量:" prop="goodsNumber">
          <el-input v-model.number="formData.goodsNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否退款:" prop="hasRefund">
          <el-switch v-model="formData.hasRefund" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="核销码:" prop="hxNumber">
          <el-input v-model="formData.hxNumber" :clearable="true"  placeholder="请输入核销码" />
       </el-form-item>
        <el-form-item label="下单ip:" prop="ip">
          <el-input v-model="formData.ip" :clearable="true"  placeholder="请输入下单ip" />
       </el-form-item>
        <el-form-item label="是否可以核销:" prop="isCanHx">
          <el-switch v-model="formData.isCanHx" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="用户删除:" prop="isDelUser">
          <el-switch v-model="formData.isDelUser" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="订单已经结束:" prop="isEnd">
          <el-switch v-model="formData.isEnd" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否有收益:" prop="isHasBenefit">
          <el-switch v-model="formData.isHasBenefit" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="需要配送:" prop="isNeedLogistics">
          <el-switch v-model="formData.isNeedLogistics" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否已经支付:" prop="isPay">
          <el-switch v-model="formData.isPay" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否有积分:" prop="isScoreOrder">
          <el-switch v-model="formData.isScoreOrder" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否拼团成功:" prop="isSuccessPingtuan">
          <el-switch v-model="formData.isSuccessPingtuan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="订单号:" prop="orderNumber">
          <el-input v-model="formData.orderNumber" :clearable="true"  placeholder="请输入订单号" />
       </el-form-item>
        <el-form-item label="订单类型:" prop="orderType">
          <el-input v-model.number="formData.orderType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父订单号:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取单号:" prop="qudanhao">
          <el-input v-model="formData.qudanhao" :clearable="true"  placeholder="请输入取单号" />
       </el-form-item>
        <el-form-item label="退款状态:" prop="refundStatus">
          <el-input v-model.number="formData.refundStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="获得积分:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="扣除积分:" prop="scoreDeduction">
          <el-input v-model.number="formData.scoreDeduction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商店id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自提商店id:" prop="shopIdZt">
          <el-input v-model.number="formData.shopIdZt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自提商店名称:" prop="shopNameZt">
          <el-input v-model="formData.shopNameZt" :clearable="true"  placeholder="请输入自提商店名称" />
       </el-form-item>
        <el-form-item label="订单状态，1未发货:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="小费金额:" prop="trips">
          <el-input v-model.number="formData.trips" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
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
  createBeeOrder,
  updateBeeOrder,
  findBeeOrder
} from '@/plugin/beeshop/api/beeOrder'

defineOptions({
    name: 'BeeOrderForm'
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
            amount: 0,
            amountCard: 0,
            amountCoupons: 0,
            amountLogistics: 0,
            amountReal: 0,
            amountRefundTotal: 0,
            amountTax: 0,
            amountTaxGst: 0,
            amountTaxService: 0,
            autoDeliverStatus: undefined,
            dateClose: new Date(),
            datePay: new Date(),
            goodsNumber: undefined,
            hasRefund: false,
            hxNumber: '',
            ip: '',
            isCanHx: false,
            isDelUser: false,
            isEnd: false,
            isHasBenefit: false,
            isNeedLogistics: false,
            isPay: false,
            isScoreOrder: false,
            isSuccessPingtuan: false,
            orderNumber: '',
            orderType: undefined,
            pid: undefined,
            qudanhao: '',
            refundStatus: undefined,
            remark: '',
            score: undefined,
            scoreDeduction: undefined,
            shopId: undefined,
            shopIdZt: undefined,
            shopNameZt: '',
            status: undefined,
            trips: undefined,
            type: undefined,
            uid: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeOrder({ ID: route.query.id })
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
               res = await createBeeOrder(formData.value)
               break
             case 'update':
               res = await updateBeeOrder(formData.value)
               break
             default:
               res = await createBeeOrder(formData.value)
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
