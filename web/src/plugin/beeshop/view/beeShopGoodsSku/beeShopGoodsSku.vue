<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
               @keyup.enter="onSubmit">
        <el-form-item label="id" prop="id">

          <el-input v-model.number="searchInfo.id" placeholder="搜索条件"/>

        </el-form-item>
        <el-form-item label="商品id" prop="goodsId">

          <el-input v-model.number="searchInfo.goodsId" placeholder="搜索条件"/>

        </el-form-item>
        <el-form-item label="sku编号" prop="code">
          <el-input v-model="searchInfo.code" placeholder="搜索条件"/>

        </el-form-item>
        <el-form-item label="价格" prop="price">

          <el-input v-model.number="searchInfo.startPrice" placeholder="最小值"/>
          —
          <el-input v-model.number="searchInfo.endPrice" placeholder="最大值"/>

        </el-form-item>
        <el-form-item label="库存" prop="stores">

          <el-input v-model.number="searchInfo.startStores" placeholder="最小值"/>
          —
          <el-input v-model.number="searchInfo.endStores" placeholder="最大值"/>

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->

          <el-form-item label="重量" prop="weight">

            <el-input v-model.number="searchInfo.startWeight" placeholder="最小值"/>
            —
            <el-input v-model.number="searchInfo.endWeight" placeholder="最大值"/>

          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开
          </el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">
          删除
        </el-button>
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
        <el-table-column type="selection" width="55"/>

        <el-table-column sortable align="left" label="id字段" prop="id" width="120"/>
        <el-table-column align="left" label="商品id" prop="goodsId" width="120"/>
        <el-table-column align="left" label="sku编号" prop="code" width="120"/>
        <!--        <el-table-column align="left" label="返现类型" prop="fxType" width="120" />-->
        <el-table-column align="left" label="市场价" prop="originalPrice" width="120"/>
        <!--        <el-table-column align="left" label="拼团价" prop="pingtuanPrice" width="120" />-->
        <el-table-column align="left" label="价格" prop="price" width="120"/>
        <el-table-column align="left" label="属性关系" prop="propertyChildIds" width="120">
          <template #default="scope">
            <template v-if="scope.row.propertyChildIds">
              <template v-for="(item, index) in scope.row.propertyChildIds.split(',')" :key="index">
                <div v-if="item && item !== ''">
                  <template v-for="(item2, index2) in item.split(':')" :key="index2">
                    {{ formatByList(item2, goodsPropList, 'id', 'name') }}
                    {{ index2 === 0 ? ':' : '' }}
                  </template>
                </div>

              </template>
            </template>
          </template>
        </el-table-column>
        <!--        <el-table-column align="left" label="需要积分" prop="score" width="120" />-->
        <el-table-column align="left" label="库存" prop="stores" width="120"/>
        <el-table-column align="left" label="重量" prop="weight" width="120"/>
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
            <el-button type="primary" link icon="edit" class="table-button"
                       @click="updateBeeShopGoodsSkuFunc(scope.row)">变更
            </el-button>
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
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段"/>
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
        <el-form-item label="商品id:" prop="goodsId">
          <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入商品id" @change="onGoodsIdChange"/>
        </el-form-item>
        <el-form-item label="sku编号:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入sku编号"/>
        </el-form-item>
        <!--            <el-form-item label="返现类型:"  prop="fxType" >-->
        <!--              <el-input v-model.number="formData.fxType" :clearable="true" placeholder="请输入返现类型" />-->
        <!--            </el-form-item>-->
        <el-form-item label="原价:" prop="originalPrice">
          <el-input-number v-model="formData.originalPrice" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <!--            <el-form-item label="拼团价:"  prop="pingtuanPrice" >-->
        <!--              <el-input-number v-model="formData.pingtuanPrice"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <el-form-item label="价格:" prop="price">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <el-form-item label="属性关系:" prop="propertyChildIds">
          <el-card class="box-card">
            <template v-for="(item,index) in formData.propertyMapList" :key="index">
            <el-row
                style="width: 100%"
            >
              <el-col :span="4">
                {{ formatByList(item.label, goodsPropList, 'id', 'name') }}：
              </el-col>
              <el-col :span="10">
                <el-select
                    v-model="item.value"
                    :clearable="true"
                >
                  <el-option
                      v-for="(item2,key2) in goodsPropMapList[item.label]"
                      :key="item2.id"
                      :value="item2.id"
                      :label="item2.name"
                  />
                </el-select>
              </el-col>
            </el-row>
          </template>
          </el-card>
        </el-form-item>
        <!--            <el-form-item label="需要积分:"  prop="score" >-->
        <!--              <el-input-number v-model="formData.score"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <el-form-item label="库存:" prop="stores">
          <el-input v-model.number="formData.stores" :clearable="true" placeholder="请输入库存"/>
        </el-form-item>
        <el-form-item label="重量:" prop="weight">
          <el-input-number v-model="formData.weight" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeShopGoodsSku,
  deleteBeeShopGoodsSku,
  deleteBeeShopGoodsSkuByIds,
  findBeeShopGoodsSku,
  getBeeShopGoodsSkuList,
  updateBeeShopGoodsSku
} from '@/plugin/beeshop/api/beeShopGoodsSku'

