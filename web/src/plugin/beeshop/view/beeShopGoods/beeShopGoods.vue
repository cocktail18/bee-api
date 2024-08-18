<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="条码" prop="barCode">
         <el-input v-model="searchInfo.barCode" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分类id" prop="categoryId">
          <el-select
              v-model="searchInfo.categoryId"
              :clearable="true"
          >
            <el-option
                v-for="(item,key) in goodsCategoryList"
                :key="item.id"
                :value="item.id"
                :label="item.name"
            />
          </el-select>

        </el-form-item>
<!--        <el-form-item label="运费模版" prop="logisticsId">-->
<!--            -->
<!--           <el-input v-model.number="searchInfo.logisticsId" placeholder="搜索条件" />-->

<!--        </el-form-item>-->
        <el-form-item label="商品名字" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商店id" prop="shopId">
            
             <el-input v-model.number="searchInfo.shopId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="开始售卖时间" prop="sellBeginTime">
            
            <template #label>
            <span>
              开始售卖时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startSellBeginTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endSellBeginTime ? time.getTime() > searchInfo.endSellBeginTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endSellBeginTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startSellBeginTime ? time.getTime() < searchInfo.startSellBeginTime.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="结束售卖时间" prop="sellEndTime">
            
            <template #label>
            <span>
              结束售卖时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startSellEndTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endSellEndTime ? time.getTime() > searchInfo.endSellEndTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endSellEndTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startSellEndTime ? time.getTime() < searchInfo.startSellEndTime.getTime() : false"></el-date-picker>

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
        <el-table-column align="left" label="商品名字" prop="name" width="120" />
        <el-table-column align="left" label="图片" prop="pic" width="120">
          <template #default="scope">
            <CustomPic
                pic-type="file"
                style="margin-top:8px"
                :pic-src="scope.row.pic"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="条码" prop="barCode" width="120" />
        <el-table-column align="left" label="售后说明" prop="afterSale" width="120" >
          <template #default="scope">
            <template v-for="(item, index) in scope.row.afterSale.split(',')"  :key="index">
              <span v-if="item && item != ''">
               {{ formatByList(item, beeGoodsAfterSale, 'value', 'label') }},
              </span>
            </template>
          </template>
        </el-table-column>
        <el-table-column align="left" label="分类" prop="categoryId" width="120" >
          <template #default="scope">{{ formatByList(scope.row.categoryId, goodsCategoryList, 'id', 'name') }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="积分奖励" prop="commission" width="120" />-->
<!--        <el-table-column align="left" label="积分奖励结算类型" prop="commissionSettleType" width="120" />-->
<!--        <el-table-column align="left" label="积分奖励类型" prop="commissionType" width="120" />-->
<!--        <el-table-column align="left" label="积分奖励用户类型" prop="commissionUserType" width="120" />-->
<!--        <el-table-column align="left" label="有附加" prop="hasAddition" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.hasAddition) }}</template>-->
<!--        </el-table-column>-->
<!--        <el-table-column align="left" label="有团购" prop="hasTourJourney" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.hasTourJourney) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="隐藏" prop="hidden" width="120" >
          <template #default="scope">{{ formatBoolean(scope.row.limitation) }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="允许砍价" prop="kanjia" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.kanjia) }}</template>-->
<!--        </el-table-column>-->
<!--        <el-table-column align="left" label="砍价价格" prop="kanjiaPrice" width="120" />-->
<!--        <el-table-column align="left" label="限制" prop="limitation" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.limitation) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="运费模版" prop="logisticsId" width="120" >
          <template #default="scope">{{ formatByList(scope.row.logisticsId, beeLogisticsList, 'id', 'name') }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="最大使用优惠券数量" prop="maxCoupons" width="120" />-->
<!--        <el-table-column align="left" label="秒杀" prop="miaosha" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.miaosha) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="最低购买数量" prop="minBuyNumber" width="120" />
        <el-table-column align="left" label="最低价格" prop="minPrice" width="120" />
