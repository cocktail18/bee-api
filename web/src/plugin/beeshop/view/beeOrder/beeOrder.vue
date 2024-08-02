<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
               @keyup.enter="onSubmit">
        <el-form-item label="订单号" prop="pid">
          <el-input v-model="searchInfo.orderNumber" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="核销码" prop="pid">
          <el-input v-model="searchInfo.hxNumber" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="已删除" prop="isDeleted">
          <el-select v-model="searchInfo.isDeleted" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true">
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false">
            </el-option>
          </el-select>
        </el-form-item>

        <!--        <el-form-item label="删除时间" prop="dateDelete">-->
        <!--            -->
        <!--            <template #label>-->
        <!--            <span>-->
        <!--              删除时间-->
        <!--              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">-->
        <!--                <el-icon><QuestionFilled /></el-icon>-->
        <!--              </el-tooltip>-->
        <!--            </span>-->
        <!--          </template>-->
        <!--            <el-date-picker v-model="searchInfo.startDateDelete" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endDateDelete ? time.getTime() > searchInfo.endDateDelete.getTime() : false"></el-date-picker>-->
        <!--            —-->
        <!--            <el-date-picker v-model="searchInfo.endDateDelete" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startDateDelete ? time.getTime() < searchInfo.startDateDelete.getTime() : false"></el-date-picker>-->

        <!--        </el-form-item>-->

        <el-form-item label="付款金额" prop="amountReal">

          <el-input v-model.number="searchInfo.startAmountReal" placeholder="最小值"/>
          —
          <el-input v-model.number="searchInfo.endAmountReal" placeholder="最大值"/>

        </el-form-item>
        <!--        <el-form-item label="自动发货状态" prop="autoDeliverStatus">-->
        <!--            -->
        <!--             <el-input v-model.number="searchInfo.autoDeliverStatus" placeholder="搜索条件" />-->

        <!--        </el-form-item>-->

        <el-form-item label="支付订单时间" prop="datePay">

          <template #label>
            <span>
              支付订单时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled/></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startDatePay" type="datetime" placeholder="开始日期"
                          :disabled-date="time=> searchInfo.endDatePay ? time.getTime() > searchInfo.endDatePay.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endDatePay" type="datetime" placeholder="结束日期"
                          :disabled-date="time=> searchInfo.startDatePay ? time.getTime() < searchInfo.startDatePay.getTime() : false"></el-date-picker>

        </el-form-item>

        <el-form-item label="订单状态" prop="status">

          <el-select v-model="searchInfo.status" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeOrderStatus" :key="key" :label="item.label"
                       :value="parseInt(item.value)"/>
          </el-select>

        </el-form-item>
        <el-form-item label="小费金额" prop="trips">

          <el-input v-model.number="searchInfo.startTrips" placeholder="最小值"/>
          —
          <el-input v-model.number="searchInfo.endTrips" placeholder="最大值"/>

        </el-form-item>
        <el-form-item label="订单类型" prop="type">

          <el-select v-model="searchInfo.type" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeOrderStatus" :key="key" :label="item.label"
                       :value="parseInt(item.value)"/>
          </el-select>
        </el-form-item>
        <el-form-item label="用户id" prop="uid">

          <el-input v-model.number="searchInfo.uid" placeholder="搜索条件"/>

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="运费" prop="amountLogistics">

            <el-input v-model.number="searchInfo.startAmountLogistics" placeholder="最小值"/>
            —
            <el-input v-model.number="searchInfo.endAmountLogistics" placeholder="最大值"/>

          </el-form-item>
          <el-form-item label="关闭订单时间" prop="dateClose">

            <template #label>
            <span>
              关闭订单时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled/></el-icon>
              </el-tooltip>
            </span>
            </template>
            <el-date-picker v-model="searchInfo.startDateClose" type="datetime" placeholder="开始日期"
                            :disabled-date="time=> searchInfo.endDateClose ? time.getTime() > searchInfo.endDateClose.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDateClose" type="datetime" placeholder="结束日期"
                            :disabled-date="time=> searchInfo.startDateClose ? time.getTime() < searchInfo.startDateClose.getTime() : false"></el-date-picker>

          </el-form-item>
          <el-form-item label="是否已退款" prop="hasRefund">
            <el-select v-model="searchInfo.hasRefund" clearable placeholder="请选择">
              <el-option
                  key="true"
                  label="是"
                  value="true">
              </el-option>
              <el-option
                  key="false"
                  label="否"
                  value="false">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="订单已经结束" prop="isEnd">
            <el-select v-model="searchInfo.isEnd" clearable placeholder="请选择">
              <el-option
                  key="true"
                  label="是"
                  value="true">
              </el-option>
              <el-option
                  key="false"
                  label="否"
                  value="false">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="父订单号" prop="pid">

            <el-input v-model.number="searchInfo.pid" placeholder="搜索条件"/>

          </el-form-item>
          <el-form-item label="创建时间" prop="dateAdd">

            <template #label>
            <span>
              创建时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled/></el-icon>
              </el-tooltip>
            </span>
            </template>
            <el-date-picker v-model="searchInfo.startDateAdd" type="datetime" placeholder="开始日期"
                            :disabled-date="time=> searchInfo.endDateAdd ? time.getTime() > searchInfo.endDateAdd.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDateAdd" type="datetime" placeholder="结束日期"
                            :disabled-date="time=> searchInfo.startDateAdd ? time.getTime() < searchInfo.startDateAdd.getTime() : false"></el-date-picker>

          </el-form-item>
          <el-form-item label="更新时间" prop="dateUpdate">

            <template #label>
            <span>
              更新时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled/></el-icon>
              </el-tooltip>
            </span>
            </template>
            <el-date-picker v-model="searchInfo.startDateUpdate" type="datetime" placeholder="开始日期"
                            :disabled-date="time=> searchInfo.endDateUpdate ? time.getTime() > searchInfo.endDateUpdate.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDateUpdate" type="datetime" placeholder="结束日期"
                            :disabled-date="time=> searchInfo.startDateUpdate ? time.getTime() < searchInfo.startDateUpdate.getTime() : false"></el-date-picker>

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
        <el-table-column align="left" label="父订单号" prop="pid" width="120"/>
        <el-table-column align="left" label="订单号" prop="orderNumber" width="120"/>
        <el-table-column align="left" label="用户id" prop="uid" width="120"/>
        <el-table-column align="left" label="商品金额" prop="amount" width="120"/>
        <!--        <el-table-column align="left" label="优惠卡抵扣" prop="amountCard" width="120" />-->
        <el-table-column align="left" label="优惠券抵扣" prop="amountCoupons" width="120"/>
        <el-table-column align="left" label="运费" prop="amountLogistics" width="120"/>
        <el-table-column align="left" label="付款金额" prop="amountReal" width="120"/>
        <el-table-column align="left" label="退款总金额" prop="amountRefundTotal" width="120"/>
        <!--        <el-table-column align="left" label="总税金额" prop="amountTax" width="120" />-->
        <!--        <el-table-column align="left" label="消费税" prop="amountTaxGst" width="120" />-->
        <!--        <el-table-column align="left" label="税服务费用" prop="amountTaxService" width="120" />-->
        <!--        <el-table-column align="left" label="自动发货状态" prop="autoDeliverStatus" width="120" />-->
        <el-table-column align="left" label="关闭订单时间" prop="dateClose" width="180">
          <template #default="scope">{{ formatDate(scope.row.dateClose) }}</template>
        </el-table-column>
        <el-table-column align="left" label="支付订单时间" prop="datePay" width="180">
          <template #default="scope">{{ formatDate(scope.row.datePay) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商品总数量" prop="goodsNumber" width="120"/>
        <el-table-column align="left" label="是否退款" prop="hasRefund" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.hasRefund) }}</template>
        </el-table-column>
        <el-table-column align="left" label="核销码" prop="hxNumber" width="120"/>
        <el-table-column align="left" label="下单ip" prop="ip" width="120"/>
        <el-table-column align="left" label="是否可以核销" prop="isCanHx" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isCanHx) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户删除" prop="isDelUser" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isDelUser) }}</template>
        </el-table-column>
        <el-table-column align="left" label="订单已经结束" prop="isEnd" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isEnd) }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="是否有收益" prop="isHasBenefit" width="120">-->
        <!--            <template #default="scope">{{ formatBoolean(scope.row.isHasBenefit) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="需要配送" prop="isNeedLogistics" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isNeedLogistics) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否已经支付" prop="isPay" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isPay) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否有积分" prop="isScoreOrder" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isScoreOrder) }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="是否拼团成功" prop="isSuccessPingtuan" width="120">-->
        <!--            <template #default="scope">{{ formatBoolean(scope.row.isSuccessPingtuan) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="订单类型" prop="orderType" width="120"/>
        <el-table-column align="left" label="取单号" prop="qudanhao" width="120"/>
        <el-table-column align="left" label="退款状态" prop="refundStatus" width="120"/>
        <el-table-column align="left" label="备注" prop="remark" width="120"/>
        <el-table-column align="left" label="获得积分" prop="score" width="120"/>
        <el-table-column align="left" label="扣除积分" prop="scoreDeduction" width="120"/>
        <el-table-column align="left" label="商店id" prop="shopId" width="120"/>
        <el-table-column align="left" label="自提商店id" prop="shopIdZt" width="120"/>
        <el-table-column align="left" label="自提商店名称" prop="shopNameZt" width="120"/>
        <el-table-column align="left" label="订单状态" prop="status" width="120">
          <template #default="scope">{{ formatEnum(scope.row.status, beeOrderStatus) }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="小费金额" prop="trips" width="120" />-->
        <el-table-column align="left" label="订单类型" prop="type" width="120">
          <template #default="scope">{{ formatEnum(scope.row.type, beeOrderType) }}</template>
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
            <div>
              <el-button type="primary" link class="table-button" @click="gotoDetailPage(scope.row)">查看详情</el-button>
            </div>
            <div>
              <el-button type="primary" link icon="edit" class="table-button" @click="updateBeeOrderFunc(scope.row)">
                变更
              </el-button>
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
        <el-form-item label="父订单号:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入父订单号"/>
        </el-form-item>
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入用户id"/>
        </el-form-item>
        <el-form-item label="取单号:" prop="qudanhao">
          <el-input v-model="formData.qudanhao" :clearable="true" placeholder="请输入取单号"/>
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
        <el-form-item label="商品金额:" prop="amount">
          <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <!--            <el-form-item label="优惠卡抵扣:"  prop="amountCard" >-->
        <!--              <el-input-number v-model="formData.amountCard"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <el-form-item label="优惠券抵扣:" prop="amountCoupons">
          <el-input-number v-model="formData.amountCoupons" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <el-form-item label="运费:" prop="amountLogistics">
          <el-input-number v-model="formData.amountLogistics" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <el-form-item label="付款金额:" prop="amountReal">
          <el-input-number v-model="formData.amountReal" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <el-form-item label="退款总金额:" prop="amountRefundTotal">
          <el-input-number v-model="formData.amountRefundTotal" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <!--            <el-form-item label="总税金额:"  prop="amountTax" >-->
        <!--              <el-input-number v-model="formData.amountTax"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <!--            <el-form-item label="消费税:"  prop="amountTaxGst" >-->
        <!--              <el-input-number v-model="formData.amountTaxGst"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <!--            <el-form-item label="税服务费用:"  prop="amountTaxService" >-->
        <!--              <el-input-number v-model="formData.amountTaxService"  style="width:100%" :precision="2" :clearable="true"  />-->
        <!--            </el-form-item>-->
        <!--            <el-form-item label="自动发货状态:"  prop="autoDeliverStatus" >-->
        <!--              <el-input v-model.number="formData.autoDeliverStatus" :clearable="true" placeholder="请输入自动发货状态" />-->
        <!--            </el-form-item>-->
        <el-form-item label="关闭订单时间:" prop="dateClose">
          <el-date-picker v-model="formData.dateClose" type="date" style="width:100%" placeholder="选择日期"
                          :clearable="true"/>
        </el-form-item>
        <el-form-item label="支付订单时间:" prop="datePay">
          <el-date-picker v-model="formData.datePay" type="date" style="width:100%" placeholder="选择日期"
                          :clearable="true"/>
        </el-form-item>
        <el-form-item label="商品总数量:" prop="goodsNumber">
          <el-input v-model.number="formData.goodsNumber" :clearable="true" placeholder="请输入商品总数量"/>
        </el-form-item>
        <el-form-item label="是否已退款:" prop="hasRefund">
          <el-switch v-model="formData.hasRefund" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="核销码:" prop="hxNumber">
          <el-input v-model="formData.hxNumber" :clearable="true" placeholder="请输入核销码"/>
        </el-form-item>
        <el-form-item label="下单ip:" prop="ip">
          <el-input v-model="formData.ip" :clearable="true" placeholder="请输入下单ip"/>
        </el-form-item>
        <el-form-item label="是否可以核销:" prop="isCanHx">
          <el-switch v-model="formData.isCanHx" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="用户删除:" prop="isDelUser">
          <el-switch v-model="formData.isDelUser" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="订单已经结束:" prop="isEnd">
          <el-switch v-model="formData.isEnd" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <!--            <el-form-item label="是否有收益:"  prop="isHasBenefit" >-->
        <!--              <el-switch v-model="formData.isHasBenefit" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
        <!--            </el-form-item>-->
        <el-form-item label="需要配送:" prop="isNeedLogistics">
          <el-switch v-model="formData.isNeedLogistics" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="是否已经支付:" prop="isPay">
          <el-switch v-model="formData.isPay" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="是否有积分:" prop="isScoreOrder">
          <el-switch v-model="formData.isScoreOrder" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <!--            <el-form-item label="是否拼团成功:"  prop="isSuccessPingtuan" >-->
        <!--              <el-switch v-model="formData.isSuccessPingtuan" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>-->
        <!--            </el-form-item>-->
        <el-form-item label="订单号:" prop="orderNumber">
          <el-input v-model="formData.orderNumber" :clearable="true" placeholder="请输入订单号"/>
        </el-form-item>
        <el-form-item label="订单类型:" prop="orderType">
          <el-input v-model.number="formData.orderType" :clearable="true" placeholder="请输入订单类型"/>
        </el-form-item>
        <el-form-item label="退款状态:" prop="refundStatus">
          <el-input v-model.number="formData.refundStatus" :clearable="true" placeholder="请输入退款状态"/>
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注"/>
        </el-form-item>
        <el-form-item label="获得积分:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入获得积分"/>
        </el-form-item>
        <el-form-item label="扣除积分:" prop="scoreDeduction">
          <el-input v-model.number="formData.scoreDeduction" :clearable="true" placeholder="请输入扣除积分"/>
        </el-form-item>
        <el-form-item label="商店id:" prop="shopId">
          <el-input v-model.number="formData.shopId" :clearable="true" placeholder="请输入商店id"/>
        </el-form-item>
        <el-form-item label="自提商店id:" prop="shopIdZt">
          <el-input v-model.number="formData.shopIdZt" :clearable="true" placeholder="请输入自提商店id"/>
        </el-form-item>
        <el-form-item label="自提商店名称:" prop="shopNameZt">
          <el-input v-model="formData.shopNameZt" :clearable="true" placeholder="请输入自提商店名称"/>
        </el-form-item>
        <el-form-item label="订单状态:" prop="status">
          <el-select v-model="formData.status" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeOrderStatus" :key="key" :label="item.label"
                       :value="parseInt(item.value)"/>
          </el-select>
        </el-form-item>
        <!--            <el-form-item label="小费金额:"  prop="trips" >-->
        <!--              <el-input v-model.number="formData.trips" :clearable="true" placeholder="请输入小费金额" />-->
        <!--            </el-form-item>-->
        <el-form-item label="订单类型:" prop="type">
          <el-select v-model="formData.type" clearable placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in beeOrderType" :key="key" :label="item.label" :value="parseInt(item.value)"/>
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createBeeOrder,
  deleteBeeOrder,
  deleteBeeOrderByIds,
  updateBeeOrder,
  findBeeOrder,
  getBeeOrderList
} from '@/plugin/beeshop/api/beeOrder'

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
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, reactive} from 'vue'
import {useRouter} from "vue-router";

