<script setup>
import { ref, onMounted } from 'vue'

const services = ref([])
const loading = ref(true)
const error = ref(null)

const fetchServices = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/status')

    if (!response.ok) {
      throw new Error('Error al obtener los servicios')
    }

    services.value = await response.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchServices()
})
</script>

<template>
  <div class="dashboard">
    <header class="header">
      <h1>Uptime Monitor</h1>
      <p>Service health monitoring dashboard</p>
    </header>

    <p v-if="loading">⏳ Loading services...</p>
    <p v-if="error" class="error">❌ {{ error }}</p>

    <section v-if="!loading && !error" class="services">
      <div
        v-for="(service, index) in services"
        :key="index"
        class="service-card"
      >
        <h3>{{ service.name }}</h3>
        <p class="url">{{ service.url }}</p>

        <span
          class="status"
          :class="service.status === 'UP' ? 'up' : 'down'"
        >
          {{ service.status }}
        </span>

        <p v-if="service.responseTime > 0">
          ⏱ {{ service.responseTime }} ms
        </p>
        <p v-else>
          ⏱ No response
        </p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.dashboard {
  padding: 2rem;
  font-family: system-ui, sans-serif;
}

.header {
  margin-bottom: 2rem;
}

.services {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1rem;
}

.service-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 1rem;
  background: #fff;
}

.url {
  font-size: 0.85rem;
  color: #666;
  word-break: break-all;
}

.status {
  display: inline-block;
  margin: 0.5rem 0;
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: bold;
}

.up {
  background: #e6f4ea;
  color: #1e7e34;
}

.down {
  background: #fdecea;
  color: #c82333;
}
</style>