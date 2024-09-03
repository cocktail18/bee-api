
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="id字段" prop="id">

          <el-input v-model.number="searchInfo.id" placeholder="搜索条件" />

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
        <el-table-column align="left" label="排队名" prop="name" width="120" />
        <el-table-column align="left" label="当前叫号" prop="curNumber" width="120" />
        <el-table-column align="left" label="当前取号" prop="numberGet" width="120" />
        <el-table-column align="left" label="最大叫号" prop="maxNumber" width="120" />
        <el-table-column align="left" label="每个号预计等待的时间，分钟" prop="minutes" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" >
          <template #default="scope">{{ formatEnum(scope.row.status, beeQueueStatus) }}</template>
        </el-table-column>
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
            <el-button type="text" link @click="showCallQueue(scope.row)">叫号台</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeQueueFunc(scope.row)">变更</el-button>
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
        <el-form-item label="排队名:"  prop="name" >
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入排队名" />
        </el-form-item>
        <el-form-item label="当前叫号:"  prop="curNumber" >
          <el-input v-model.number="formData.curNumber" :clearable="true" placeholder="请输入当前叫号" />
        </el-form-item>
        <el-form-item label="当前取号:"  prop="numberGet" >
          <el-input v-model.number="formData.numberGet" :clearable="true" placeholder="请输入当前取号" />
        </el-form-item>
        <el-form-item label="最大叫号:"  prop="maxNumber" >
          <el-input v-model.number="formData.maxNumber" :clearable="true" placeholder="请输入最大叫号" />
        </el-form-item>
        <el-form-item label="每个号预计等待的时间，分钟:"  prop="minutes" >
          <el-input v-model.number="formData.minutes" :clearable="true" placeholder="请输入每个号预计等待的时间，分钟" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeQueueStatus" :key="key" :label="item.label" :value="parseInt(item.value)" />
          </el-select>
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
      </el-form>
    </el-drawer>

    <el-dialog v-model="queueCallFormVisible" title="叫号台" width="500">
      <el-form :model="queueCallForm">
        <el-form-item label="当前叫号：（点叮咚会更新当前叫号）" prop="curNumber">
          <el-input v-model="queueCallForm.curNumber" autocomplete="off" />
        </el-form-item>
        <el-form-item label="等待人数：">
          <el-input readonly v-model="queueCallForm.waitingNumber" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="queueCallFormVisible = false">Cancel</el-button>
          <el-button type="primary" @click="dingDong">
            叮咚
          </el-button>
          <el-button type="primary" @click="reCall">
            重呼
          </el-button>
          <el-button type="primary" @click="passUser">
            无人过号
          </el-button>
          <el-button type="primary" @click="nextUser">
            下一位
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createBeeQueue,
  deleteBeeQueue,
  deleteBeeQueueByIds,
  updateBeeQueue,
  findBeeQueue,
  getBeeQueueList, callBeeQueue, reCallBeeQueue, passBeeQueue, nextBeeQueue
} from '@/plugin/beeshop/api/beeQueue'

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
import SpeakTTS from "speak-tts";

defineOptions({
  name: 'BeeQueue'
})

const speakTTS = new SpeakTTS();
if (speakTTS.hasBrowserSupport()) {
  speakTTS.init().then(() => {
    console.log('SpeakTTS is ready');
  }).catch(e => {
    console.error('An error occurred while initializing SpeakTTS:', e);
  });
} else {
  console.warn('Browser does not support SpeakTTS');
}

const speakText = (text) => {
  if (speakTTS.hasBrowserSupport()) {
    speakTTS.speak({
      text: text,
      queue: false,
    }).then(() => {
      console.log('Text has been spoken');
    }).catch(e => {
      console.error('An error occurred while speaking text:', e);
    });
  }
}

const beeQueueStatus = ref([])
const init = async () => {
  // 获取字典（可能为空）
  beeQueueStatus.value = await getDictFunc('BeeQueueStatus')
}
init()


// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)
const queueCallFormVisible = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  id: undefined,
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
  name: '',
  curNumber: undefined,
  numberGet: undefined,
  maxNumber: undefined,
  minutes: undefined,
  status: undefined,
})

const queueCallForm = ref({
  id: undefined,
  name: '',
  curNumber: undefined,
  waitingNumber: 0,
  numberGet: undefined,
  maxNumber: undefined,
  minutes: undefined,
  status: undefined,
})



// 验证规则
const rule = reactive({
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
  const table = await getBeeQueueList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteBeeQueueFunc(row)
  })
}

// 删除行
const showCallQueue = async (row) => {
  const res = await findBeeQueue({id: row.id})
  if (res.code === 0) {
    queueCallForm.value = res.data
    queueCallForm.value.waitingNumber = res.data.numberGet - res.data.curNumber
    queueCallFormVisible.value = true
  }
}

const dingDong = async () => {
  const res = await callBeeQueue(queueCallForm.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '操作成功'
    })
    await showCallQueue({id: queueCallForm.value.id})
    speakText(queueCallForm.value.curNumber + '号')
  }
}

const reCall = async () => {
  const res = await reCallBeeQueue(queueCallForm.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '操作成功'
    })
    await showCallQueue({id: queueCallForm.value.id})
    speakText(queueCallForm.value.curNumber + '号')
  }
}

const passUser = async () => {
  const res = await passBeeQueue(queueCallForm.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '操作成功'
    })
    await showCallQueue({id: queueCallForm.value.id})
    speakText(queueCallForm.value.curNumber + '号')
  }
}

const nextUser = async () => {
  const res = await nextBeeQueue(queueCallForm.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '操作成功'
    })
    await showCallQueue({id: queueCallForm.value.id})
    speakText(queueCallForm.value.curNumber + '号')
  }
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
    const res = await deleteBeeQueueByIds({ ids })
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
const updateBeeQueueFunc = async(row) => {
  const res = await findBeeQueue({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteBeeQueueFunc = async (row) => {
  const res = await deleteBeeQueue({ id: row.id })
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
    name: '',
    curNumber: undefined,
    numberGet: undefined,
    maxNumber: undefined,
    minutes: undefined,
    status: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBeeQueue(formData.value)
        break
      case 'update':
        res = await updateBeeQueue(formData.value)
        break
      default:
        res = await createBeeQueue(formData.value)
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

