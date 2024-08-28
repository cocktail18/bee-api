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
          <td class="table-cell">订单id {{ beeOrderData.id }}</td>
          <td class="table-cell">订单号 {{ beeOrderData.orderNumber }}</td>
          <td class="table-cell">用户编号 {{ beeOrderData.uid }}</td>
          <td class="table-cell">昵称 {{ beeUserInfo.nick }}</td>
        </tr>
        <tr>
          <td class="table-cell">交易时间 {{ formatDate(beeOrderData.dateAdd) }}</td>
          <td class="table-cell">打包费 {{ beeOrderData.trips }}</td>
          <td class="table-cell">商品金额 {{ beeOrderData.amount }}</td>
          <td class="table-cell">实收金额 {{ beeOrderData.amountReal }}</td>
        </tr>
        <tr>
          <td class="table-cell">优惠金额 {{ beeOrderData.amountCoupons }}</td>
          <td class="table-cell">运费金额 {{ beeOrderData.amountLogistics }}</td>
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
    <div class="table-container" style="padding-left: 60px">
      <el-row
          :gutter="15"
          class="py-1"
      >
        <el-col :span="4" v-if="!beeOrderData.isPay"><el-button type="warning" @click="payOrderOffline(beeOrderData)">线下支付</el-button></el-col>
        <el-col :span="4" v-if="beeOrderData.status === 1"><el-button type="warning" @click="shipBeeOrder(beeOrderData)">已发货</el-button></el-col>
        <el-col :span="4" v-if="beeOrderData.status === 2"><el-button type="warning" @click="confirmShippedBeeOrder(beeOrderData)">已确认收货</el-button></el-col>
      </el-row>
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

  <div v-if="beeOrderData.isNeedLogistics">
    <div class="static-content-item">
      <el-divider direction="horizontal" content-position="left">配送信息</el-divider>
    </div>
    <div class="table-container">
      <table class="table-layout" style="text-align: center">
        <thead>
        <tr>
          <th class="table-cell">id</th>
          <th class="table-cell" width="200">第三方配送订单号</th>
          <th class="table-cell" width="200">配送订单号</th>
          <th class="table-cell" width="100">状态</th>
          <th class="table-cell" width="100">配送费</th>
          <th class="table-cell" width="100">实际支付金额</th>
          <th class="table-cell" width="100">取消配送扣费</th>
          <th class="table-cell" width="100">是否已取消</th>
          <th class="table-cell" width="100">更新时间</th>
          <th class="table-cell" width="100">操作</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in beeOrderPeisongList" :key="index">
          <td class="table-cell">{{ item.id }}</td>
          <td class="table-cell">{{ item.peisongOrderId }} </td>
          <td class="table-cell">{{ item.peisongOrderNo }} </td>
          <td class="table-cell">{{ formatEnum(item.status, beeOrderPeisongStatus) }} </td>
          <td class="table-cell">{{ item.money }} </td>
          <td class="table-cell">{{ item.realMoney }} </td>
          <td class="table-cell">{{ item.deductFee }} </td>
          <td class="table-cell">{{ item.isCancel ? '是' : '否' }} </td>
          <td class="table-cell">{{ formatDate(item.dateUpdate) }} </td>
          <td class="table-cell">
            <div><el-button type="text" @click="showDetailModal(item.id)">查看详情</el-button></div>
            <div><el-button v-if="!item.isCancel" type="text" @click="showCancelModal(item.id)">取消配送</el-button></div>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
    <div style="margin-left: 30px">
      <el-button type="primary" @click="notifyDelivery(beeOrderData.id)">通知配送商取货</el-button>
    </div>
  </div>
    <el-dialog
        v-model="cancelModalVisible"
        title="取消配送"
        width="500"
        :before-close="handleCloseCancelModal"
    >
      <div>
        <el-form-item label="取消原因" prop="orderNumber">
          <el-select v-model="cancelFormData.reasonId" clearable placeholder="请选择取消原因" :clearable="false">
            <el-option v-for="(item,key) in beeDeliveryCancelReason" :key="key" :label="item.label"
                       :value="parseInt(item.value)"/>
          </el-select>
        </el-form-item>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cancelModalVisible = false">取消</el-button>
          <el-button type="primary" @click="cancelDelivery">
            确定
          </el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
        v-model="detailModalVisible"
        title="配送详情"
        width="500"
        :before-close="handleClosePeisongDetailModal"
    >
      <el-row
          :gutter="15"
          class="py-1"
      >
        <el-col :span="12" class="detail-col">配送ID：{{peisongOrderDetail.orderId}}</el-col>
        <el-col :span="12" class="detail-col">订单状态：{{peisongOrderDetail.statusMsg}}</el-col>
        <el-col :span="12" class="detail-col">骑手姓名：{{peisongOrderDetail.transporterName}}</el-col>
        <el-col :span="12" class="detail-col">骑手电话：{{peisongOrderDetail.transporterPhone}}</el-col>
        <el-col :span="12" class="detail-col">骑手经度：{{peisongOrderDetail.transporterLng}}</el-col>
        <el-col :span="12" class="detail-col">骑手纬度：{{peisongOrderDetail.transporterLat}}</el-col>
        <el-col :span="12" class="detail-col">订单总费用：{{peisongOrderDetail.deliveryFee}}</el-col>
        <el-col :span="12" class="detail-col">小费：{{peisongOrderDetail.tips}}元</el-col>
        <el-col :span="12" class="detail-col">优惠券费用：{{peisongOrderDetail.couponFee}}元</el-col>
        <el-col :span="12" class="detail-col">保价费：{{peisongOrderDetail.insuranceFee}}元</el-col>
        <el-col :span="12" class="detail-col">订单实际费用消耗：{{peisongOrderDetail.actualFee}}</el-col>
        <el-col :span="12" class="detail-col">配送距离,单位为米：{{peisongOrderDetail.distance}}</el-col>
        <el-col :span="12" class="detail-col">发单时间：{{peisongOrderDetail.createTime}}</el-col>
        <el-col :span="12" class="detail-col">接单时间：{{peisongOrderDetail.acceptTime}}</el-col>
        <el-col :span="12" class="detail-col">取货时间：{{peisongOrderDetail.fetchTime}}</el-col>
        <el-col :span="12" class="detail-col">送达时间：{{peisongOrderDetail.finishTime}}</el-col>
        <el-col :span="12" class="detail-col">取消时间：{{peisongOrderDetail.cancelTime}}</el-col>
        <el-col :span="12" class="detail-col">收货码：{{peisongOrderDetail.orderFinishCode}}</el-col>
        <el-col :span="12" class="detail-col">违约金：{{peisongOrderDetail.deductFee}}</el-col>
      </el-row>
      <div class="dialog-footer">
        <el-button type="primary" @click="detailModalVisible = false">
          确定
        </el-button>
      </div>
    </el-dialog>

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
          <th class="table-cell" width="800">备注</th>
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
  findBeeOrder, updateBeeOrderExtJsonStr, markBeeOrderPaid, updateBeeOrderStatus, shippedBeeOrder
} from '@/plugin/beeshop/api/beeOrder'

