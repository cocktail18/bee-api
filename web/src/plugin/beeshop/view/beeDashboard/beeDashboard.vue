<template>
  <div
    class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 py-2 gap-4 md:gap-2 gva-container2"
  >
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/shop-user-admin/beeUser')"
    >
      <gva-chart :data="userData" title="用户数" />
    </gva-card>
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/shop-order-admin/beeOrder')"
    >
      <gva-chart :data="orderData" title="订单统计" />
    </gva-card>
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/bee_goods_admin/beeShopGoods')"
    >
      <gva-chart :data="goodsData" title="商品数量" />
    </gva-card>
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/beeOrderTodo')"
    >
      <gva-chart :data="orderTodoData" title="订单待办" />
    </gva-card>
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/beeFinancialManager/beePayLog')"
    >
      <gva-chart :data="payData" title="总支付金额" />
    </gva-card>
    <gva-card
      custom-class="col-span-1 lg:col-span-2 h-32"
      @click="gotoPage('/bee_index/beeFinancialManager/beePayLog')"
    >
      <gva-chart :data="payNumData" title="支付人数" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-6" title="支付流水">
      <el-row>
        <el-col :span="4">
          <el-select v-model="shopId" placeholder="请选择门店">
            <el-option
              v-for="(item, index) in shops"
              :key="index"
              :value="item.id"
              :label="item.name"
              >{{ item.name }}</el-option
            >
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-date-picker
            v-model="dates"
            type="daterange"
            style="width: 100%; box-sizing: border-box"
            range-separator="到"
            start-placeholder="请选择开始日期"
            end-placeholder="请选择结束日期"
          />
        </el-col>
        <el-col
          :span="1"
          style="display: flex; justify-content: center; align-items: center"
        >
          <el-button @click="refreshOrders">
            <el-icon><Search /></el-icon>
          </el-button>
        </el-col>
      </el-row>
      <el-table
        :data="orders"
        stripe
        style="width: 100%; height: 200px"
        :span-method="spanMethodHandler"
      >
        <el-table-column prop="orderNumber" label="订单号" width="200" />
        <el-table-column
          prop="amount"
          label="订单金额"
          align="center"
          show-overflow-tooltip
        >
        </el-table-column>
        <el-table-column prop="datePay" label="支付时间">
          <template #default="scope">
            {{ dateFmt(scope.row.dateAdd) }}
          </template>
        </el-table-column>
      </el-table>
    </gva-card>

    <gva-card
      title="待办事项"
      custom-class="col-span-1 md:col-span-6 lg:col-span-6 row-span-2"
    >
      <div>
        <el-table
          :data="orderTodo.list"
          stripe
          style="width: 100%"
          @row-click="gotoTodoPage"
        >
          <el-table-column prop="orderNumber" label="订单号" width="200" />
          <el-table-column
            prop="id"
            label="事项"
            align="center"
            show-overflow-tooltip
          >
            <template #default="scope">
              {{ formatEnum(scope.row.status, beeOrderStatus) }}
            </template>
          </el-table-column>
          <el-table-column prop="amount" label="金额" width="100" />
          <el-table-column prop="dateAdd" label="支付时间">
            <template #default="scope">{{
              formatDate(scope.row.dateAdd)
            }}</template>
          </el-table-column>
        </el-table>
      </div>
    </gva-card>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import { GvaChart, GvaCard } from "./componenst";
