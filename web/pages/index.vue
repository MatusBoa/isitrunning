<template>
  <div class="space-y-12">
    <div class="mx-auto text-center">
      <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-28" />
      <div class="font-semibold text-3xl">Its all good man</div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <SiteHeartbeat v-for="monitor in data" :monitor="monitor.monitor" :heartbeats="monitor.monitor_heartbeats" />
    </div>
  </div>
</template>

<script setup lang="ts">
const { data } = await useFetch('http://localhost:1323/monitors')
</script>
