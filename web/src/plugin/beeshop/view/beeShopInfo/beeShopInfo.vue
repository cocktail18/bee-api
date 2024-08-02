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
        <el-table-column align="left" label="省id" prop="provinceId" width="120" >
          <template #default="scope">{{ formatStrByTree(scope.row.provinceId, regionTree, "id", "name") }}</template>
        </el-table-column>
        <el-table-column align="left" label="城市id" prop="cityId" width="120">
          <template #default="scope">{{ formatStrByTree(scope.row.cityId, regionTree, "id", "name") }}</template>
        </el-table-column>
        <el-table-column align="left" label="地区id" prop="districtId" width="120" >
          <template #default="scope">{{ formatStrByTree(scope.row.districtId, regionTree, "id", "name") }}</template>
        </el-table-column>
        <el-table-column align="left" label="地址" prop="address" width="120" />
<!--        <el-table-column align="left" label="绑定用户id" prop="bindUid" width="120" />-->
<!--        <el-table-column align="left" label="支付方式" prop="cyTablePayMod" width="120" />-->

        <el-table-column align="left" label="生鲜配送" prop="expressType" width="120" />
        <el-table-column align="left" label="接单需商家确认" prop="goodsNeedCheck" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.goodsNeedCheck) }}</template>
        </el-table-column>
        <el-table-column align="left" label="纬度" prop="latitude" width="120" />
        <el-table-column align="left" label="联系人手机" prop="linkPhone" width="120" />
        <el-table-column align="left" label="经度" prop="longitude" width="120" />
        <el-table-column align="left" label="店名" prop="name" width="120" />
        <el-table-column align="left" label="门店编号" prop="number" width="120" />
        <el-table-column align="left" label="喜欢人数" prop="numberFav" width="120" />
        <el-table-column align="left" label="商品评分" prop="numberGoodReputation" width="120" />
        <el-table-column align="left" label="下单数" prop="numberOrder" width="120" />
        <el-table-column align="left" label="店铺评分" prop="numberReputation" width="120" />
        <el-table-column align="left" label="开启点餐" prop="openScan" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.openScan) }}</template>
        </el-table-column>
        <el-table-column align="left" label="开启外卖" prop="openWaimai" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.openWaimai) }}</template>
        </el-table-column>
        <el-table-column align="left" label="开启自取" prop="openZiqu" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.openZiqu) }}</template>
        </el-table-column>
        <el-table-column align="left" label="开店时间" prop="openingHours" width="120" />
        <el-table-column align="left" label="排序" prop="paixu" width="120" />
<!--        <el-table-column align="left" label="推荐状态" prop="recommendStatus" width="120" />-->
        <el-table-column align="left" label="服务距离" prop="serviceDistance" width="120" />
<!--        <el-table-column align="left" label="状态" prop="status" width="120" />-->
        <el-table-column align="left" label="发票" prop="taxGst" width="120" />
        <el-table-column align="left" label="发票服务" prop="taxService" width="120" />
<!--        <el-table-column align="left" label="店铺类型" prop="type" width="120" />-->
        <el-table-column align="left" label="营业状态(0:启用1:关闭)" prop="workStatus" width="120" />
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
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeShopInfoFunc(scope.row)">变更</el-button>
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

            <el-form-item label="省:"  prop="provinceId" >
              <el-select
                  v-model="formData.provinceId"
              >
                <el-option
                    v-for="(item,key) in regionTree"
                    :key="item.id"
                    :value="''+item.id"
                    :label="item.name"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="城市:"  prop="cityId" >
              <el-select
                  v-model="formData.cityId"
              >
                <template v-if="formData.provinceId!='' && regionTreeByPid[parseInt(formData.provinceId)]">
                <el-option
                    v-for="(item,key) in regionTreeByPid[parseInt(formData.provinceId)].childs"
                    :key="item.id"
                    :value="''+item.id"
                    :label="item.name"
                />
                </template>
              </el-select>
            </el-form-item>
            <el-form-item label="地区:"  prop="districtId" >
              <el-select
                  v-model="formData.districtId"
              >
                <template v-if="formData.cityId!='' && regionTreeByPid[parseInt(formData.cityId)]">
                <el-option
                    v-if="formData.cityId!==''"
                    v-for="(item,key) in regionTreeByPid[parseInt(formData.cityId)].childs"
                    :key="item.id"
                    :value="''+item.id"
                    :label="item.name"
                />
                </template>
              </el-select>
            </el-form-item>
            <el-form-item label="地址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
            </el-form-item>
