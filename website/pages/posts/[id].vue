<template>
  <h1>post {{ postId }}</h1>
</template>

<script setup>
const route = useRoute();
const postId = route.params.id;
</script>

<style>
</style>