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
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="绑定用户id:" prop="bindUid">
          <el-input v-model.number="formData.bindUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="城市id:" prop="cityId">
          <el-input v-model="formData.cityId" :clearable="true"  placeholder="请输入城市id" />
       </el-form-item>
        <el-form-item label="支付方式:" prop="cyTablePayMod">
          <el-input v-model.number="formData.cyTablePayMod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="唯一id:" prop="districtId">
          <el-input v-model="formData.districtId" :clearable="true"  placeholder="请输入唯一id" />
       </el-form-item>
        <el-form-item label="快递类型:" prop="expressType">
          <el-input v-model="formData.expressType" :clearable="true"  placeholder="请输入快递类型" />
       </el-form-item>
        <el-form-item label="商品需要检查:" prop="goodsNeedCheck">
          <el-switch v-model="formData.goodsNeedCheck" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="纬度:" prop="latitude">
          <el-input-number v-model="formData.latitude" :precision="10" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="绑定手机:" prop="linkPhone">
          <el-input v-model="formData.linkPhone" :clearable="true"  placeholder="请输入绑定手机" />
       </el-form-item>
        <el-form-item label="经度:" prop="longitude">
          <el-input-number v-model="formData.longitude" :precision="10" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="店名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入店名" />
       </el-form-item>
        <el-form-item label="座机号码:" prop="number">
          <el-input v-model="formData.number" :clearable="true"  placeholder="请输入座机号码" />
       </el-form-item>
        <el-form-item label="喜欢人数:" prop="numberFav">
          <el-input v-model.number="formData.numberFav" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品评分:" prop="numberGoodReputation">
          <el-input v-model.number="formData.numberGoodReputation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下单数:" prop="numberOrder">
          <el-input v-model.number="formData.numberOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺评分:" prop="numberReputation">
          <el-input v-model.number="formData.numberReputation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开启扫描:" prop="openScan">
          <el-switch v-model="formData.openScan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开启外卖:" prop="openWaimai">
          <el-switch v-model="formData.openWaimai" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开启自取:" prop="openZiqu">
          <el-switch v-model="formData.openZiqu" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开店时间:" prop="openingHours">
          <el-input v-model="formData.openingHours" :clearable="true"  placeholder="请输入开店时间" />
       </el-form-item>
        <el-form-item label="排序:" prop="paixu">
          <el-input v-model.number="formData.paixu" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="省id:" prop="provinceId">
          <el-input v-model="formData.provinceId" :clearable="true"  placeholder="请输入省id" />
       </el-form-item>
        <el-form-item label="推荐状态:" prop="recommendStatus">
          <el-input v-model.number="formData.recommendStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务距离:" prop="serviceDistance">
          <el-input-number v-model="formData.serviceDistance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发票:" prop="taxGst">
          <el-input v-model.number="formData.taxGst" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发票服务:" prop="taxService">
          <el-input v-model.number="formData.taxService" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入店铺类型" />
       </el-form-item>
        <el-form-item label="营业状态:" prop="workStatus">
          <el-input v-model.number="formData.workStatus" :clearable="true" placeholder="请输入" />
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
  createBeeShopInfo,
  updateBeeShopInfo,
  findBeeShopInfo
} from '@/plugin/beeshop/api/beeShopInfo'

defineOptions({
    name: 'BeeShopInfoForm'
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
            address: '',
            bindUid: undefined,
            cityId: '',
            cyTablePayMod: undefined,
            districtId: '',
            expressType: '',
            goodsNeedCheck: false,
            latitude: 0,
            linkPhone: '',
            longitude: 0,
            name: '',
            number: '',
            numberFav: undefined,
            numberGoodReputation: undefined,
            numberOrder: undefined,
            numberReputation: undefined,
            openScan: false,
            openWaimai: false,
            openZiqu: false,
            openingHours: '',
            paixu: undefined,
            provinceId: '',
            recommendStatus: undefined,
            serviceDistance: 0,
            status: undefined,
            taxGst: undefined,
            taxService: undefined,
            type: '',
            workStatus: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeShopInfo({ ID: route.query.id })
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
               res = await createBeeShopInfo(formData.value)
               break
             case 'update':
               res = await updateBeeShopInfo(formData.value)
               break
             default:
               res = await createBeeShopInfo(formData.value)
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
