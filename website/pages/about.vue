<script setup>
definePageMeta({
  middleware: "auth",
  validate: async (route) => {
    return route.name == "about";
  },
});
</script>

<template>
  <div>
    <h1>It's Nuxt3!</h1>
    <p>About Page</p>
  </div>
</template>