<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="添加时间" prop="dateAdd">
            
            <template #label>
            <span>
              添加时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startDateAdd" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endDateAdd ? time.getTime() > searchInfo.endDateAdd.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDateAdd" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startDateAdd ? time.getTime() < searchInfo.startDateAdd.getTime() : false"></el-date-picker>

        </el-form-item>

        <el-form-item label="订单id" prop="orderId">
            
             <el-input v-model.number="searchInfo.orderId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="是否已支付" prop="purchase">
         <el-input v-model="searchInfo.purchase" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="退款状态" prop="refundStatus">
            
             <el-input v-model.number="searchInfo.refundStatus" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="用户id" prop="uid">
            
             <el-input v-model.number="searchInfo.uid" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="更新时间" prop="dateUpdate">

            <template #label>
            <span>
              更新时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
            </template>
            <el-date-picker v-model="searchInfo.startDateUpdate" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endDateUpdate ? time.getTime() > searchInfo.endDateUpdate.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDateUpdate" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startDateUpdate ? time.getTime() < searchInfo.startDateUpdate.getTime() : false"></el-date-picker>

          </el-form-item>
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
<!--            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
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
        
        <el-table-column align="left" sortable label="id字段" prop="id" width="120" />
        <el-table-column align="left" sortable label="订单id" prop="orderId" width="120" />
        <el-table-column align="left" label="用户id" prop="uid" width="120" />
        <el-table-column align="left" label="售后类型" prop="afterSale" width="120" />
        <el-table-column align="left" label="总金额" prop="amount" width="120" />
        <el-table-column align="left" label="优惠券金额" prop="amountCoupon" width="120" />
        <el-table-column align="left" label="单价" prop="amountSingle" width="120" />
        <el-table-column align="left" label="基本单价" prop="amountSingleBase" width="120" />
        <el-table-column align="left" label="奖励发放结束" prop="buyRewardEnd" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.buyRewardEnd) }}</template>
        </el-table-column>
        <el-table-column align="left" label="分类" prop="categoryId" width="120" />
        <el-table-column align="left" label="餐饮桌子状态" prop="cyTableStatus" width="120" >
          <template #default="scope">{{ formatEnum(scope.row.cyTableStatus, beeCyTableStatus) }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="返现类型" prop="fxType" width="120" />-->
        <el-table-column align="left" label="商品id" prop="goodsId" width="120" />
        <el-table-column align="left" label="商品名" prop="goodsName" width="120" />
        <el-table-column align="left" label="是否积分订单" prop="isScoreOrder" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isScoreOrder) }}</template>
        </el-table-column>
        <el-table-column align="left" label="数量" prop="number" width="120" />
        <el-table-column align="left" label="未发货数量" prop="numberNoFahuo" width="120" />

        <el-table-column align="left" label="人数" prop="persion" width="120" />
        <el-table-column align="left" label="商品图" prop="pic" width="120">
          <template #default="scope">
            <CustomPic
                pic-type="file"
                style="margin-top:8px"
                :pic-src="scope.row.pic"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="价格id" prop="priceId" width="120" />
        <el-table-column align="left" label="属性" prop="property" width="120" />
        <el-table-column align="left" label="属性id" prop="propertyChildIds" width="120" />
        <el-table-column align="left" label="是否已支付" prop="purchase" width="120" />
        <el-table-column align="left" label="退款状态" prop="refundStatus" width="120" />
        <el-table-column align="left" label="积分" prop="score" width="120" />
        <el-table-column align="left" label="商店id" prop="shopId" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="类型" prop="type" width="120" />

        <el-table-column align="left" label="单位" prop="unit" width="120" />
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
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeOrderGoodsFunc(scope.row)">变更</el-button>
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

            <el-form-item label="订单id:"  prop="orderId" >
              <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入订单id" />
            </el-form-item>
            <el-form-item label="用户id:"  prop="uid" >
              <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入用户id" />
            </el-form-item>