defineOptions({
    name: 'BeeOrderDetail'
})

// 自动获取字典
import {formatDate, formatEnum, getDictFunc} from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import {ElMessage, ElMessageBox} from 'element-plus'
import { ref, reactive } from 'vue'
import {getBeeOrderLogList} from "@/plugin/beeshop/api/beeOrderLog";
import {getBeeOrderGoodsList} from "@/plugin/beeshop/api/beeOrderGoods";
import {getBeeOrderLogisticsList} from "@/plugin/beeshop/api/beeOrderLogistics";
import {findBeeUser} from "@/plugin/beeshop/api/beeUser";
import CustomPic from "@/components/customPic/index.vue";
import {
  cancelBeeOrderPeisong, getBeeOrderPeisongDetail,
  getBeeOrderPeisongList,
  notifyBeeOrderPeisong
} from "@/plugin/beeshop/api/beeOrderPeisong";

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
const beeOrderPeisongList = ref([])
const beeDeliveryCancelReason = ref([])
const beeOrderPeisongStatus = ref([])
const cancelModalVisible = ref(false)
const detailModalVisible = ref(false)
const peisongOrderDetail = ref({})
// 初始化方法
const init = async () => {
  beeOrderLogType.value = await getDictFunc('BeeOrderLogType')
  beeOrderStatus.value = await getDictFunc('OrderStatus')
  beeOrderPeisongStatus.value = await getDictFunc('BeeOrderPeisongStatus')
  beeDeliveryCancelReason.value = await getDictFunc('BeeDeliveryCancelReason')
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

  const beeOrderPeisongListRes = await getBeeOrderPeisongList({ page: 1, pageSize: 10000, orderId: orderId })
  if (beeOrderPeisongListRes.code === 0) {
    beeOrderPeisongList.value = beeOrderPeisongListRes.data.list
  }
}

