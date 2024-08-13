<template>
  <div class="space-y-12">
    <div class="mx-auto text-center">
      <Icon name="material-symbols:check-circle-rounded" class="text-green-600 mr-2 size-28" />
      <div class="font-semibold text-3xl">Its all good man</div>
    </div>

    <div class="grid grid-cols-3 gap-6">
      <SiteHeartbeat hostname="deb0fdaa-786f-4526-bdc0-f732c273df2f" />
      <SiteHeartbeat hostname="simplo.cz" />
      <SiteHeartbeat hostname="gateway01.simplo.cz" />
    </div>
  </div>
</template>

<script setup lang="ts">

</script>
