<template>
  <Card>
    <div class="p-6 space-y-4">
      <div class="flex items-center">
        <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-5" />
        {{ props.hostname }}
      </div>

      <div class="flex space-x-[3px] flex-nowrap">
        <HealthTick v-for="beat in heartbeats" :beat="beat" />
      </div>
    </div>
  </Card>
</template>
<script setup lang="ts">
import {Card} from "~/components/ui/card";
import HealthTick from "~/components/HealthTick.vue";

const props = defineProps(['hostname'])
const heartbeats = reactive([])

import pusher from "~/services/pusher";

pusher.subscribe(props.hostname).bind('heartbeat', (message) => {
  heartbeats.push(message)
});
</script>