import { dayjs } from "element-plus";
import { getBeeUserList } from "@/plugin/beeshop/api/beeUser";
import {
  getBeePayLogList,
  getBeePayTotal,
} from "@/plugin/beeshop/api/beePayLog";
import { getBeeOrderList, orderList } from "@/plugin/beeshop/api/beeOrder";
import { getBeeShopGoodsList } from "@/plugin/beeshop/api/beeShopGoods";
import { formatDate, formatEnum, getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { getAllBeeShopInfos } from "@/plugin/beeshop/api/beeShopInfo";
const route = useRoute();
const router = useRouter();

const beeOrderStatus = ref([]);
const userData = ref([]);
const goodsData = ref([]);
const orderData = ref([]);
const payData = ref([]);
const payNumData = ref([]);
const orderTodoData = ref([]);
const orderTodo = ref({});
const shops = ref([]);
const orders = ref([]);
const totalOrders = ref(0);
const orderSum = ref(0);
const pageNum = ref(0);
const pageSize = ref(10);
const dates = ref([]);
const shopId = ref("");
const init = async () => {
  beeOrderStatus.value = await getDictFunc("OrderStatus");
  for (let i = 7; i >= 0; i--) {
    // 获取最近7天数据
    const start = new Date(0);
    const end = dayjs()
      .add(-1 * i, "day")
      .endOf("day")
      .toDate();

    const funcs = [];
    funcs.push(
      getBeeUserList({
        page: 1,
        pageSize: 1,
        startDateAdd: start,
        endDateAdd: end,
      })
    );
    funcs.push(
      getBeePayLogList({
        page: 1,
        pageSize: 1,
        startDateAdd: start,
        endDateAdd: end,
        distinct: "uid",
      })
    );
    funcs.push(
      getBeeOrderList({
        page: 1,
        pageSize: 1,
        startDateAdd: start,
        endDateAdd: end,
      })
    );
    funcs.push(
      getBeeShopGoodsList({
        page: 1,
        pageSize: 1,
        startDateAdd: start,
        endDateAdd: end,
      })
    );
    funcs.push(
      getBeePayTotal({
        page: 1,
        pageSize: 1,
        startDateAdd: start,
        endDateAdd: end,
        sum: "money",
      })
    );
    const results = await Promise.all(funcs);

    const userTable = results[0];
    if (userTable.code === 0) {
      userData.value.push(userTable.data.total);
    }
    const payLogUidTable = results[1];
    if (payLogUidTable.code === 0) {
      payNumData.value.push(payLogUidTable.data.total);
    }
    const orderTable = await results[2];
    if (orderTable.code === 0) {
      orderData.value.push(orderTable.data.total);
    }
    const goodsTable = await results[3];
    if (goodsTable.code === 0) {
      goodsData.value.push(goodsTable.data.total);
    }
    const payDataTable = await results[4];
    if (payDataTable.code === 0) {
      payData.value.push(payDataTable.data.total);
    }
  }
  await refreshOrders();
  getAllBeeShopInfos().then((res) => {
    if (res.code == 0) {
      shops.value = res.data.list;
    }
  });
  const orderTable = await getBeeOrderList({
    page: 1,
    pageSize: 100,
    status: 1,
  });
  if (orderTable.code === 0) {
    orderTodoData.value.push(orderTable.data.total);
    orderTodoData.value.push(orderTable.data.total);
    orderTodo.value = orderTable.data;
  }
};
const spanMethodHandler = (row, column, rowIndex, columnIndex) => {
  console.log(row, column, rowIndex, columnIndex);
  // if(row - 1 == rowIndex && columnIndex == 0){
  //   return [1,2]
  // }else {
  //   return [0,0]
  // }
};
const dateFmt = (time) => {
  if (time) {
    return dayjs(time).format("YYYY-MM-DD hh:mm:ss");
  }
  return "";
};
init();
// 加载订单流水
const refreshOrders = async () => {
  const ordersRes =  await orderList({
    page: pageNum.value,
    pageSize: 10,
    shopId: shopId.value,
    startDateAdd:dates.value[0],
    endDateAdd:dates.value[1]
  });
  if (ordersRes.code == 0) {
    console.log(ordersRes.data.list, "234234");
    orders.value = [
      ...ordersRes.data.list,
      { amount: ordersRes.data.sum, orderNumber: "合计" },
    ];
    totalOrders.value = ordersRes.data.total;
  }

};
const gotoPage = async (page) => {
  page = "/layout" + page;
  await router.push(page);
};
// const formatDate = (time)=>{
//   return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
// }
const gotoTodoPage = (row) => {
  gotoPage("/bee_index/shop-order-admin/beeOrderDetail?id=" + row.id);
};
</script>

<style lang="scss" scoped></style>