<!--            <el-form-item label="绑定用户id:"  prop="bindUid" >-->
<!--              <el-input v-model.number="formData.bindUid" :clearable="true" placeholder="请输入绑定用户id" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="支付方式:"  prop="cyTablePayMod" >-->
<!--              <el-input v-model.number="formData.cyTablePayMod" :clearable="true" placeholder="请输入支付方式" />-->
<!--            </el-form-item>-->
            <el-form-item label="生鲜配送:"  prop="expressType" >
              <el-input v-model="formData.expressType" :clearable="true"  placeholder="达达配送填 dada ,其他留空不要填写" />
            </el-form-item>
            <el-form-item label="接单需商家确认:"  prop="goodsNeedCheck" >
              <el-switch v-model="formData.goodsNeedCheck" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="纬度:"  prop="latitude" >
              <el-input-number v-model="formData.latitude"  style="width:100%" :precision="10" :clearable="true"  />
            </el-form-item>
            <el-form-item label="绑定手机:"  prop="linkPhone" >
              <el-input v-model="formData.linkPhone" :clearable="true"  placeholder="请输入绑定手机" />
            </el-form-item>
            <el-form-item label="经度:"  prop="longitude" >
              <el-input-number v-model="formData.longitude"  style="width:100%" :precision="10" :clearable="true"  />
            </el-form-item>
            <el-form-item label="店名:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入店名" />
            </el-form-item>
            <el-form-item label="门店编号:"  prop="number" >
              <el-input v-model="formData.number" :clearable="true"  placeholder="达达、美团配送的时候需要填写" />
            </el-form-item>
            <el-form-item label="喜欢人数:"  prop="numberFav" >
              <el-input v-model.number="formData.numberFav" :clearable="true" placeholder="请输入喜欢人数" />
            </el-form-item>
            <el-form-item label="商品评分:"  prop="numberGoodReputation" >
              <el-input v-model.number="formData.numberGoodReputation" :clearable="true" placeholder="请输入商品评分" />
            </el-form-item>
            <el-form-item label="下单数:"  prop="numberOrder" >
              <el-input v-model.number="formData.numberOrder" :clearable="true" placeholder="请输入下单数" />
            </el-form-item>
            <el-form-item label="店铺评分:"  prop="numberReputation" >
              <el-input v-model.number="formData.numberReputation" :clearable="true" placeholder="请输入店铺评分" />
            </el-form-item>
            <el-form-item label="开启点餐:"  prop="openScan" >
              <el-switch v-model="formData.openScan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="开启外卖:"  prop="openWaimai" >
              <el-switch v-model="formData.openWaimai" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="开启自取:"  prop="openZiqu" >
              <el-switch v-model="formData.openZiqu" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="开店时间:"  prop="openingHours" >
              <el-input v-model="formData.openingHours" :clearable="true"  placeholder="eg: 10:00-23:30" />
            </el-form-item>
            <el-form-item label="排序:"  prop="paixu" >
              <el-input v-model.number="formData.paixu" :clearable="true" placeholder="请输入排序" />
            </el-form-item>
<!--            <el-form-item label="推荐状态:"  prop="recommendStatus" >-->
<!--              <el-input v-model.number="formData.recommendStatus" :clearable="true" placeholder="请输入推荐状态" />-->
<!--            </el-form-item>-->
            <el-form-item label="服务距离:"  prop="serviceDistance" >
              <el-input-number v-model="formData.serviceDistance"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
<!--            <el-form-item label="状态:"  prop="status" >-->
<!--              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态" />-->
<!--            </el-form-item>-->
            <el-form-item label="发票:"  prop="taxGst" >
              <el-input v-model.number="formData.taxGst" :clearable="true" placeholder="请输入发票" />
            </el-form-item>
            <el-form-item label="发票服务:"  prop="taxService" >
              <el-input v-model.number="formData.taxService" :clearable="true" placeholder="请输入发票服务" />
            </el-form-item>