<!--            <el-form-item label="已删除:"  prop="isDeleted" >-->
<!--              <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="添加时间:"  prop="dateAdd" >-->
<!--              <el-date-picker v-model="formData.dateAdd" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="更新时间:"  prop="dateUpdate" >-->
<!--              <el-date-picker v-model="formData.dateUpdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="删除时间:"  prop="dateDelete" >-->
<!--              <el-date-picker v-model="formData.dateDelete" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />-->
<!--            </el-form-item>-->
            <el-form-item label="售后类型:"  prop="afterSale" >
              <el-input v-model="formData.afterSale" :clearable="true"  placeholder="请输入售后类型" />
            </el-form-item>
            <el-form-item label="总金额:"  prop="amount" >
              <el-input-number v-model="formData.amount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="优惠券金额:"  prop="amountCoupon" >
              <el-input-number v-model="formData.amountCoupon"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="单价:"  prop="amountSingle" >
              <el-input-number v-model="formData.amountSingle"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="基本单价:"  prop="amountSingleBase" >
              <el-input-number v-model="formData.amountSingleBase"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="奖励发放结束:"  prop="buyRewardEnd" >
              <el-switch v-model="formData.buyRewardEnd" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="分类:"  prop="categoryId" >
              <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入分类" />
            </el-form-item>
            <el-form-item label="餐饮桌子状态:"  prop="cyTableStatus" >
              <el-input v-model.number="formData.cyTableStatus" :clearable="true" placeholder="请输入餐饮桌子状态" >
                <el-select v-model="formData.cyTableStatus" clearable placeholder="请选择" :clearable="false">
                  <el-option v-for="(item,key) in beeCyTableStatus" :key="key" :label="item.label" :value="parseInt(item.value)" />
                </el-select>
              </el-input>
            </el-form-item>
<!--            <el-form-item label="返现类型:"  prop="fxType" >-->
<!--              <el-input v-model.number="formData.fxType" :clearable="true" placeholder="请输入返现类型" />-->
<!--            </el-form-item>-->
            <el-form-item label="商品id:"  prop="goodsId" >
              <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入商品id" />
            </el-form-item>
            <el-form-item label="商品名:"  prop="goodsName" >
              <el-input v-model="formData.goodsName" :clearable="true"  placeholder="请输入商品名" />
            </el-form-item>
            <el-form-item label="是否积分订单:"  prop="isScoreOrder" >
              <el-switch v-model="formData.isScoreOrder" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="数量:"  prop="number" >
              <el-input v-model.number="formData.number" :clearable="true" placeholder="请输入数量" />
            </el-form-item>
            <el-form-item label="未发货数量:"  prop="numberNoFahuo" >
              <el-input v-model.number="formData.numberNoFahuo" :clearable="true" placeholder="请输入未发货数量" />
            </el-form-item>
            <el-form-item label="人数:"  prop="persion" >
              <el-input v-model.number="formData.persion" :clearable="true" placeholder="请输入人数" />
            </el-form-item>
            <el-form-item label="商品图:"  prop="pic" >
              <el-input v-model="formData.pic" :clearable="true"  placeholder="请输入商品图" />
              <SelectImage
                  v-model="formData.pic"
              />
            </el-form-item>
            <el-form-item label="价格id:"  prop="priceId" >
              <el-input v-model.number="formData.priceId" :clearable="true" placeholder="请输入价格id" />
            </el-form-item>
            <el-form-item label="属性:"  prop="property" >
              <el-input v-model="formData.property" :clearable="true"  placeholder="请输入属性" />
            </el-form-item>
            <el-form-item label="属性id:"  prop="propertyChildIds" >
              <el-input v-model="formData.propertyChildIds" :clearable="true"  placeholder="请输入属性id" />
            </el-form-item>
            <el-form-item label="是否已支付:"  prop="purchase" >
              <el-input v-model="formData.purchase" :clearable="true"  placeholder="请输入是否已支付" />
            </el-form-item>
            <el-form-item label="退款状态:"  prop="refundStatus" >
              <el-input v-model.number="formData.refundStatus" :clearable="true" placeholder="请输入退款状态" />
            </el-form-item>
            <el-form-item label="积分:"  prop="score" >
              <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入积分" />
            </el-form-item>
            <el-form-item label="商店id:"  prop="shopId" >
              <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入商店id" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态" />
            </el-form-item>
            <el-form-item label="类型:"  prop="type" >
              <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入类型" />
            </el-form-item>
            <el-form-item label="单位:"  prop="unit" >
              <el-input v-model="formData.unit" :clearable="true"  placeholder="请输入单位" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeOrderGoods,
  deleteBeeOrderGoods,
  deleteBeeOrderGoodsByIds,
  updateBeeOrderGoods,
  findBeeOrderGoods,
  getBeeOrderGoodsList
} from '@/plugin/beeshop/api/beeOrderGoods'

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
import CustomPic from "@/components/customPic/index.vue";
import SelectImage from "@/components/selectImage/selectImage.vue";