<!--        <el-table-column align="left" label="需要积分数量" prop="minScore" width="120" />-->

        <el-table-column align="left" label="喜欢数量" prop="numberFav" width="120" />
        <el-table-column align="left" label="商品评分" prop="numberGoodReputation" width="120" />
        <el-table-column align="left" label="订单数量" prop="numberOrders" width="120" />
        <el-table-column align="left" label="评分数量" prop="numberReputation" width="120" />
        <el-table-column align="left" label="销售数量" prop="numberSells" width="120" />
        <el-table-column align="left" label="原价" prop="originalPrice" width="120" />
<!--        <el-table-column align="left" label="海外直邮" prop="overseas" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.overseas) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column sortable align="left" label="排序" prop="paixu" width="120" />
<!--        <el-table-column align="left" label="用户" prop="persion" width="120" />-->

<!--        <el-table-column align="left" label="允许拼团" prop="pingtuan" width="120">-->
<!--            <template #default="scope">{{ formatBoolean(scope.row.pingtuan) }}</template>-->
<!--        </el-table-column>-->
<!--        <el-table-column align="left" label="拼团价格" prop="pingtuanPrice" width="120" />-->
        <el-table-column align="left" label="商品属性" prop="propertyIds" width="120" >
          <template #default="scope">
            <template v-for="(item, index) in scope.row.propertyIds.split(',')"  :key="index">
              <span v-if="item && item != ''">
               {{ formatByList(item, goodsPropList, 'id', 'name') }},
              </span>
            </template>
          </template>

        </el-table-column>
<!--        <el-table-column align="left" label="推荐状态" prop="recommendStatus" width="120" />-->
<!--        <el-table-column align="left" label="秒杀最低购买数量" prop="seckillBuyNumber" width="120" />-->
        <el-table-column sortable align="left" label="商店id" prop="shopId" width="120" />
        <el-table-column align="left" label="商品状态" prop="status" width="120" >
          <template #default="scope">{{ formatEnum(scope.row.status, beeGoodsStatus) }}</template>
        </el-table-column>
<!--        <el-table-column align="left" label="库存" prop="stores" width="120" />-->
        <el-table-column align="left" label="自动下架" prop="stores0Unsale" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.stores0Unsale) }}</template>
        </el-table-column>
         <el-table-column align="left" label="开始售卖时间" prop="sellBeginTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.sellBeginTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="结束售卖时间" prop="sellEndTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.sellEndTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="重量，kg" prop="weight" width="120" />
<!--        <el-table-column align="left" label="类型" prop="type" width="120" />-->
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
              <div>
                <el-button type="primary" link class="table-button" @click="gotoSkuPage(scope.row)">查看sku</el-button>
                <el-button type="primary" link class="table-button" @click="gotoAdditionPage(scope.row)">查看配件信息</el-button>
              </div>
              <div>
                <el-button type="primary" link class="table-button" @click="gotoImagePage(scope.row)">查看商品图</el-button>
                <el-button type="primary" link class="table-button" @click="gotoContentPage(scope.row)">查看详情</el-button>
              </div>
             <div>
               <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeShopGoodsFunc(scope.row)">变更</el-button>
               <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
             </div>
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
            <el-form-item label="条码:"  prop="barCode" >
              <el-input v-model="formData.barCode" :clearable="true"  placeholder="请输入条码" />
            </el-form-item>
            <el-form-item label="售后说明:"  prop="afterSale" >
              <bee-mul-select v-model="formData.afterSale" :options="beeGoodsAfterSale" :field="'value'" :label="'label'"  ></bee-mul-select>
            </el-form-item>
            <el-form-item label="分类:"  prop="categoryId" >
              <el-select
                  v-model="formData.categoryId"
                  :clearable="true"
              >
                <el-option
                    v-for="(item,key) in goodsCategoryList"
                    :key="item.id"
                    :value="item.id"
                    :label="item.name"
                />
              </el-select>
            </el-form-item>
