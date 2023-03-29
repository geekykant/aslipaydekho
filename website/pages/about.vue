<script setup>
definePageMeta({
  middleware: "auth",
  validate: async (route) => {
    return route.name == "about";
  },
});
</script>

<template>
  <div class="content-wrapper">
    <h1>About AsliPayDekho</h1>
    <div style="overflow: auto">
      <n-timeline horizontal>
        <n-timeline-item content="We are here. Starting." time="2023-03-28"/>
        <n-timeline-item content="Checkpoint" />
        <n-timeline-item content="Checkpoint" />
        <n-timeline-item content="Checkpoint" />
        <n-timeline-item content="Checkpoint" />
        <n-timeline-item content="Checkpoint" />
      </n-timeline>
    </div>
  </div>
</template>

<style scoped></style>