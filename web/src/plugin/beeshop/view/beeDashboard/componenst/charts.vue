<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="">
    <div class="flex items-center justify-between mb-2">
      <div v-if="title" class="font-bold">
        {{ title }}
      </div>
      <slot v-else name="title" />
    </div>
    <div class=" w-full relative">
      <div v-if="type !== 4">
        <div class="mt-4 text-gray-600 text-3xl font-mono">
          <el-statistic :value="data[data.length-1]" />
        </div>
        <div class="mt-2 text-green-600 text-sm font-bold font-mono">
          {{data[data.length-1] >= data[data.length-2] ? '+' : '-'}}{{ ((data[data.length-1]-data[data.length-2])*100/data[data.length-2]).toFixed(2) }}% <el-icon><TopRight /></el-icon>
        </div>
      </div>
      <div class=" absolute top-0 right-2 w-[50%] h-20">
        <charts-people-number v-if="type === 1" :data="data" height="100%" />
      </div>
    </div>
  </div>
</template>

<script setup>

import {reactive, ref} from 'vue'
import chartsPeopleNumber from "./charts-people-numbers.vue"
import chartsContentNumber from "./charts-content-numbers.vue"

defineProps({
  type :{
    type : Number,
    default : 1
  },
  title : {
    type : String,
    default : ""
  },
  data : {
    type : Array,
    default : []
  }
})

</script>

<style scoped lang="scss">

</style>
