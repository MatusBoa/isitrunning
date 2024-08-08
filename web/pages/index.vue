<template>
  <div class="space-y-12">
    <div class="mx-auto text-center">
      <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-28" />
      <div class="font-semibold text-3xl">Its all good man</div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <div v-for="n in 6">
        <Card>
          <div class="p-6 space-y-4">
            <div class="flex items-center">
              <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-5" />
              koterba.sk
            </div>

            <div class="flex space-x-[3px] flex-nowrap">
              <HealthTick v-for="n in 42" />
            </div>
          </div>
        </Card>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {Card, CardContent, CardHeader, CardTitle} from "~/components/ui/card";
import HealthTick from "~/components/HealthTick.vue";
import pusher from "~/services/pusher";

pusher.subscribe('test').bind('stats', (message) => {
  console.log(message)
});
</script>