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
        <el-form-item label="条码:" prop="barCode">
          <el-input v-model="formData.barCode" :clearable="true"  placeholder="请输入条码" />
       </el-form-item>
        <el-form-item label="售后说明,1仅退款，2退款退货，3换货:" prop="afterSale">
          <el-input v-model="formData.afterSale" :clearable="true"  placeholder="请输入售后说明,1仅退款，2退款退货，3换货" />
       </el-form-item>
        <el-form-item label="分类id:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分奖励:" prop="commission">
          <el-input-number v-model="formData.commission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="积分奖励结算类型:" prop="commissionSettleType">
          <el-input v-model.number="formData.commissionSettleType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分奖励类型:" prop="commissionType">
          <el-input v-model.number="formData.commissionType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分奖励用户类型:" prop="commissionUserType">
          <el-input v-model.number="formData.commissionUserType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="有附加:" prop="hasAddition">
          <el-switch v-model="formData.hasAddition" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="有团购:" prop="hasTourJourney">
          <el-switch v-model="formData.hasTourJourney" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="隐藏:" prop="hidden">
          <el-input v-model.number="formData.hidden" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="允许砍价:" prop="kanjia">
          <el-switch v-model="formData.kanjia" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="砍价价格:" prop="kanjiaPrice">
          <el-input-number v-model="formData.kanjiaPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="限制:" prop="limitation">
          <el-switch v-model="formData.limitation" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="物流id:" prop="logisticsId">
          <el-input v-model.number="formData.logisticsId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最大使用优惠券数量:" prop="maxCoupons">
          <el-input v-model.number="formData.maxCoupons" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="秒杀:" prop="miaosha">
          <el-switch v-model="formData.miaosha" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="最低购买数量:" prop="minBuyNumber">
          <el-input v-model.number="formData.minBuyNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最低价格:" prop="minPrice">
          <el-input-number v-model="formData.minPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="需要积分数量:" prop="minScore">
          <el-input-number v-model="formData.minScore" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商品名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入商品名字" />
       </el-form-item>
        <el-form-item label="喜欢数量:" prop="numberFav">
          <el-input v-model.number="formData.numberFav" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品评分:" prop="numberGoodReputation">
          <el-input-number v-model="formData.numberGoodReputation" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="订单数量:" prop="numberOrders">
          <el-input v-model.number="formData.numberOrders" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="评分数量:" prop="numberReputation">
          <el-input v-model.number="formData.numberReputation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="销售数量:" prop="numberSells">
          <el-input v-model.number="formData.numberSells" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="原价:" prop="originalPrice">
          <el-input-number v-model="formData.originalPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="会员价:" prop="vipPrice">
          <el-input-number v-model="formData.vipPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="海外直邮:" prop="overseas">
          <el-switch v-model="formData.overseas" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="排序:" prop="paixu">
          <el-input v-model.number="formData.paixu" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户:" prop="persion">
          <el-input v-model.number="formData.persion" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片:" prop="pic">
          <el-input v-model="formData.pic" :clearable="true"  placeholder="请输入图片" />
       </el-form-item>
        <el-form-item label="允许拼团:" prop="pingtuan">
          <el-switch v-model="formData.pingtuan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="拼团价格:" prop="pingtuanPrice">
          <el-input-number v-model="formData.pingtuanPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="属性id:" prop="propertyIds">
          <el-input v-model="formData.propertyIds" :clearable="true"  placeholder="请输入属性id" />
       </el-form-item>
        <el-form-item label="推荐状态:" prop="recommendStatus">
          <el-input v-model.number="formData.recommendStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="秒杀最低购买数量:" prop="seckillBuyNumber">
          <el-input v-model.number="formData.seckillBuyNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商店id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存:" prop="stores">
          <el-input v-model.number="formData.stores" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自动下架:" prop="stores0Unsale">
          <el-switch v-model="formData.stores0Unsale" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开始售卖时间:" prop="sellBeginTime">
          <el-date-picker v-model="formData.sellBeginTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束售卖时间:" prop="sellEndTime">
          <el-date-picker v-model="formData.sellEndTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="重量，kg:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
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
  createBeeShopGoods,
  updateBeeShopGoods,
  findBeeShopGoods
} from '@/plugin/beeshop/api/beeShopGoods'

defineOptions({
    name: 'BeeShopGoodsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from "@/components/selectImage/selectImage.vue";

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
            barCode: '',
            afterSale: '',
            categoryId: undefined,
            commission: 0,
            commissionSettleType: undefined,
            commissionType: undefined,
            commissionUserType: undefined,
            hasAddition: false,
            hasTourJourney: false,
            hidden: undefined,
            kanjia: false,
            kanjiaPrice: 0,
            limitation: false,
            logisticsId: undefined,
            maxCoupons: undefined,
            miaosha: false,
            minBuyNumber: undefined,
            minPrice: 0,
            minScore: 0,
            name: '',
            numberFav: undefined,
            numberGoodReputation: 0,
            numberOrders: undefined,
            numberReputation: undefined,
            numberSells: undefined,
            originalPrice: 0,
            vipPrice: 0,
            overseas: false,
            paixu: undefined,
            persion: undefined,
            pic: '',
            pingtuan: false,
            pingtuanPrice: 0,
            propertyIds: '',
            recommendStatus: undefined,
            seckillBuyNumber: undefined,
            shopId: undefined,
            status: undefined,
            stores: undefined,
            stores0Unsale: false,
            sellBeginTime: new Date(),
            sellEndTime: new Date(),
            weight: 0,
            type: undefined,
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
      const res = await findBeeShopGoods({ ID: route.query.id })
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
               res = await createBeeShopGoods(formData.value)
               break
             case 'update':
               res = await updateBeeShopGoods(formData.value)
               break
             default:
               res = await createBeeShopGoods(formData.value)
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