<!--            <el-form-item label="店铺类型:"  prop="type" >-->
<!--              <el-input v-model="formData.type" :clearable="true"  placeholder="请输入店铺类型" />-->
<!--            </el-form-item>-->
            <el-form-item label="营业状态:"  prop="workStatus" >
              <el-input v-model.number="formData.workStatus" :clearable="true" placeholder="0:启用1:关闭" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeShopInfo,
  deleteBeeShopInfo,
  deleteBeeShopInfoByIds,
  findBeeShopInfo,
  getBeeShopInfoList,
  updateBeeShopInfo
} from '@/plugin/beeshop/api/beeShopInfo'

// 全量引入格式化工具 请按需保留
import {buildTreeMap, formatBoolean, formatDate, formatStrByTree} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {reactive, ref} from 'vue'
import {getBeeRegionCityTree} from "@/plugin/beeshop/api/beeRegion";

defineOptions({
    name: 'BeeShopInfo'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)
const regionTree = ref([])
const regionTreeByPid = ref({})

const init = async () => {
  const regionTreeRes = await getBeeRegionCityTree({})
  if (regionTreeRes.code === 0) {
    regionTree.value = regionTreeRes.data.list
    buildTreeMap(regionTree.value, regionTreeByPid.value)
    console.log(regionTreeByPid)
  }
}
init()


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        id: undefined,
        isDeleted: false,
        dateAdd: new Date(),
        dateUpdate: new Date(),
        dateDelete: new Date(),
        address: '',
        bindUid: undefined,
        cityId: '',
        cyTablePayMod: undefined,
        districtId: '',
        expressType: '',
        goodsNeedCheck: false,
        latitude: 0,
        linkPhone: '',
        longitude: 0,
        name: '',
        number: '',
        numberFav: undefined,
        numberGoodReputation: undefined,
        numberOrder: undefined,
        numberReputation: undefined,
        openScan: false,
        openWaimai: false,
        openZiqu: false,
        openingHours: '',
        paixu: undefined,
        provinceId: '',
        recommendStatus: undefined,
        serviceDistance: 0,
        status: 1,
        taxGst: undefined,
        taxService: undefined,
        type: '',
        workStatus: undefined,
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
    if (searchInfo.value.goodsNeedCheck === ""){
        searchInfo.value.goodsNeedCheck=null
    }
    if (searchInfo.value.openScan === ""){
        searchInfo.value.openScan=null
    }
    if (searchInfo.value.openWaimai === ""){
        searchInfo.value.openWaimai=null
    }
    if (searchInfo.value.openZiqu === ""){
        searchInfo.value.openZiqu=null
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
  const table = await getBeeShopInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteBeeShopInfoFunc(row)
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
      const res = await deleteBeeShopInfoByIds({ ids })
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
const updateBeeShopInfoFunc = async(row) => {
    const res = await findBeeShopInfo({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBeeShopInfoFunc = async (row) => {
    const res = await deleteBeeShopInfo({ id: row.id })
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
        address: '',
        bindUid: undefined,
        cityId: '',
        cyTablePayMod: undefined,
        districtId: '',
        expressType: '',
        goodsNeedCheck: false,
        latitude: 0,
        linkPhone: '',
        longitude: 0,
        name: '',
        number: '',
        numberFav: undefined,
        numberGoodReputation: undefined,
        numberOrder: undefined,
        numberReputation: undefined,
        openScan: false,
        openWaimai: false,
        openZiqu: false,
        openingHours: '',
        paixu: undefined,
        provinceId: '',
        recommendStatus: undefined,
        serviceDistance: 0,
        status: 1,
        taxGst: undefined,
        taxService: undefined,
        type: '',
        workStatus: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createBeeShopInfo(formData.value)
                  break
                case 'update':
                  res = await updateBeeShopInfo(formData.value)
                  break
                default:
                  res = await createBeeShopInfo(formData.value)
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
