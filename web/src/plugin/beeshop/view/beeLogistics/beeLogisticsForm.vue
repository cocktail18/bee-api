<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入"/>
        </el-form-item>
        <!--        <el-form-item label="商店用户ID:" prop="userId">-->
        <!--          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />-->
        <!--       </el-form-item>-->
        <el-form-item label="模板名:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入模板名"/>
        </el-form-item>
        <el-form-item label="计价方式:" prop="feeType">
          <el-select v-model="formData.feeType" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeLogisticType" :key="key" :label="item.label"
                       :value="parseInt(item.value)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="是否包邮:" prop="isFree">
          <el-switch v-model="formData.isFree" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <!--        <el-form-item label="已删除:" prop="isDeleted">-->
        <!--          <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
        <!--       </el-form-item>-->
        <!--        <el-form-item label="创建时间:" prop="dateAdd">-->
        <!--          <el-date-picker v-model="formData.dateAdd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>-->
        <!--       </el-form-item>-->
        <!--        <el-form-item label="更新时间:" prop="dateUpdate">-->
        <!--          <el-date-picker v-model="formData.dateUpdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>-->
        <!--       </el-form-item>-->
        <!--        <el-form-item label="删除时间:" prop="dateDelete">-->
        <!--          <el-date-picker v-model="formData.dateDelete" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>-->
        <!--       </el-form-item>-->
        <el-divider content-position="left">运费设置：</el-divider>
        <el-row style="width: 100%">
          <el-col :span="3" class="grid-cell">默认:</el-col>
          <el-col :span="17" class="grid-cell">
            <div style="display: flex">
              <el-input v-model.number="logisticsDetailDefault.firstNumber" :clearable="true" placeholder="请输入">
                <template #append>{{ formatEnum(formData.feeType, beeLogisticsUnit) }}内</template>
              </el-input>
              <el-input v-model.number="logisticsDetailDefault.firstAmount" :clearable="true" placeholder="请输入">
                <template #append>元</template>
              </el-input>
              <el-input v-model.number="logisticsDetailDefault.addNumber" :clearable="true" placeholder="请输入">
                <template #prepend>每增加</template>
                <template #append>{{ formatEnum(formData.feeType, beeLogisticsUnit) }}</template>
              </el-input>
              <el-input v-model.number="logisticsDetailDefault.addAmount" :clearable="true" placeholder="请输入">
                <template #prepend>增加运费</template>
                <template #append>元</template>
              </el-input>
            </div>
          </el-col>
        </el-row>
        <el-row style="width: 100%" v-for="(item,key) in logisticsExceptionList" :key="key">
          <el-col :span="3" class="grid-cell text">{{ item.names.join(",") }}:</el-col>
          <el-col :span="17" class="grid-cell">
            <div style="display: flex">
              <el-input v-model.number="item.firstNumber" :clearable="true" placeholder="请输入">
                <template #append>{{ formatEnum(formData.feeType, beeLogisticsUnit) }}内</template>
              </el-input>
              <el-input v-model.number="item.firstAmount" :clearable="true" placeholder="请输入">
                <template #append>元</template>
              </el-input>
              <el-input v-model.number="item.addNumber" :clearable="true" placeholder="请输入">
                <template #prepend>每增加</template>
                <template #append>{{ formatEnum(formData.feeType, beeLogisticsUnit) }}</template>
              </el-input>
              <el-input v-model.number="item.addAmount" :clearable="true" placeholder="请输入">
                <template #prepend>增加运费</template>
                <template #append>元</template>
              </el-input>
