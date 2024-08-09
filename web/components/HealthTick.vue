<template>
  <div class="bg-green-500 rounded-full flex-shrink-0" style="height: 30px; width: 6px" v-if="props.beat.status_code < 400"></div>
  <div class="bg-amber-500 rounded-full flex-shrink-0" style="height: 30px; width: 6px" v-if="props.beat.status_code >= 400  && props.beat.status_code < 500"></div>
  <div class="bg-red-500 rounded-full flex-shrink-0" style="height: 30px; width: 6px" v-if="props.beat.status_code >= 500"></div>
</template>

<script setup lang="ts">

const props = defineProps(['beat'])

</script>