<!--            <el-form-item label="积分奖励:"  prop="commission" >-->
<!--              <el-input-number v-model="formData.commission"  style="width:100%" :precision="2" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="积分奖励结算类型:"  prop="commissionSettleType" >-->
<!--              <el-input v-model.number="formData.commissionSettleType" :clearable="true" placeholder="请输入积分奖励结算类型" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="积分奖励类型:"  prop="commissionType" >-->
<!--              <el-input v-model.number="formData.commissionType" :clearable="true" placeholder="请输入积分奖励类型" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="积分奖励用户类型:"  prop="commissionUserType" >-->
<!--              <el-input v-model.number="formData.commissionUserType" :clearable="true" placeholder="请输入积分奖励用户类型" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="有附加:"  prop="hasAddition" >-->
<!--              <el-switch v-model="formData.hasAddition" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="有团购:"  prop="hasTourJourney" >-->
<!--              <el-switch v-model="formData.hasTourJourney" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
            <el-form-item label="隐藏:"  prop="hidden" >
              <el-switch v-model="formData.hidden" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
<!--            <el-form-item label="允许砍价:"  prop="kanjia" >-->
<!--              <el-switch v-model="formData.kanjia" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="砍价价格:"  prop="kanjiaPrice" >-->
<!--              <el-input-number v-model="formData.kanjiaPrice"  style="width:100%" :precision="2" :clearable="true"  />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="限制:"  prop="limitation" >-->
<!--              <el-switch v-model="formData.limitation" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
            <el-form-item label="运费模版:"  prop="logisticsId" >
              <el-select
                  v-model="formData.logisticsId"
                  :clearable="true"
              >
                <el-option
                    v-for="(item,key) in beeLogisticsList"
                    :key="item.id"
                    :value="item.id"
                    :label="item.name"
                />
              </el-select>
            </el-form-item>
<!--            <el-form-item label="最大使用优惠券数量:"  prop="maxCoupons" >-->
<!--              <el-input v-model.number="formData.maxCoupons" :clearable="true" placeholder="请输入最大使用优惠券数量" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="秒杀:"  prop="miaosha" >-->
<!--              <el-switch v-model="formData.miaosha" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
            <el-form-item label="最低购买数量:"  prop="minBuyNumber" >
              <el-input v-model.number="formData.minBuyNumber" :clearable="true" placeholder="请输入最低购买数量" />
            </el-form-item>
            <el-form-item label="最低价格:"  prop="minPrice" >
              <el-input-number v-model="formData.minPrice"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
<!--            <el-form-item label="需要积分数量:"  prop="minScore" >-->
<!--              <el-input-number v-model="formData.minScore"  style="width:100%" :precision="2" :clearable="true"  />-->
<!--            </el-form-item>-->
            <el-form-item label="商品名字:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入商品名字" />
            </el-form-item>
<!--            <el-form-item label="喜欢数量:"  prop="numberFav" >-->
<!--              <el-input v-model.number="formData.numberFav" :clearable="true" placeholder="请输入喜欢数量" />-->
<!--            </el-form-item>-->
            <el-form-item label="商品评分:"  prop="numberGoodReputation" >
              <el-input-number v-model="formData.numberGoodReputation"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
<!--            <el-form-item label="订单数量:"  prop="numberOrders" >-->
<!--              <el-input v-model.number="formData.numberOrders" :clearable="true" placeholder="请输入订单数量" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="评分数量:"  prop="numberReputation" >-->
<!--              <el-input v-model.number="formData.numberReputation" :clearable="true" placeholder="请输入评分数量" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="销售数量:"  prop="numberSells" >-->
<!--              <el-input v-model.number="formData.numberSells" :clearable="true" placeholder="请输入销售数量" />-->
<!--            </el-form-item>-->
            <el-form-item label="原价:"  prop="originalPrice" >
              <el-input-number v-model="formData.originalPrice"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