<!--              <el-input v-model="item.remark" :clearable="true" placeholder="请输入备注"/>-->
            </div>
          </el-col>
          <el-col :span="4" class="grid-cell">
            <el-form-item>
              <el-button size="small" type="text" @click="openSetLogisticsExceptionRegionDialogForUpdate(key)">修改</el-button>
              <el-button size="small" type="text" @click="delLogisticsException(key)" style="color: red">删除</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-button type="text" @click="openSetLogisticsExceptionRegionDialogForUpdate(-1)">为指定地区设置运费</el-button>
        <el-drawer destroy-on-close size="800" v-model="dialogSetLogisticsExceptionRegionFormVisible" :show-close="false" :before-close="closeSetExceptionLogisticsRegionFormDialog">
          <Suspense>
          <BeeRegionTree :regionTreeCheckedIds="regionTreeLogisticsExceptionCheckedIds" @on-ok="onSetLogisticsExceptionRegionFormDialogOk"></BeeRegionTree>
          </Suspense>
        </el-drawer>
        <el-divider content-position="left">条件包邮设置：</el-divider>
        <el-row style="width: 100%">
          <el-col :span="3" class="grid-cell">默认:</el-col>
          <el-col :span="12" class="grid-cell">
            <div style="display: flex">
              <el-input v-model.number="freeShippingDefault.number" :clearable="true" placeholder="请输入">
                <template #prepend>满</template>
              </el-input>
              <el-select v-model="freeShippingDefault.type" clearable placeholder="请选择" :clearable="false">
                <el-option v-for="(item,key) in beeLogisticsUnit" :key="key" :label="item.label"
                           :value="parseInt(item.value)"/>
              </el-select>
              <span style="width: 100px">包邮</span>
            </div>
          </el-col>
        </el-row>
        <el-row style="width: 100%" v-for="(item,key) in freeShippingList" :key="key">
          <el-col :span="3" class="grid-cell text">{{ item.names.join(",") }}</el-col>
          <el-col :span="12" class="grid-cell">
            <div style="display: flex">
              <el-input v-model.number="item.number" :clearable="true" placeholder="请输入">
                <template #prepend>满</template>
              </el-input>
              <el-select v-model="item.type" clearable placeholder="请选择" :clearable="false">
                <el-option v-for="(item,key) in beeLogisticsUnit" :key="key" :label="item.label"
                           :value="parseInt(item.value)"/>
              </el-select>
<!--              <el-input v-model="item.remark" :clearable="true" placeholder="请输入备注"/>-->
              <span style="width: 100px">包邮</span>
            </div>
          </el-col>
          <el-col :span="4" class="grid-cell">
            <el-form-item>
              <el-button size="small" type="text" @click="onSetLogisticsExceptionRegionFormDialogOk(key)">修改</el-button>
              <el-button size="small" type="text" @click="delFreeShipping(key)" style="color: red">删除</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-button type="text" @click="openSetFreeShippingFormDialogForUpdate(-1)">为指定地区设置包邮规则</el-button>
        <el-drawer destroy-on-close size="800" v-model="dialogSetFreeShippingFormVisible" :show-close="false" :before-close="closeSetFreeShippingFormDialog">
          <Suspense>
            <BeeRegionTree :regionTreeCheckedIds="regionTreeFreeShippingCheckedIds" @on-ok="onSetFreeShippingRegionFormDialogOk"></BeeRegionTree>
          </Suspense>
        </el-drawer>
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
  createBeeLogistics,
  updateBeeLogistics,
  findBeeLogistics
} from '@/plugin/beeshop/api/beeLogistics'

defineOptions({
  name: 'BeeLogisticsForm'
})

// 自动获取字典
import {formatEnum, getDictFunc} from '@/utils/format'
import {useRoute, useRouter} from "vue-router"
import {ElMessage} from 'element-plus'
import {ref, reactive} from 'vue'
import BeeRegionTree from '../components/beeRegionTree/bee-region-tree.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  id: undefined,
  userId: undefined,
  name: '',
  feeType: undefined,
  isFree: false,
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
  freeShippingSetting: '',
  detailsJson: '',
})
// 验证规则
const rule = reactive({})

const elFormRef = ref()
const beeLogisticType = ref([])
const beeLogisticsUnit = ref([])
const logisticsExceptionList = ref([])
const freeShippingDefault = ref({
  default: true,
  names: ["默认"],
  cityIds: ["0"],
  type: 0,
  number: 5,
  regionId: "0",
  remark: ""
})
const freeShippingList = ref([])
const logisticsDetailDefault = ref({
  default: true,
  names: ["默认运费"],
  cityIds: ["0"],
  firstNumber: 1,
  firstAmount: 5,
  addNumber: 1,
  addAmount: 5,
  type: 0,
  cityId: "0",
  remark: ""
})
// 初始化方法
const init = async () => {
  beeLogisticType.value = await getDictFunc('BeeLogisticsType')
  beeLogisticsUnit.value = await getDictFunc('BeeLogisticsUnit')
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例

  if (route.query.id) {
    const logisticsId = route.query.id
    const res = await findBeeLogistics({id: logisticsId})
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
    const detailList = JSON.parse(res.data.detailsJson)
    const freeShippingSettingList = JSON.parse(res.data.freeShippingSetting)
    const _logisticsExceptionList = []
    const _freeShippingSettingList = []
    detailList.forEach((item) => {
      if (item.default) {
        logisticsDetailDefault.value = item
      } else {
        _logisticsExceptionList.push(item)
      }
    })
    logisticsExceptionList.value = _logisticsExceptionList

    freeShippingSettingList.forEach((item) => {
      if (item.default) {
        freeShippingDefault.value = item
      } else {
        _freeShippingSettingList.push(item)
      }
    })
    freeShippingList.value = _freeShippingSettingList
  } else {
    type.value = 'create'
  }
}

