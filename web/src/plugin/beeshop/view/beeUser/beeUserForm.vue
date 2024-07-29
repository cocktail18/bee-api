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
        <el-form-item label="用户show_id:" prop="showUid">
          <el-input v-model.number="formData.showUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="头像:" prop="avatarUrl">
          <SelectImage
              v-model="formData.avatarUrl"
          />
          <el-input v-model="formData.avatarUrl" :clearable="true"  placeholder="请输入头像" />
       </el-form-item>
        <el-form-item label="会员卡号:" prop="cardNumber">
          <el-input v-model="formData.cardNumber" :clearable="true"  placeholder="请输入会员卡号" />
       </el-form-item>
        <el-form-item label="城市:" prop="city">
          <el-input v-model="formData.city" :clearable="true"  placeholder="请输入城市" />
       </el-form-item>
        <el-form-item label="登录时间:" prop="dateLogin">
          <el-date-picker v-model="formData.dateLogin" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-input v-model.number="formData.gender" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册ip:" prop="ipAdd">
          <el-input v-model="formData.ipAdd" :clearable="true"  placeholder="请输入注册ip" />
       </el-form-item>
        <el-form-item label="登录ip:" prop="ipLogin">
          <el-input v-model="formData.ipLogin" :clearable="true"  placeholder="请输入登录ip" />
       </el-form-item>
        <el-form-item label="人脸识别:" prop="isFaceCheck">
          <el-switch v-model="formData.isFaceCheck" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="身份证识别:" prop="isIdcardCheck">
          <el-switch v-model="formData.isIdcardCheck" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="管理员:" prop="isManager">
          <el-switch v-model="formData.isManager" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="销售人员:" prop="isSeller">
          <el-switch v-model="formData.isSeller" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="已发放注册优惠券:" prop="isSendRegisterCoupons">
          <el-switch v-model="formData.isSendRegisterCoupons" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="店长:" prop="isShopManager">
          <el-switch v-model="formData.isShopManager" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="团长:" prop="isTeamLeader">
          <el-switch v-model="formData.isTeamLeader" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="团员:" prop="isTeamMember">
          <el-switch v-model="formData.isTeamMember" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="关注用户:" prop="isUserAttendant">
          <el-switch v-model="formData.isUserAttendant" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否虚拟人:" prop="isVirtual">
          <el-switch v-model="formData.isVirtual" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="名字:" prop="nick">
          <el-input v-model="formData.nick" :clearable="true"  placeholder="请输入名字" />
       </el-form-item>
        <el-form-item label="所在省:" prop="province">
          <el-input v-model="formData.province" :clearable="true"  placeholder="请输入所在省" />
       </el-form-item>
        <el-form-item label="vip等级:" prop="vipLevel">
          <el-input v-model.number="formData.vipLevel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册来源:" prop="source">
          <el-input v-model.number="formData.source" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
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
  createBeeUser,
  updateBeeUser,
  findBeeUser
} from '@/plugin/beeshop/api/beeUser'

defineOptions({
    name: 'BeeUserForm'
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
            dateDelete: new Date(),
            showUid: undefined,
            avatarUrl: '',
            cardNumber: '',
            city: '',
            dateLogin: new Date(),
            gender: undefined,
            ipAdd: '',
            ipLogin: '',
            isFaceCheck: false,
            isIdcardCheck: false,
            isManager: false,
            isSeller: false,
            isSendRegisterCoupons: false,
            isShopManager: false,
            isTeamLeader: false,
            isTeamMember: false,
            isUserAttendant: false,
            isVirtual: false,
            nick: '',
            province: '',
            vipLevel: undefined,
            source: undefined,
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeUser({ ID: route.query.id })
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
               res = await createBeeUser(formData.value)
               break
             case 'update':
               res = await updateBeeUser(formData.value)
               break
             default:
               res = await createBeeUser(formData.value)
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