init()

const saveExtData = async () => {
  const extJsonStr = JSON.stringify(beeOrderExtList.value.reduce((acc, cur) => {
    acc[cur.key] = cur.value
    return acc
  }, {}))
  const res = await updateBeeOrderExtJsonStr({ id: beeOrderData.value.id, extJsonStr })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })
  }
}

const payOrderOffline = async (row) => {
  ElMessageBox.confirm('确定已收到支付款?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = []
    ids.push(row.id)
    const res = await markBeeOrderPaid({'ids': ids})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      await init()
    }
  })
}

const notifyDelivery = async (orderId) => {
  ElMessageBox.confirm('确定通知取货?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await notifyBeeOrderPeisong({'orderId': orderId})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '通知成功'
      })
      await init()
    }
  })
}

const cancelDelivery = async () => {
  ElMessageBox.confirm('确定取消配送?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await cancelBeeOrderPeisong(cancelFormData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '取消成功'
      })
      cancelModalVisible.value = false
      await init()
    }
  })
}

const shipBeeOrder = async (row) => {
  let tipMsg = '确定已发货?'
  if (beeOrderPeisongList.value.length === 0 && row.peisongType === "kd") {
    tipMsg = '没找到配送单，确定已发货?'
  }
  ElMessageBox.confirm(tipMsg, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await shippedBeeOrder({'id': row.id})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      await init()
    }
  })
}

const confirmShippedBeeOrder = async (row) => {
  ElMessageBox.confirm('确定已收货?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = []
    ids.push(row.id)
    const res = await updateBeeOrderStatus({'id': row.id, 'status': 3})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      await init()
    }
  })
}

const cancelFormData = ref({
  id: 0,
  reasonId: 1,
  reason: '',
})

const showCancelModal = (peisongId) => {
  cancelFormData.value = {
    id: peisongId,
    reasonId: 1,
    reason: '',
  }
  cancelModalVisible.value = true
}

const handleCloseCancelModal = () => {
  cancelFormData.value = {
    id: 0,
    reasonId: 1,
    reason: '',
  }
  cancelModalVisible.value = false
}

const showDetailModal = async (peisongId) => {
  const res = await getBeeOrderPeisongDetail({id: peisongId})
  if (res.code === 0) {
    peisongOrderDetail.value = res.data
    detailModalVisible.value = true
  }
}

const handleClosePeisongDetailModal = () => {
  peisongOrderDetail.value = {}
  detailModalVisible.value = false
}
</script>

<style>
.py-1{
  width: 100%;
}
.detail-col{
  padding-bottom: 7px;
}
.table-layout{
  margin-left: 30px;
}
</style>
