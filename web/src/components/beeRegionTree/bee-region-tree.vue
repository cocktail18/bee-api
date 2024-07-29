<template>
  <div class="tree-content">
    <el-scrollbar>
      <el-tree
          ref="regionTreeRef"
          :data="regionTree"
          :default-checked-keys="regionTreeCheckedIds"
          :props="regionTreeProps"
          :default-expanded-keys="regionTreeExpandedKeys"
          highlight-current
          :accordion="onlyOne"
          node-key="onlyId"
          show-checkbox
      >
        <template #default="{ node, data }">
          <div class="flex items-center justify-between w-full pr-1">
            <span>{{ data.name }} </span>
          </div>
        </template>
      </el-tree>
    </el-scrollbar>
  </div>
  <div class="tree-footer">
    <el-button type="primary" @click="onOk">确定</el-button>
  </div>
</template>

<script setup>

import {useUserStore} from "@/pinia";
import {ref} from "vue";
import {getBeeRegionCityTree} from "@/api/bee/beeRegion";

defineOptions({
  name: 'BeeRegionTree',
})

const userStore = useUserStore()
const regionTreeRef = ref(null)
const regionTreeCheckedIds = ref([])
const regionTree = ref([])
const regionTreeExpandedKeys = ref([])

const emit = defineEmits(['on-ok'])

const regionTreeProps = ref({
  children: 'childs',
  label: 'name',
  disabled: 'disabled'
})

const props = defineProps({
  regionTreeCheckedIds: {
    type: Array,
    required: false,
    default: []
  },
  regionTreeExpandedKeys: {
    type: Array,
    required: false,
    default: []
  },
  onlyOne: {
    type: Boolean,
    required: false,
    default: false
  }
})

const init = async () => {
  const regionTreeRes = await getBeeRegionCityTree({})
  if (regionTreeRes.code === 0) {
    const list = regionTreeRes.data.list
    if (props.onlyOne) {
      list.forEach(node => {
        markDisable(node)
      })
    }
    regionTree.value = list
  }
}
init()

const markDisable = (node) => {
  if (node.childs && node.childs.length > 0) {
    node.disabled = true
    node.childs.forEach(item => {
      markDisable(item)
    })
  } else {
    node.disabled = false
  }
}


const markAllChilds = (node, markMap) => {
  if (node.childs && node.childs.length > 0) {
    node.childs.forEach(item => {
      markMap[item.id] = 1
      markAllChilds(item, markMap)
    })
  }
}

const onOk = () => {
  // const res = []
  // const skipIds = {}
  const nodes = []
  regionTreeRef.value.getCheckedNodes().forEach(item => {
    nodes.push(item)
  })
  emit('on-ok', nodes)
  // //根据pid 进行排序
  // nodes.sort((a, b) => {
  //   return a.level - b.level
  // })
  //
  // // 如果某个节点下的所有子节点被选中，则只需要返回该父节点
  // nodes.forEach(item => {
  //   // if (skipIds[item.id]) return
  //   if (item.level !== 3) return // 只要第三级
  //   res.push({id:item.id, name:item.name})
  // })
  // emit('on-ok', res)
}
</script>


<style scoped lang="scss">

</style>
