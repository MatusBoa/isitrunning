<template>
  <Card>
    <div class="p-6 space-y-4">
      <div class="flex items-center">
        <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-5" />
        {{ props.monitor.url }}
      </div>

      <div class="flex space-x-[4px] flex-nowrap">
        <TooltipProvider v-for="beat in heartbeats" :delay-duration="1">
          <Tooltip>
            <TooltipTrigger><HealthTick :beat="beat" /></TooltipTrigger>
            <TooltipContent>
              {{ beat.created_at }}
              <br>
              {{ beat.response_time }}
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </div>
    </div>
  </Card>
</template>
<script setup lang="ts">
import {Card} from "~/components/ui/card";
import HealthTick from "~/components/HealthTick.vue";

const props = defineProps(['monitor', 'heartbeats'])
const heartbeats = reactive(props.heartbeats)

import pusher from "~/services/pusher";

pusher.subscribe(props.monitor.uuid).bind('heartbeat', (message) => {
  if (heartbeats.length >= 27) {
    heartbeats.shift()
  }
  
  heartbeats.push(message)
});
</script>