<template>
  <main class="bg-gray-50">
    <nav class="p-4 border-b flex align-middle items-center">
      <div class="mr-auto text-xl font-semibold">Is It Running</div>

      <div>
        <Button as="a" href="/add" variant="outline">
          <Icon name="material-symbols:add" class="mr-2" />
          Add monitor
        </Button>
      </div>
    </nav>

    <div class="container mx-auto mt-6">
      <NuxtPage/>
    </div>
  </main>
</template>
<script setup lang="ts">
</script>