<!--            <el-form-item label="海外直邮:"  prop="overseas" >-->
<!--              <el-switch v-model="formData.overseas" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
            <el-form-item label="排序:"  prop="paixu" >
              <el-input v-model.number="formData.paixu" :clearable="true" placeholder="请输入排序" />
            </el-form-item>
<!--            <el-form-item label="用户:"  prop="persion" >-->
<!--              <el-input v-model.number="formData.persion" :clearable="true" placeholder="请输入用户" />-->
<!--            </el-form-item>-->
            <el-form-item label="图片:"  prop="pic" >
              <el-input v-model="formData.pic" :clearable="true"  placeholder="请输入图片" />
              <SelectImage
                  v-model="formData.pic"
              />
            </el-form-item>
<!--            <el-form-item label="允许拼团:"  prop="pingtuan" >-->
<!--              <el-switch v-model="formData.pingtuan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
<!--            </el-form-item>-->
<!--            <el-form-item label="拼团价格:"  prop="pingtuanPrice" >-->
<!--              <el-input-number v-model="formData.pingtuanPrice"  style="width:100%" :precision="2" :clearable="true"  />-->
<!--            </el-form-item>-->
            <el-form-item label="商品属性:"  prop="propertyIds" >
              <bee-mul-select v-model="formData.propertyIds" :options="goodsPropList.filter(item=>item.propertyId===0)" :field="'id'" :label="'name'"  ></bee-mul-select>
            </el-form-item>
<!--            <el-form-item label="推荐状态:"  prop="recommendStatus" >-->
<!--              <el-input v-model.number="formData.recommendStatus" :clearable="true" placeholder="请输入推荐状态" />-->
<!--            </el-form-item>-->
<!--            <el-form-item label="秒杀最低购买数量:"  prop="seckillBuyNumber" >-->
<!--              <el-input v-model.number="formData.seckillBuyNumber" :clearable="true" placeholder="请输入秒杀最低购买数量" />-->
<!--            </el-form-item>-->
            <el-form-item label="商店id:"  prop="shopId" >
              <el-select v-model="formData.shopId" clearable placeholder="请选择" :clearable="false">
                <el-option v-for="(item,key) in shopList" :key="key" :label="item.id + (item.isDeleted ? '(已删除)' :'') " :value="parseInt(item.id)" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品状态:"  prop="status" >
              <el-select v-model="formData.status" clearable placeholder="请选择" :clearable="false">
                <el-option v-for="(item,key) in beeGoodsStatus" :key="key" :label="item.label" :value="parseInt(item.value)" />
              </el-select>
            </el-form-item>
<!--            <el-form-item label="库存:"  prop="stores" >-->
<!--              <el-input v-model.number="formData.stores" :clearable="true" placeholder="请输入库存" />-->
<!--            </el-form-item>-->
            <el-form-item label="自动下架:"  prop="stores0Unsale" >
              <el-switch v-model="formData.stores0Unsale" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="开始售卖时间:"  prop="sellBeginTime" >
              <el-date-picker v-model="formData.sellBeginTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="结束售卖时间:"  prop="sellEndTime" >
              <el-date-picker v-model="formData.sellEndTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="重量，kg:"  prop="weight" >
              <el-input-number v-model="formData.weight"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
