<template>
  <h1>Header</h1>
</template>

<script setup lang="ts">
let websiteTitle: string = "AsliPayDekho";
useHead({
  titleTemplate: (titleChunk) => {
    return titleChunk ? `${titleChunk} - ${websiteTitle}` : websiteTitle;
  },
});
</script>

<style>
</style>