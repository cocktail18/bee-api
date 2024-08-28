<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="订单id" prop="orderId">

          <el-input v-model.number="searchInfo.orderId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="配送订单id" prop="peisongOrderId">
          <el-input v-model="searchInfo.peisongOrderId" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="id"
          @selection-change="handleSelectionChange"
          @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column sortable align="left" label="id字段" prop="id" width="120" />
        <el-table-column align="left" label="订单id" prop="orderId" width="120" />
        <el-table-column align="left" label="配送订单id" prop="peisongOrderId" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">{{ formatEnum(scope.row.status, beeOrderPeisongStatus) }}</template>
        </el-table-column>
        <el-table-column align="left" label="配送订单号" prop="peisongOrderNo" width="120" />
        <el-table-column align="left" label="配送类型" prop="type" width="120" >
          <template #default="scope">{{ formatEnum(scope.row.type, beeDelivery) }}</template>
        </el-table-column>
        <el-table-column align="left" label="配送费" prop="money" width="120" />
        <el-table-column align="left" label="实际支付金额" prop="realMoney" width="120" />
        <el-table-column align="left" label="取消配送扣费" prop="deductFee" width="120" />
        <el-table-column align="left" label="是否取消" prop="isCancel" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isCancel) }}</template>
        </el-table-column>
        <el-table-column align="left" label="配送错误信息" prop="errMsg" width="120" />
        <el-table-column align="left" label="第三方状态" prop="thirdStatus" width="120" />
        <el-table-column align="left" label="重试次数" prop="retryTimes" width="120" />
        <el-table-column align="left" label="上次重试时间" prop="lastRetryUnix" width="120" />
<!--        <el-table-column align="left" label="请求数据" prop="reqData" width="120" />-->
        <el-table-column align="left" label="已删除" prop="isDeleted" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isDeleted) }}</template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="dateAdd" width="180">
          <template #default="scope">{{ formatDate(scope.row.dateAdd) }}</template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" prop="dateUpdate" width="180">
          <template #default="scope">{{ formatDate(scope.row.dateUpdate) }}</template>
        </el-table-column>
        <el-table-column align="left" label="删除时间" prop="dateDelete" width="180">
          <template #default="scope">{{ formatDate(scope.row.dateDelete) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeOrderPeisongFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="id字段:"  prop="id" >
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
        </el-form-item>
<!--        <el-form-item label="已删除:"  prop="isDeleted" >-->
<!--          <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="创建时间:"  prop="dateAdd" >-->
<!--          <el-date-picker v-model="formData.dateAdd" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="更新时间:"  prop="dateUpdate" >-->
<!--          <el-date-picker v-model="formData.dateUpdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="删除时间:"  prop="dateDelete" >-->
<!--          <el-date-picker v-model="formData.dateDelete" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--        </el-form-item>-->
        <el-form-item label="订单id:"  prop="orderId" >
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入订单id" />
        </el-form-item>
        <el-form-item label="配送订单id:"  prop="peisongOrderId" >
          <el-input v-model="formData.peisongOrderId" :clearable="true"  placeholder="请输入配送订单id" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeOrderPeisongStatus" :key="key" :label="item.label" :value="parseInt(item.value)" />
          </el-select>
        </el-form-item>
        <el-form-item label="配送订单号:" prop="peisongOrderNo">
          <el-input v-model="formData.peisongOrderNo" :clearable="true"  placeholder="请输入配送订单号" />
        </el-form-item>
        <el-form-item label="配送类型:" prop="type">
          <el-select v-model="formData.type" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeDelivery" :key="key" :label="item.label" :value="parseInt(item.value)" />
          </el-select>
        </el-form-item>
        <el-form-item label="配送费:" prop="money">
          <el-input-number v-model="formData.money" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="实际支付金额:" prop="realMoney">
          <el-input-number v-model="formData.realMoney" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="取消配送扣费:" prop="deductFee">
          <el-input-number v-model="formData.deductFee" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="是否取消:" prop="isCancel">
          <el-switch v-model="formData.isCancel" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="配送错误信息:" prop="errMsg">
          <el-input v-model="formData.errMsg" :clearable="true"  placeholder="请输入配送错误信息" />
        </el-form-item>
        <el-form-item label="第三方状态:" prop="thirdStatus">
          <el-input v-model="formData.thirdStatus" :clearable="true"  placeholder="请输入第三方状态" />
        </el-form-item>
        <el-form-item label="重试次数:" prop="retryTimes">
          <el-input v-model.number="formData.retryTimes" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上次重试时间:" prop="lastRetryUnix">
          <el-input v-model.number="formData.lastRetryUnix" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="请求数据:" prop="reqData">
          <el-input type="textarea" v-model="formData.reqData" :clearable="true"  placeholder="请输入请求数据" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeOrderPeisong,
  deleteBeeOrderPeisong,
  deleteBeeOrderPeisongByIds,
  updateBeeOrderPeisong,
  findBeeOrderPeisong,
  getBeeOrderPeisongList
} from '@/plugin/beeshop/api/beeOrderPeisong'

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  filterDataSource,
  ReturnArrImg,
  onDownloadFile,
  formatEnum
} from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'BeeOrderPeisong'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

const beeOrderPeisongStatus = ref([])
const beeDelivery = ref([])

const init = async () => {
  beeOrderPeisongStatus.value = await getDictFunc('BeeOrderPeisongStatus')
  beeDelivery.value = await getDictFunc('BeeDelivery')
}
init()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  id: undefined,
  userId: undefined,
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
  orderId: undefined,
  peisongOrderId: '',
  status: undefined,
  peisongOrderNo: '',
  type: undefined,
  money: 0,
  realMoney: 0,
  deductFee: 0,
  isCancel: false,
  errMsg: '',
  thirdStatus: '',
  retryTimes: undefined,
  lastRetryUnix: undefined,
  reqData: '',
})



// 验证规则
const rule = reactive({
  orderId : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  },
  ],
  peisongOrderId : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  sort: 'id',
  order: 'descending',
})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    id: 'id',
  }

  let sort = sortMap[prop]
  if(!sort){
    sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {
    sort: 'id',
    order: 'descending',
  }
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.isDeleted === ""){
      searchInfo.value.isDeleted=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBeeOrderPeisongList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteBeeOrderPeisongFunc(row)
  })
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const ids = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.id)
    })
    const res = await deleteBeeOrderPeisongByIds({ ids })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBeeOrderPeisongFunc = async(row) => {
  const res = await findBeeOrderPeisong({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteBeeOrderPeisongFunc = async (row) => {
  const res = await deleteBeeOrderPeisong({ id: row.id })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    id: undefined,
    isDeleted: false,
    dateAdd: new Date(),
    dateUpdate: new Date(),
    dateDelete: new Date(),
    orderId: undefined,
    peisongOrderId: '',
    status: undefined,
    peisongOrderNo: '',
    type: undefined,
    money: 0,
    realMoney: 0,
    deductFee: 0,
    isCancel: false,
    errMsg: '',
    thirdStatus: '',
    retryTimes: undefined,
    lastRetryUnix: undefined,
    reqData: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBeeOrderPeisong(formData.value)
        break
      case 'update':
        res = await updateBeeOrderPeisong(formData.value)
        break
      default:
        res = await createBeeOrderPeisong(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>

