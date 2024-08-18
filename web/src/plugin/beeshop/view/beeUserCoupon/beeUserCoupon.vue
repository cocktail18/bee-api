<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

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
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="id字段" prop="id" width="120" />
        <el-table-column align="left" label="用户id" prop="uid" width="120" />
         <el-table-column align="left" label="可以使用时间" prop="dateStart" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateStart) }}</template>
         </el-table-column>
        <el-table-column align="left" label="过期时间（毫秒）" prop="expiryMillis" width="120" />
        <el-table-column align="left" label="金额" prop="money" width="120" />
        <el-table-column align="left" label="满xx可用" prop="moneyHreshold" width="120" />
        <el-table-column align="left" label="优惠券面额范围" prop="moneyMin" width="120" />
        <el-table-column align="left" label="优惠券面额范围" prop="moneyMax" width="120" />
        <el-table-column align="left" label="优惠券类型" prop="moneyType" width="120" />
        <el-table-column align="left" label="优惠券名字" prop="name" width="120" />
        <el-table-column align="left" label="仅抵扣运费" prop="onlyFreight" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.onlyFreight) }}</template>
        </el-table-column>
        <el-table-column align="left" label="优惠券id" prop="pid" width="120" />
        <el-table-column align="left" label="店铺id" prop="shopId" width="120" />
        <el-table-column align="left" label="管理员批量发放" prop="source" width="120" />
        <el-table-column align="left" label="使用状态，0未使用，1使用中，2已使用" prop="status" width="120" />
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
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeUserCouponFunc(scope.row)">变更</el-button>
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
            <!--            <el-form-item label="已删除:"  prop="isDeleted" >-->
<!--              <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="创建时间:"  prop="dateAdd" >-->
<!--              <el-date-picker v-model="formData.dateAdd" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="更新时间:"  prop="dateUpdate" >-->
<!--              <el-date-picker v-model="formData.dateUpdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="删除时间:"  prop="dateDelete" >-->
<!--              <el-date-picker v-model="formData.dateDelete" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
            <el-form-item label="用户id:"  prop="uid" >
              <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入用户id" />
            </el-form-item>
            <el-form-item label="可以使用时间:"  prop="dateStart" >
              <el-date-picker v-model="formData.dateStart" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="过期时间（毫秒）:"  prop="expiryMillis" >
              <el-input v-model.number="formData.expiryMillis" :clearable="true" placeholder="请输入过期时间（毫秒）" />
            </el-form-item>
            <el-form-item label="金额:"  prop="money" >
              <el-input-number v-model="formData.money"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="满xx可用:"  prop="moneyHreshold" >
              <el-input v-model.number="formData.moneyHreshold" :clearable="true" placeholder="请输入满xx可用" />
            </el-form-item>
            <el-form-item label="优惠券面额范围:"  prop="moneyMin" >
              <el-input v-model.number="formData.moneyMin" :clearable="true" placeholder="请输入优惠券面额范围" />
            </el-form-item>
            <el-form-item label="优惠券面额范围:"  prop="moneyMax" >
              <el-input v-model.number="formData.moneyMax" :clearable="true" placeholder="请输入优惠券面额范围" />
            </el-form-item>
            <el-form-item label="优惠券类型:"  prop="moneyType" >
              <el-input v-model.number="formData.moneyType" :clearable="true" placeholder="请输入优惠券类型" />
            </el-form-item>
            <el-form-item label="优惠券名字:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入优惠券名字" />
            </el-form-item>
            <el-form-item label="仅抵扣运费:"  prop="onlyFreight" >
              <el-switch v-model="formData.onlyFreight" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="优惠券id:"  prop="pid" >
              <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入优惠券id" />
            </el-form-item>
            <el-form-item label="店铺id:"  prop="shopId" >
              <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入店铺id" />
            </el-form-item>
            <el-form-item label="管理员批量发放:"  prop="source" >
              <el-input v-model="formData.source" :clearable="true"  placeholder="请输入管理员批量发放" />
            </el-form-item>
            <el-form-item label="使用状态，0未使用，1使用中，2已使用:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入使用状态，0未使用，1使用中，2已使用" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeUserCoupon,
  deleteBeeUserCoupon,
  deleteBeeUserCouponByIds,
  updateBeeUserCoupon,
  findBeeUserCoupon,
  getBeeUserCouponList
} from '@/plugin/beeshop/api/beeUserCoupon'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'BeeUserCoupon'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        id: undefined,
        isDeleted: false,
        dateAdd: new Date(),
        dateUpdate: new Date(),
        dateDelete: undefined,
        uid: undefined,
        dateStart: new Date(),
        expiryMillis: undefined,
        money: 0,
        moneyHreshold: undefined,
        moneyMin: undefined,
        moneyMax: undefined,
        moneyType: undefined,
        name: '',
        onlyFreight: false,
        pid: undefined,
        shopId: undefined,
        source: '',
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
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
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
    if (searchInfo.value.onlyFreight === ""){
        searchInfo.value.onlyFreight=null
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
  const table = await getBeeUserCouponList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteBeeUserCouponFunc(row)
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
      const res = await deleteBeeUserCouponByIds({ ids })
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
const updateBeeUserCouponFunc = async(row) => {
    const res = await findBeeUserCoupon({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBeeUserCouponFunc = async (row) => {
    const res = await deleteBeeUserCoupon({ id: row.id })
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
        dateDelete: undefined,
        uid: undefined,
        dateStart: new Date(),
        expiryMillis: undefined,
        money: 0,
        moneyHreshold: undefined,
        moneyMin: undefined,
        moneyMax: undefined,
        moneyType: undefined,
        name: '',
        onlyFreight: false,
        pid: undefined,
        shopId: undefined,
        source: '',
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
                  res = await createBeeUserCoupon(formData.value)
                  break
                case 'update':
                  res = await updateBeeUserCoupon(formData.value)
                  break
                default:
                  res = await createBeeUserCoupon(formData.value)
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