<!--            <el-form-item label="类型:"  prop="type" >-->
<!--              <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入类型" />-->
<!--            </el-form-item>-->
            <el-form-item label="单位:"  prop="unit" >
              <el-input v-model="formData.unit" :clearable="true"  placeholder="请输入单位" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeShopGoods,
  deleteBeeShopGoods,
  deleteBeeShopGoodsByIds,
  updateBeeShopGoods,
  findBeeShopGoods,
  getBeeShopGoodsList
} from '@/plugin/beeshop/api/beeShopGoods'

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  filterDataSource,
  ReturnArrImg,
  onDownloadFile,
  formatByList, formatEnum
} from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from "@/components/selectImage/selectImage.vue";
import CustomPic from "@/components/customPic/index.vue";
import {getBeeShopGoodsCategoryList} from "@/plugin/beeshop/api/beeShopGoodsCategory";
import {getBeeLogisticsList} from "@/plugin/beeshop/api/beeLogistics";
import {getBeeShopGoodsPropList} from "@/plugin/beeshop/api/beeShopGoodsProp";
import {useRouter} from "vue-router";
import BeeMulSelect from "@/plugin/beeshop/view/components/beeEnumSelect/beeMulSelect.vue";
import {getBeeShopInfoList} from "@/plugin/beeshop/api/beeShopInfo";

defineOptions({
    name: 'BeeShopGoods'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)


const beeGoodsStatus = ref([])
const beeGoodsAfterSale = ref([])
const init = async () => {
  // 获取字典（可能为空）
  beeGoodsStatus.value = await getDictFunc('BeeGoodsStatus')
  beeGoodsAfterSale.value = await getDictFunc('BeeGoodsAfterSale')
}
init()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        id: undefined,
        isDeleted: false,
        dateAdd: new Date(),
        dateUpdate: new Date(),
        dateDelete: undefined,
        barCode: '',
        afterSale: '',
        categoryId: undefined,
        commission: 0,
        commissionSettleType: undefined,
        commissionType: undefined,
        commissionUserType: undefined,
        hasAddition: false,
        hasTourJourney: false,
        hidden: undefined,
        kanjia: false,
        kanjiaPrice: 0,
        limitation: false,
        logisticsId: undefined,
        maxCoupons: undefined,
        miaosha: false,
        minBuyNumber: undefined,
        minPrice: 0,
        minScore: 0,
        name: '',
        numberFav: undefined,
        numberGoodReputation: 0,
        numberOrders: undefined,
        numberReputation: undefined,
        numberSells: undefined,
        originalPrice: 0,
        overseas: false,
        paixu: undefined,
        persion: undefined,
        pic: '',
        pingtuan: false,
        pingtuanPrice: 0,
        propertyIds: '',
        recommendStatus: undefined,
        seckillBuyNumber: undefined,
        shopId: undefined,
        status: undefined,
        stores: undefined,
        stores0Unsale: false,
        sellBeginTime: new Date(),
        sellEndTime: new Date(),
        weight: 0,
        type: undefined,
        unit: '',
        })



