<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const products = ref([])

// 调用后端API示例
const fetchProducts = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/products')
    products.value = await res.json()
  } catch (err) {
    console.error('获取产品失败:', err)
  }
}
</script>

<template>
  <main>
    <h1>欢迎来到RevoCore</h1>
    <button @click="fetchProducts">加载产品</button>
    
    <div v-if="products.length">
      <div v-for="product in products" :key="product.id">
        {{ product.name }}
      </div>
    </div>
  </main>
</template>