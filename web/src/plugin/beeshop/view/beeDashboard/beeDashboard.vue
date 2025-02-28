<template>
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 py-2 gap-4 md:gap-2 gva-container2">
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/shop-user-admin/beeUser')">
      <gva-chart :data="userData" title="用户数" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/shop-order-admin/beeOrder')">
      <gva-chart :data="orderData" title="订单统计" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/bee_goods_admin/beeShopGoods')">
      <gva-chart :data="goodsData" title="商品数量" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/beeOrderTodo')">
      <gva-chart :data="orderTodoData" title="订单待办" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/beeFinancialManager/beePayLog')">
      <gva-chart :data="payData" title="总支付金额" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32" @click="gotoPage('/bee_index/beeFinancialManager/beePayLog')">
      <gva-chart :data="payNumData" title="支付人数" />
    </gva-card>

    <gva-card title="待办事项" custom-class="col-span-1 md:col-span-6 lg:col-span-6 row-span-2">
      <div>
        <el-table :data="orderTodo.list" stripe style="width: 100%" @row-click="gotoTodoPage">
          <el-table-column prop="orderNumber" label="订单号"  width="200" />
          <el-table-column prop="id" label="事项" align="center" show-overflow-tooltip >
            <template #default="scope">
              {{formatEnum(scope.row.status, beeOrderStatus)}}
            </template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="100" />
          <el-table-column prop="datePay" label="支付时间" width="100"  >
            <template #default="scope">{{ formatDate(scope.row.datePay) }}</template>
          </el-table-column>
        </el-table>
      </div>
    </gva-card>
  </div>
</template>

<script setup>
import {reactive, ref} from 'vue'
import { GvaChart, GvaCard } from "./componenst"
import { dayjs } from 'element-plus'
import {getBeeUserList} from "@/plugin/beeshop/api/beeUser";
import {getBeePayLogList, getBeePayTotal} from "@/plugin/beeshop/api/beePayLog";
import {getBeeOrderList} from "@/plugin/beeshop/api/beeOrder";
import {getBeeShopGoodsList} from "@/plugin/beeshop/api/beeShopGoods";
import {formatDate, formatEnum, getDictFunc} from "@/utils/format";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const router = useRouter()

const beeOrderStatus = ref([])
const userData = ref([])
const goodsData = ref([])
const orderData = ref([])
const payData = ref([])
const payNumData = ref([])
const orderTodoData = ref([])
const orderTodo = ref({})

const init = async () => {
  beeOrderStatus.value = await getDictFunc('OrderStatus')
  for (let i = 7; i >= 0; i--) {
    // 获取最近7天数据
    const start = new Date(0)
    const end = dayjs().add(-1 * i, 'day').endOf('day').toDate()

    const funcs = [];
    funcs.push(getBeeUserList({page: 1, pageSize: 1, startDateAdd: start, endDateAdd: end}))
    funcs.push(getBeePayLogList({page: 1, pageSize: 1, startDateAdd: start, endDateAdd: end, distinct: 'uid'}))
    funcs.push(getBeeOrderList({page: 1, pageSize: 1, startDateAdd: start, endDateAdd: end}))
    funcs.push(getBeeShopGoodsList({page: 1, pageSize: 1, startDateAdd: start, endDateAdd: end}))
    funcs.push(getBeePayTotal({page: 1, pageSize: 1, startDateAdd: start, endDateAdd: end, sum: 'money'}))
    const results = await Promise.all(funcs)
    console.log(results)

    const userTable = results[0]
    if (userTable.code === 0) {
      userData.value.push(userTable.data.total)
    }
    const payLogUidTable = results[1]
    if (payLogUidTable.code === 0) {
      payNumData.value.push(payLogUidTable.data.total)
    }
    const orderTable = await results[2]
    if (orderTable.code === 0) {
      orderData.value.push(orderTable.data.total)
    }
    const goodsTable = await results[3]
    if (goodsTable.code === 0) {
      goodsData.value.push(goodsTable.data.total)
    }
    const payDataTable = await results[4]
    if (payDataTable.code === 0) {
      payData.value.push(payDataTable.data.total)
    }
  }

  const orderTable = await getBeeOrderList({page: 1, pageSize: 100, status: 1})
  if (orderTable.code === 0) {
    orderTodoData.value.push(orderTable.data.total)
    orderTodoData.value.push(orderTable.data.total)
    orderTodo.value = orderTable.data
  }
}

init()

const gotoPage = async(page) => {
  page = '/layout' + page
  await router.push(page)
}

const gotoTodoPage = (row) => {
  gotoPage('/bee_index/shop-order-admin/beeOrderDetail?id='+row.id)
}

</script>

<style lang="scss" scoped>
</style>