// 验证规则
const rule = reactive({
  categoryId : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  name : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  pic : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  propertyIds : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  unit : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
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
        sellBeginTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startSellBeginTime && !searchInfo.value.endSellBeginTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startSellBeginTime && searchInfo.value.endSellBeginTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startSellBeginTime && searchInfo.value.endSellBeginTime && (searchInfo.value.startSellBeginTime.getTime() === searchInfo.value.endSellBeginTime.getTime() || searchInfo.value.startSellBeginTime.getTime() > searchInfo.value.endSellBeginTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        sellEndTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startSellEndTime && !searchInfo.value.endSellEndTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startSellEndTime && searchInfo.value.endSellEndTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startSellEndTime && searchInfo.value.endSellEndTime && (searchInfo.value.startSellEndTime.getTime() === searchInfo.value.endSellEndTime.getTime() || searchInfo.value.startSellEndTime.getTime() > searchInfo.value.endSellEndTime.getTime())) {
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
const searchInfo = ref({sort: 'id', order: 'descending'})
const router = useRouter()
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            id: 'id',
            dateAdd: 'date_add',
            paixu: 'paixu',
            shopId: 'shop_id',
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
    if (searchInfo.value.hasAddition === ""){
        searchInfo.value.hasAddition=null
    }
    if (searchInfo.value.hasTourJourney === ""){
        searchInfo.value.hasTourJourney=null
    }
    if (searchInfo.value.kanjia === ""){
        searchInfo.value.kanjia=null
    }
    if (searchInfo.value.limitation === ""){
        searchInfo.value.limitation=null
    }
    if (searchInfo.value.miaosha === ""){
        searchInfo.value.miaosha=null
    }
    if (searchInfo.value.overseas === ""){
        searchInfo.value.overseas=null
    }
    if (searchInfo.value.pingtuan === ""){
        searchInfo.value.pingtuan=null
    }
    if (searchInfo.value.stores0Unsale === ""){
        searchInfo.value.stores0Unsale=null
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

  const table = await getBeeShopGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}


const shopList = ref([])
const getShopList = async() => {
  const table = await getBeeShopInfoList({ page: 1, pageSize: 10000 })
  if (table.code === 0) {
    shopList.value = table.data.list
  }
}
getShopList()

const goodsCategoryList = ref([])
const getGoodsCategoryList = async() => {
  const table = await getBeeShopGoodsCategoryList({ page: 1, pageSize: 10000 })
  if (table.code === 0) {
    goodsCategoryList.value = table.data.list
  }
}
getGoodsCategoryList()

const goodsPropList = ref([])
const getGoodsPropList = async() => {
  const table = await getBeeShopGoodsPropList({ page: 1, pageSize: 100000 })
  if (table.code === 0) {
    goodsPropList.value = table.data.list
  }
}
getGoodsPropList()

const beeLogisticsList = ref([])
const getBeeLogisticsListFunc = async() => {
  const table = await getBeeLogisticsList({ page: 1, pageSize: 10000})
  if (table.code === 0) {
    beeLogisticsList.value = table.data.list
  }
}
getBeeLogisticsListFunc()

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
            deleteBeeShopGoodsFunc(row)
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
      const res = await deleteBeeShopGoodsByIds({ ids })
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

const gotoSkuPage = async(row) => {
   await router.push('beeShopGoodsSku?goodsId=' + row.id)
}

const gotoImagePage = async(row) => {
   await router.push('beeShopGoodsImages?goodsId=' + row.id)
}

const gotoContentPage = async(row) => {
   await router.push('beeShopGoodsContent?goodsId=' + row.id)
}

const gotoAdditionPage = async(row) => {
   await router.push('beeShopGoodsAddition?goodsId=' + row.id)
}

// 更新行
const updateBeeShopGoodsFunc = async(row) => {
    const res = await findBeeShopGoods({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBeeShopGoodsFunc = async (row) => {
    const res = await deleteBeeShopGoods({ id: row.id })
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
        barCode: '',
        afterSale: '',
        categoryId: undefined,
        commission: 0,
        commissionSettleType: undefined,
        commissionType: undefined,
        commissionUserType: undefined,
        hasAddition: false,
        hasTourJourney: false,
        hidden: undefined,
        kanjia: false,
        kanjiaPrice: 0,
        limitation: false,
        logisticsId: undefined,
        maxCoupons: undefined,
        miaosha: false,
        minBuyNumber: undefined,
        minPrice: 0,
        minScore: 0,
        name: '',
        numberFav: undefined,
        numberGoodReputation: 0,
        numberOrders: undefined,
        numberReputation: undefined,
        numberSells: undefined,
        originalPrice: 0,
        overseas: false,
        paixu: undefined,
        persion: undefined,
        pic: '',
        pingtuan: false,
        pingtuanPrice: 0,
        propertyIds: '',
        recommendStatus: undefined,
        seckillBuyNumber: undefined,
        shopId: undefined,
        status: undefined,
        stores: undefined,
        stores0Unsale: false,
        sellBeginTime: new Date(),
        sellEndTime: new Date(),
        weight: 0,
        type: undefined,
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
                  res = await createBeeShopGoods(formData.value)
                  break
                case 'update':
                  res = await updateBeeShopGoods(formData.value)
                  break
                default:
                  res = await createBeeShopGoods(formData.value)
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
