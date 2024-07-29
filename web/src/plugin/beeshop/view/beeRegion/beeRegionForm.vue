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
        <el-form-item label="firstLetter字段:" prop="firstLetter">
          <el-input v-model="formData.firstLetter" :clearable="true"  placeholder="请输入firstLetter字段" />
       </el-form-item>
        <el-form-item label="jianpin字段:" prop="jianpin">
          <el-input v-model="formData.jianpin" :clearable="true"  placeholder="请输入jianpin字段" />
       </el-form-item>
        <el-form-item label="level字段:" prop="level">
          <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="name字段:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入name字段" />
       </el-form-item>
        <el-form-item label="nameEn字段:" prop="nameEn">
          <el-input v-model="formData.nameEn" :clearable="true"  placeholder="请输入nameEn字段" />
       </el-form-item>
        <el-form-item label="pinyin字段:" prop="pinyin">
          <el-input v-model="formData.pinyin" :clearable="true"  placeholder="请输入pinyin字段" />
       </el-form-item>
        <el-form-item label="pid字段:" prop="pid">
          <el-input v-model="formData.pid" :clearable="true"  placeholder="请输入pid字段" />
       </el-form-item>
        <el-form-item label="regionId字段:" prop="regionId">
          <el-input v-model.number="formData.regionId" :clearable="true" placeholder="请输入" />
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
  createBeeRegion,
  updateBeeRegion,
  findBeeRegion
} from '@/plugin/beeshop/api/beeRegion'

defineOptions({
    name: 'BeeRegionForm'
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
            firstLetter: '',
            jianpin: '',
            level: undefined,
            name: '',
            nameEn: '',
            pinyin: '',
            pid: '',
            regionId: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeRegion({ ID: route.query.id })
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
               res = await createBeeRegion(formData.value)
               break
             case 'update':
               res = await updateBeeRegion(formData.value)
               break
             default:
               res = await createBeeRegion(formData.value)
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
