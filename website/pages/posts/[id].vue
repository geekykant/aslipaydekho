<template>
  <div>
    <h1>post {{ postId }}</h1>
  </div>
</template>

<script setup>
const route = useRoute();
const postId = route.params.id;
</script>

<style>
</style>