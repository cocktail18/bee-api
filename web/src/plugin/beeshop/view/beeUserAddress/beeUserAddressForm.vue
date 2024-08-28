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
        <el-form-item label="详细地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入详细地址" />
       </el-form-item>
        <el-form-item label="地区:" prop="areaStr">
          <el-input v-model="formData.areaStr" :clearable="true"  placeholder="请输入地区" />
       </el-form-item>
        <el-form-item label="地区码:" prop="cityId">
          <el-input v-model="formData.cityId" :clearable="true"  placeholder="请输入地区码" />
       </el-form-item>
        <el-form-item label="城市:" prop="cityStr">
          <el-input v-model="formData.cityStr" :clearable="true"  placeholder="请输入城市" />
       </el-form-item>
        <el-form-item label="地区id:" prop="districtId">
          <el-input v-model="formData.districtId" :clearable="true"  placeholder="请输入地区id" />
       </el-form-item>
        <el-form-item label="默认地址:" prop="isDefault">
          <el-switch v-model="formData.isDefault" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="纬度:" prop="latitude">
          <el-input-number v-model="formData.latitude" :precision="6" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="联系人:" prop="linkMan">
          <el-input v-model="formData.linkMan" :clearable="true"  placeholder="请输入联系人" />
       </el-form-item>
        <el-form-item label="经度:" prop="longitude">
          <el-input-number v-model="formData.longitude" :precision="6" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="联系电话:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入联系电话" />
       </el-form-item>
        <el-form-item label="省代码:" prop="provinceId">
          <el-input v-model="formData.provinceId" :clearable="true"  placeholder="请输入省代码" />
       </el-form-item>
        <el-form-item label="省份:" prop="provinceStr">
          <el-input v-model="formData.provinceStr" :clearable="true"  placeholder="请输入省份" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
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
  createBeeUserAddress,
  updateBeeUserAddress,
  findBeeUserAddress
} from '@/plugin/beeshop/api/beeUserAddress'

defineOptions({
    name: 'BeeUserAddressForm'
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
            address: '',
            areaStr: '',
            cityId: '',
            cityStr: '',
            districtId: '',
            isDefault: false,
            latitude: 0,
            linkMan: '',
            longitude: 0,
            mobile: '',
            provinceId: '',
            provinceStr: '',
            status: undefined,
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
      const res = await findBeeUserAddress({ ID: route.query.id })
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
               res = await createBeeUserAddress(formData.value)
               break
             case 'update':
               res = await updateBeeUserAddress(formData.value)
               break
             default:
               res = await createBeeUserAddress(formData.value)
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