defineOptions({
  name: 'BeeOrder'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

const router = useRouter()
const beeOrderStatus = ref([])
const beeOrderType = ref([])
const init = async () => {
  beeOrderStatus.value = await getDictFunc('OrderStatus')
  beeOrderType.value = await getDictFunc('OrderType')
}
init()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  id: undefined,
  isDeleted: false,
  dateAdd: new Date(),
  dateUpdate: new Date(),
  dateDelete: new Date(),
  amount: 0,
  amountCard: 0,
  amountCoupons: 0,
  amountLogistics: 0,
  amountReal: 0,
  amountRefundTotal: 0,
  amountTax: 0,
  amountTaxGst: 0,
  amountTaxService: 0,
  autoDeliverStatus: undefined,
  dateClose: new Date(),
  datePay: new Date(),
  goodsNumber: undefined,
  hasRefund: false,
  hxNumber: '',
  ip: '',
  isCanHx: false,
  isDelUser: false,
  isEnd: false,
  isHasBenefit: false,
  isNeedLogistics: false,
  isPay: false,
  isScoreOrder: false,
  isSuccessPingtuan: false,
  orderNumber: '',
  orderType: undefined,
  pid: undefined,
  qudanhao: '',
  refundStatus: undefined,
  remark: '',
  score: undefined,
  scoreDeduction: undefined,
  shopId: undefined,
  shopIdZt: undefined,
  shopNameZt: '',
  status: undefined,
  trips: undefined,
  type: undefined,
  uid: undefined,
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
  dateAdd: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startDateAdd && !searchInfo.value.endDateAdd) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startDateAdd && searchInfo.value.endDateAdd) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startDateAdd && searchInfo.value.endDateAdd && (searchInfo.value.startDateAdd.getTime() === searchInfo.value.endDateAdd.getTime() || searchInfo.value.startDateAdd.getTime() > searchInfo.value.endDateAdd.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
  dateUpdate: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startDateUpdate && !searchInfo.value.endDateUpdate) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startDateUpdate && searchInfo.value.endDateUpdate) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startDateUpdate && searchInfo.value.endDateUpdate && (searchInfo.value.startDateUpdate.getTime() === searchInfo.value.endDateUpdate.getTime() || searchInfo.value.startDateUpdate.getTime() > searchInfo.value.endDateUpdate.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
  dateDelete: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startDateDelete && !searchInfo.value.endDateDelete) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startDateDelete && searchInfo.value.endDateDelete) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startDateDelete && searchInfo.value.endDateDelete && (searchInfo.value.startDateDelete.getTime() === searchInfo.value.endDateDelete.getTime() || searchInfo.value.startDateDelete.getTime() > searchInfo.value.endDateDelete.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
  dateClose: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startDateClose && !searchInfo.value.endDateClose) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startDateClose && searchInfo.value.endDateClose) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startDateClose && searchInfo.value.endDateClose && (searchInfo.value.startDateClose.getTime() === searchInfo.value.endDateClose.getTime() || searchInfo.value.startDateClose.getTime() > searchInfo.value.endDateClose.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
  datePay: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startDatePay && !searchInfo.value.endDatePay) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startDatePay && searchInfo.value.endDatePay) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startDatePay && searchInfo.value.endDatePay && (searchInfo.value.startDatePay.getTime() === searchInfo.value.endDatePay.getTime() || searchInfo.value.startDatePay.getTime() > searchInfo.value.endDatePay.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
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
    dateAdd: 'date_add',
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
    if (searchInfo.value.hasRefund === "") {
      searchInfo.value.hasRefund = null
    }
    if (searchInfo.value.isCanHx === "") {
      searchInfo.value.isCanHx = null
    }
    if (searchInfo.value.isDelUser === "") {
      searchInfo.value.isDelUser = null
    }
    if (searchInfo.value.isEnd === "") {
      searchInfo.value.isEnd = null
    }
    if (searchInfo.value.isHasBenefit === "") {
      searchInfo.value.isHasBenefit = null
    }
    if (searchInfo.value.isNeedLogistics === "") {
      searchInfo.value.isNeedLogistics = null
    }
    if (searchInfo.value.isPay === "") {
      searchInfo.value.isPay = null
    }
    if (searchInfo.value.isScoreOrder === "") {
      searchInfo.value.isScoreOrder = null
    }
    if (searchInfo.value.isSuccessPingtuan === "") {
      searchInfo.value.isSuccessPingtuan = null
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
const getTableData = async () => {
  const table = await getBeeOrderList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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
    deleteBeeOrderFunc(row)
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
    const res = await deleteBeeOrderByIds({ids})
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
const updateBeeOrderFunc = async (row) => {
  const res = await findBeeOrder({id: row.id})
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteBeeOrderFunc = async (row) => {
  const res = await deleteBeeOrder({id: row.id})
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
    amount: 0,
    amountCard: 0,
    amountCoupons: 0,
    amountLogistics: 0,
    amountReal: 0,
    amountRefundTotal: 0,
    amountTax: 0,
    amountTaxGst: 0,
    amountTaxService: 0,
    autoDeliverStatus: undefined,
    dateClose: new Date(),
    datePay: new Date(),
    goodsNumber: undefined,
    hasRefund: false,
    hxNumber: '',
    ip: '',
    isCanHx: false,
    isDelUser: false,
    isEnd: false,
    isHasBenefit: false,
    isNeedLogistics: false,
    isPay: false,
    isScoreOrder: false,
    isSuccessPingtuan: false,
    orderNumber: '',
    orderType: undefined,
    pid: undefined,
    qudanhao: '',
    refundStatus: undefined,
    remark: '',
    score: undefined,
    scoreDeduction: undefined,
    shopId: undefined,
    shopIdZt: undefined,
    shopNameZt: '',
    status: undefined,
    trips: undefined,
    type: undefined,
    uid: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBeeOrder(formData.value)
        break
      case 'update':
        res = await updateBeeOrder(formData.value)
        break
      default:
        res = await createBeeOrder(formData.value)
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

const gotoDetailPage = async(row) => {
  await router.push('beeOrderDetail?id='+row.id)
}
</script>

<style>

</style>
