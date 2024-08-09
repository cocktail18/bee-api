<template>
  <el-select
      v-model="values"
      :clearable="true"
      multiple
  >
    <el-option
        v-for="(item,key) in options"
        :key="key"
        :value="item[field] + ''"
        :label="item[label]"
    />
  </el-select>
</template>

<script setup>

import {ref, watch} from "vue";

defineOptions({
  name: 'BeeMulSelect',
})
const values = ref([])
const modelValue = ref('')

const props = defineProps({
  values: {
    type: Array,
    required: false,
    default: []
  },
  modelValue: {
    type: String,
    required: false,
    default: ''
  },
  label: {
    type: String,
    required: false,
    default: ''
  },
  field: {
    type: String,
    required: false,
    default: ''
  },
  options: {
    type: Array,
    required: false,
    default: []
  }
})

const init = () => {
  if (props.modelValue) {
    const arr = props.modelValue.split(",")
    values.value = arr.filter(Boolean)
  }
}
init()

watch(
    () => values.value, // 监听的数据源
    (newValue, oldValue) => {
      console.log(`values 变化了：${oldValue} -> ${newValue}`);
      modelValue.value = getValueStr()
      emit('update:modelValue', modelValue.value)
    }
);

const emit = defineEmits(['update:modelValue'])


const getValueStr = () => {
  if (values.value && values.value.length > 0) {
    return ","+ values.value.join(",") + ","
  }
  return ""
}
</script>