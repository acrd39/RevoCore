<script setup>
import { ref } from 'vue'

const products = ref([])

const fetchProducts = async () => {
  try {
    const res = await uni.request({
      url: 'http://localhost:8080/api/products'
    })
    products.value = res.data
  } catch (err) {
    uni.showToast({ title: '加载失败', icon: 'error' })
  }
}
</script>

<template>
  <view class="container">
    <button @tap="fetchProducts">加载产品</button>
    <scroll-view scroll-y>
      <view v-for="item in products" :key="item.id">
        {{ item.name }}
      </view>
    </scroll-view>
  </view>
</template>