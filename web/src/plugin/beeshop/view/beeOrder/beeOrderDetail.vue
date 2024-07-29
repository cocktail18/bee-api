<template>
  <div>
    <div class="static-content-item">
      <el-divider direction="horizontal" content-position="left">订单信息</el-divider>
    </div>
    <div class="table-container">
      <table class="table-layout">
        <thead>
        <tr>
          <th class="table-cell" width="280"></th>
          <th class="table-cell" width="280"></th>
          <th class="table-cell" width="280"></th>
          <th class="table-cell" width="280"></th>
        </tr>
        </thead>
        <tbody>

        <tr>
          <td class="table-cell">用户编号 {{ beeOrderData.uid }}</td>
          <td class="table-cell">昵称 {{ beeUserInfo.nick }}</td>
          <td class="table-cell">手机号 </td>
          <td class="table-cell">用户名 </td>
        </tr>
        <tr>
          <td class="table-cell">交易时间 {{ formatDate(beeOrderData.dateAdd) }}</td>
          <td class="table-cell">订单类型 普通订单</td>
          <td class="table-cell">订单id {{ beeOrderData.id }}</td>
          <td class="table-cell">订单号 {{ beeOrderData.orderNumber }}</td>
        </tr>
        <tr>
          <td class="table-cell">商品数量 {{ beeOrderData.goodsNumber }}</td>
          <td class="table-cell">状态 {{ formatEnum(beeOrderData.status, beeOrderStatus) }}</td>
          <td class="table-cell">核销码 {{ beeOrderData.hxNumber }}</td>
          <td class="table-cell">是否支付 <font color="{{beeOrderData.isPay ? 'green' : 'red'}}">{{ beeOrderData.isPay ? '已支付' : '未支付' }}</font></td>
        </tr>
        <tr>
          <td class="table-cell">是否快递 {{ beeOrderData.isNeedLogistics ? '需要' : '不需要' }}</td>
          <td class="table-cell">有无退款 {{ beeOrderData.hasRefund ? '有' : '无' }}</td>
          <td class="table-cell">备注 {{ beeOrderData.remark }}</td>
          <td class="table-cell">更新时间 {{ formatDate(beeOrderData.dateUpdate) }}</td>
        </tr>
        <tr>
          <td class="table-cell">商品金额 {{ beeOrderData.amount }}</td>
          <td class="table-cell">实收金额 {{ beeOrderData.amountReal }}</td>
          <td class="table-cell">优惠金额 {{ beeOrderData.amountCoupons }}</td>
          <td class="table-cell">运费金额 {{ beeOrderData.amountLogistics }}</td>
        </tr>
        <tr>
          <td class="table-cell">自提门店ID {{ beeOrderData.shopIdZt }}</td>
          <td class="table-cell">自提门店 {{ beeOrderData.shopNameZt }}</td>
          <td class="table-cell">支付时间 {{ formatDate(beeOrderData.datePay) }}</td>
          <td class="table-cell">关闭时间 {{ formatDate(beeOrderData.dateClose) }}</td>
        </tr>
        <tr v-if="beeOrderLogisticsList.length > 0">
          <td class="table-cell">联系人 {{beeOrderLogisticsList[0].linkMan}}</td>
          <td class="table-cell">联系电话 {{beeOrderLogisticsList[0].mobile}}</td>
          <td class="table-cell" colspan="2">收获地址 {{beeOrderLogisticsList[0].provinceStr}}{{beeOrderLogisticsList[0].areaStr}}{{beeOrderLogisticsList[0].cityStr}}{{beeOrderLogisticsList[0].address}}</td>
        </tr>
        <tr v-if="beeOrderReputationList.length > 0">
          <td class="table-cell">评价 {{beeOrderReputationList[0].reputation}}</td>
          <td class="table-cell">备注 {{beeOrderReputationList[0].remark}}</td>
          <td class="table-cell" colspan="2">
            <template v-if="beeOrderReputationList[0].picsStr">
              <template v-for="(item, index) in beeOrderReputationList[0].picsStr.split(',')">
                <CustomPic style="margin-top:8px" :pic-src="item"/>
              </template>
            </template>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>

  <div>
    <div class="static-content-item">
      <el-divider direction="horizontal" content-position="left">其他信息</el-divider>
    </div>
    <div class="table-container">
      <el-row
          :gutter="15"
          class="py-1"
      >
        <el-col :span="12">
          <el-form ref="elFormRef" label-position="right" label-width="80px" style="margin-left: 20px">
            <template v-for="(item, index) in beeOrderExtList" :key="index">
              <el-form-item :label="item.key" prop="item.key">
                <el-input v-model="item.value" :clearable="true" placeholder="请输入" />
              </el-form-item>
            </template>
            <el-form-item>
              <el-button type="primary" @click="saveExtData">保存</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </div>

  <div>
    <div class="static-content-item">
      <el-divider direction="horizontal" content-position="left">商品信息</el-divider>
    </div>
    <div class="table-container">
      <table class="table-layout" style="text-align: center">
        <thead>
        <tr>
          <th class="table-cell">商品ID</th>
          <th class="table-cell" width="200">商品图</th>
          <th class="table-cell" width="100">商品名</th>
          <th class="table-cell" width="200">sku</th>
          <th class="table-cell" width="200">数量/规格</th>
          <th class="table-cell" width="200">单价/总价</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in beeOrderGoodsList" :key="index">
          <td class="table-cell">{{ item.goodsId }}</td>
          <td class="table-cell"><CustomPic style="margin-top:8px" :pic-src="item.pic"/> </td>
          <td class="table-cell">{{ item.goodsName }} </td>
          <td class="table-cell">{{ item.property }} </td>
          <td class="table-cell">{{ item.number }}/{{ item.unit }} </td>
          <td class="table-cell">{{ item.amountSingle }}/{{ item.amount }}  </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>

  <div>
    <div class="static-content-item">
      <el-divider direction="horizontal" content-position="left">订单记录</el-divider>
    </div>
    <div class="table-container">
      <table class="table-layout" style="text-align: left">
        <thead>
        <tr>
          <th class="table-cell" width="200">操作时间</th>
          <th class="table-cell" width="200">操作名称</th>
          <th class="table-cell" width="400">备注</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in beeOrderLog" :key="index">
          <td class="table-cell">{{ formatDate(item.dateAdd) }}</td>
          <td class="table-cell">{{ formatEnum(item.type, beeOrderLogType) }} </td>
          <td class="table-cell">{{ item.remark }} </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import {
  createBeeOrder,
  updateBeeOrder,
  findBeeOrder, updateBeeOrderExtJsonStr
} from '@/plugin/beeshop/api/beeOrder'