// 全量引入格式化工具 请按需保留
import {buildTreeMap, formatBoolean, formatByList, formatDate} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {reactive, ref} from 'vue'
import {getBeeShopGoodsPropList} from "@/plugin/beeshop/api/beeShopGoodsProp";
import {findBeeShopGoods} from "@/plugin/beeshop/api/beeShopGoods";
import {useRoute} from "vue-router";

defineOptions({
  name: 'BeeShopGoodsSku'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)
const route = useRoute()

const goodsPropList = ref([])
const goodsPropMapList = ref({})
const getGoodsPropList = async () => {
  const table = await getBeeShopGoodsPropList({page: 1, pageSize: 100000})
  if (table.code === 0) {
    goodsPropList.value = table.data.list
    goodsPropList.value.forEach(item => {
      if (!goodsPropMapList.value[item.propertyId]) {
        goodsPropMapList.value[item.propertyId] = []
      }
      goodsPropMapList.value[item.propertyId].push(item)
    })
  }
}
getGoodsPropList()

// 自动化生成的字典（可能为空）以及字段
const goodsData = ref({})
const formData = ref({
  id: undefined,
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
  goodsId: undefined,
  code: '',
  fxType: undefined,
  originalPrice: 0,
  pingtuanPrice: 0,
  price: 0,
  propertyChildIds: '',
  propertyMapList: [],
  score: 0,
  stores: undefined,
  weight: 0,
})


// 验证规则
const rule = reactive({})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
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
// 排序
const sortChange = ({prop, order}) => {
  const sortMap = {
    id: 'id',
  }

  let sort = sortMap[prop]
  if (!sort) {
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
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.isDeleted === "") {
      searchInfo.value.isDeleted = null
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

const init = () => {
  if (route.query.goodsId) {
    searchInfo.value.goodsId = route.query.goodsId
  }
}
init()

// 查询
const getTableData = async () => {
  const table = await getBeeShopGoodsSkuList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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
const setOptions = async () => {
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
    deleteBeeShopGoodsSkuFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
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
    const res = await deleteBeeShopGoodsSkuByIds({ids})
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
const updateBeeShopGoodsSkuFunc = async (row) => {
  const res = await findBeeShopGoodsSku({id: row.id})
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    await fillSkuPropertyMapList()
    dialogFormVisible.value = true
  }
}

const onGoodsIdChange = async (val) => {
  await fillSkuPropertyMapList()
}

const fillSkuPropertyMapList = async () => {
  const goodsRes = await findBeeShopGoods({id: formData.value.goodsId})
  if (goodsRes.code !== 0) return
  goodsData.value = goodsRes.data
  formData.value.propertyMapList = []
  if (goodsRes.data.propertyIds) {
    const propertyIds = goodsRes.data.propertyIds.split(",")
    const propertyChildIds = formData.value.propertyChildIds
    const propMap = {}
    propertyChildIds.split(",").forEach(propertyChildId => {
      if (propertyChildId) {
        const propertyChildIdArr = propertyChildId.split(":")
        const propertyId = propertyChildIdArr[0]
        propMap[propertyId] = propertyChildIdArr[1]
      }
    })
    propertyIds.forEach(propertyId => {
      if (!propertyId) return
      formData.value.propertyMapList.push({
        'label': parseInt(propertyId),
        'value': parseInt(propMap[propertyId]) || 0
      })
    })
    console.log(formData.value.propertyMapList)
  }
}


// 删除行
const deleteBeeShopGoodsSkuFunc = async (row) => {
  const res = await deleteBeeShopGoodsSku({id: row.id})
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
    goodsId: undefined,
    code: '',
    fxType: undefined,
    originalPrice: 0,
    pingtuanPrice: 0,
    price: 0,
    propertyChildIds: '',
    propertyMapList: [],
    score: 0,
    stores: undefined,
    weight: 0,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res

    // 格式化属性
    formData.value.propertyChildIds = ","+ formData.value.propertyMapList.map(item => {
      return `${item.label}:${item.value}`
    }).join(",") + ","
    console.log(formData.value.propertyChildIds)

    switch (type.value) {
      case 'create':
        res = await createBeeShopGoodsSku(formData.value)
        break
      case 'update':
        res = await updateBeeShopGoodsSku(formData.value)
        break
      default:
        res = await createBeeShopGoodsSku(formData.value)
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
.box-card{
  width: 100%;
}
</style>