init()
const dialogSetLogisticsExceptionRegionFormVisible = ref(false)
const dialogSetFreeShippingFormVisible = ref(false)
const checkedLogisticsExceptionIndex = ref(0)
const checkedFreeShippingIndex = ref(0)

const regionTreeLogisticsExceptionCheckedIds = ref([])
const regionTreeFreeShippingCheckedIds = ref([])
const closeSetExceptionLogisticsRegionFormDialog = () => {
  dialogSetLogisticsExceptionRegionFormVisible.value = false
}

const closeSetFreeShippingFormDialog = () => {
  dialogSetFreeShippingFormVisible.value = false
}

// 更新行
const openSetLogisticsExceptionRegionDialogForUpdate = async(index) => {
  checkedLogisticsExceptionIndex.value = index
  if (index >= 0) {
    regionTreeLogisticsExceptionCheckedIds.value = logisticsExceptionList.value[index].cityIds
  }
  openSetLogisticsExceptionRegionDialog()
}

// 更新行
const openSetFreeShippingFormDialogForUpdate = async(index) => {
  checkedFreeShippingIndex.value = index
  if (index >= 0) {
    regionTreeFreeShippingCheckedIds.value = freeShippingList.value[index].cityIds
  }
  openSetFreeShippingRegionDialog()
}

const openSetLogisticsExceptionRegionDialog = () => {
  dialogSetLogisticsExceptionRegionFormVisible.value = true
}

const openSetFreeShippingRegionDialog = () => {
  dialogSetFreeShippingFormVisible.value = true
}

const delLogisticsException = (index) => {
  logisticsExceptionList.value.splice(index, 1)
}
const delFreeShipping = (index) => {
  freeShippingList.value.splice(index, 1)
}

const onSetLogisticsExceptionRegionFormDialogOk = async (regions) => {
  const checkIds = []
  const regionStr = []
  regions.forEach(region => {
    checkIds.push(region.id)
    regionStr.push(region.name)
  })
  if (checkedLogisticsExceptionIndex.value && checkedLogisticsExceptionIndex.value >= 0) {
    logisticsExceptionList.value[checkedLogisticsExceptionIndex].cityIds = checkIds
    logisticsExceptionList.value[checkedLogisticsExceptionIndex].names = regionStr
  } else {
    logisticsExceptionList.value.push({
      cityIds: checkIds,
      names: regionStr,
      default: false,
      firstNumber: 1,
      firstAmount: 5,
      addNumber: 1,
      addAmount: 5,
    })
  }
  closeSetExceptionLogisticsRegionFormDialog()
}

const onSetFreeShippingRegionFormDialogOk = async (regions) => {
  const checkIds = []
  const regionStr = []
  regions.forEach(region => {
    checkIds.push(region.id)
    regionStr.push(region.name)
  })
  if (checkedFreeShippingIndex.value && checkedFreeShippingIndex.value >= 0) {
    freeShippingList.value[checkedFreeShippingIndex].cityIds = checkIds
    freeShippingList.value[checkedFreeShippingIndex].names = regionStr
  } else {
    freeShippingList.value.push({
      cityIds: checkIds,
      names: regionStr,
      default: false,
      type: 0,
      number: 0,
    })
  }
  closeSetFreeShippingFormDialog()
}

// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    let details = []
    let freeShippingSettings = []
    if (logisticsDetailDefault.value) {
      logisticsDetailDefault.default = true
      logisticsDetailDefault.type = formData.value.type
      logisticsDetailDefault.cityIds = []
      logisticsDetailDefault.names = []
      details.push(logisticsDetailDefault.value)
    }
    logisticsExceptionList.value.forEach((item) => {
      item.default = false
      item.type = formData.value.type
      details.push(item)
    })
    if (freeShippingDefault.value) {
      freeShippingDefault.default = true
      freeShippingDefault.cityIds = []
      freeShippingDefault.names = []
      freeShippingSettings.push(freeShippingDefault.value)
    }
    freeShippingList.value.forEach((item) => {
      item.default = false
      freeShippingSettings.push(item)
    })
    formData.value.detailsJson = JSON.stringify(details)
    formData.value.freeShippingSetting = JSON.stringify(freeShippingSettings)
    switch (type.value) {
      case 'create':
        res = await createBeeLogistics(formData.value)
        break
      case 'update':
        res = await updateBeeLogistics(formData.value)
        break
      default:
        res = await createBeeLogistics(formData.value)
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

.text {
  white-space: nowrap; /* 防止文本换行 */
  text-overflow: ellipsis; /* 文本溢出时显示省略号 */
  overflow: hidden; /* 兼容性处理 */
}
</style>