defineOptions({
    name: 'BeeOrderGoods'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

const beeCyTableStatus = ref([])
const init = async () => {
  // 获取字典（可能为空）
  beeCyTableStatus.value = await getDictFunc('BeeCyTableStatus')
}
init()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        id: undefined,
        isDeleted: false,
        dateAdd: new Date(),
        dateUpdate: new Date(),
        dateDelete: new Date(),
        afterSale: '',
        amount: 0,
        amountCoupon: 0,
        amountSingle: 0,
        amountSingleBase: 0,
        buyRewardEnd: false,
        categoryId: undefined,
        cyTableStatus: undefined,
        fxType: undefined,
        goodsId: undefined,
        goodsName: '',
        isScoreOrder: false,
        number: undefined,
        numberNoFahuo: undefined,
        orderId: undefined,
        persion: undefined,
        pic: '',
        priceId: undefined,
        property: '',
        propertyChildIds: '',
        purchase: '',
        refundStatus: undefined,
        score: undefined,
        shopId: undefined,
        status: undefined,
        type: undefined,
        uid: undefined,
        unit: '',
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
        dateAdd : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startDateAdd && !searchInfo.value.endDateAdd) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startDateAdd && searchInfo.value.endDateAdd) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startDateAdd && searchInfo.value.endDateAdd && (searchInfo.value.startDateAdd.getTime() === searchInfo.value.endDateAdd.getTime() || searchInfo.value.startDateAdd.getTime() > searchInfo.value.endDateAdd.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        dateUpdate : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startDateUpdate && !searchInfo.value.endDateUpdate) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startDateUpdate && searchInfo.value.endDateUpdate) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startDateUpdate && searchInfo.value.endDateUpdate && (searchInfo.value.startDateUpdate.getTime() === searchInfo.value.endDateUpdate.getTime() || searchInfo.value.startDateUpdate.getTime() > searchInfo.value.endDateUpdate.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            dateAdd: 'date_add',
            dateUpdate: 'date_update',
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
    if (searchInfo.value.buyRewardEnd === ""){
        searchInfo.value.buyRewardEnd=null
    }
    if (searchInfo.value.isScoreOrder === ""){
        searchInfo.value.isScoreOrder=null
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
  const table = await getBeeOrderGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteBeeOrderGoodsFunc(row)
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
      const res = await deleteBeeOrderGoodsByIds({ ids })
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
const updateBeeOrderGoodsFunc = async(row) => {
    const res = await findBeeOrderGoods({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBeeOrderGoodsFunc = async (row) => {
    const res = await deleteBeeOrderGoods({ id: row.id })
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
        afterSale: '',
        amount: 0,
        amountCoupon: 0,
        amountSingle: 0,
        amountSingleBase: 0,
        buyRewardEnd: false,
        categoryId: undefined,
        cyTableStatus: undefined,
        fxType: undefined,
        goodsId: undefined,
        goodsName: '',
        isScoreOrder: false,
        number: undefined,
        numberNoFahuo: undefined,
        orderId: undefined,
        persion: undefined,
        pic: '',
        priceId: undefined,
        property: '',
        propertyChildIds: '',
        purchase: '',
        refundStatus: undefined,
        score: undefined,
        shopId: undefined,
        status: undefined,
        type: undefined,
        uid: undefined,
        unit: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createBeeOrderGoods(formData.value)
                  break
                case 'update':
                  res = await updateBeeOrderGoods(formData.value)
                  break
                default:
                  res = await createBeeOrderGoods(formData.value)
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
