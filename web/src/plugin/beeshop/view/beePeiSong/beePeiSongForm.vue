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
        <el-form-item label="距离，单位km:" prop="distance">
          <el-input-number v-model="formData.distance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最低收费金额:" prop="fwf1Min">
          <el-input-number v-model="formData.fwf1Min" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="收费项目名称:" prop="fwf1Name">
          <el-input-number v-model="formData.fwf1Name" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="比例或者金额:" prop="fwf1Number">
          <el-input-number v-model="formData.fwf1Number" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="类型:" prop="fwf1Type">
          <el-input v-model.number="formData.fwf1Type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自提补贴:" prop="ztDiscounts">
          <el-input-number v-model="formData.ztDiscounts" :precision="2" :clearable="true"></el-input-number>
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
  createBeePeiSong,
  updateBeePeiSong,
  findBeePeiSong
} from '@/plugin/beeshop/api/beePeiSong'

defineOptions({
    name: 'BeePeiSongForm'
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
            distance: 0,
            fwf1Min: 0,
            fwf1Name: 0,
            fwf1Number: 0,
            fwf1Type: undefined,
            ztDiscounts: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeePeiSong({ ID: route.query.id })
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
               res = await createBeePeiSong(formData.value)
               break
             case 'update':
               res = await updateBeePeiSong(formData.value)
               break
             default:
               res = await createBeePeiSong(formData.value)
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