defineOptions({
    name: 'BeeOrderDetail'
})

// 自动获取字典
import {formatDate, formatEnum, getDictFunc} from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import {getBeeOrderLogList} from "@/plugin/beeshop/api/beeOrderLog";
import {getBeeOrderGoodsList} from "@/plugin/beeshop/api/beeOrderGoods";
import {getBeeOrderLogisticsList} from "@/plugin/beeshop/api/beeOrderLogistics";
import {findBeeUser} from "@/plugin/beeshop/api/beeUser";
import CustomPic from "@/components/customPic/index.vue";

const route = useRoute()
const router = useRouter()

const type = ref('')

const beeOrderLogType = ref([])
const beeOrderStatus = ref([])

const beeOrderData = ref({})
const beeUserInfo = ref({})
const beeOrderLog = ref([])
const beeOrderGoodsList = ref([])
const beeOrderLogisticsList = ref([])
const beeOrderReputationList = ref([])
const beeOrderCouponList = ref([])
const beeOrderExtList = ref([])
// 初始化方法
const init = async () => {
  beeOrderLogType.value = await getDictFunc('BeeOrderLogType')
  beeOrderStatus.value = await getDictFunc('OrderStatus')
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (!route.query.id) {
      return
    }
  const orderId =  route.query.id
  const res = await findBeeOrder({ id: orderId })
  if (res.code === 0) {
    beeOrderData.value = res.data
    const extJsonMap = JSON.parse(res.data.extJsonStr)
    beeOrderExtList.value = Object.keys(extJsonMap).map(key => {
      return {
        key: key,
        value: extJsonMap[key]
      }
    })

    // 获取用户信息
    const userInfoRes = await findBeeUser({ id: res.data.uid})
    if (userInfoRes.code === 0) {
      beeUserInfo.value = userInfoRes.data
    }
  }

  const orderLogRes = await getBeeOrderLogList({ page: 1, pageSize: 10000, orderId: orderId })
  if (orderLogRes.code === 0) {
    beeOrderLog.value = orderLogRes.data.list
  }

  const orderGoodsRes = await getBeeOrderGoodsList({ page: 1, pageSize: 10000, orderId: orderId })
  if (orderGoodsRes.code === 0) {
    beeOrderGoodsList.value = orderGoodsRes.data.list
  }

  const orderLogisticsRes = await getBeeOrderLogisticsList({ page: 1, pageSize: 10000, orderId: orderId })
  if (orderLogisticsRes.code === 0) {
    beeOrderLogisticsList.value = orderLogisticsRes.data.list
  }
}

init()

const saveExtData = async () => {
  const extJsonStr = JSON.stringify(beeOrderExtList.value.reduce((acc, cur) => {
    acc[cur.key] = cur.value
    return acc
  }, {}))
  await updateBeeOrderExtJsonStr({ id: beeOrderData.value.id, extJsonStr })
}

</script>

<style>
.py-1{
  width: 100%;
}
.table-layout{
  margin-left: 30px;
}
</style>
