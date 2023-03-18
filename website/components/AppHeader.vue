<template>
    <a-layout-header id="main-header">
      <nav id="main-nav">
        <ul>
          <li><NuxtLink to="/">Home</NuxtLink></li>
          <li><NuxtLink to="/about">About</NuxtLink></li>
          <li><NuxtLink to="/posts/1">Post1</NuxtLink></li>
          <li><NuxtLink to="/posts/2">Post2</NuxtLink></li>
        </ul>
      </nav>
      <div>
      </div>
    </a-layout-header>
</template>

<script setup lang="ts">
let websiteTitle: string = "AsliPayDekho";
useHead({
  titleTemplate: (titleChunk) => {
    return titleChunk ? `${titleChunk} - ${websiteTitle}` : websiteTitle;
  },
});
</script>

<style scoped>
  a {
    color: rgb(0,0,0,0.7);
  }

  a:hover {
    color:  #0a8a68;
  }

  .router-link-exact-active {
    color: #12b488;
    text-decoration: underline;
    font-weight: 700;
  }

 nav > ul {
    display: flex;
    justify-content: start;
    list-style-type: none
  }

  li {
    margin: 0 15px;
    font-size: 15px;
  }

  header {
    background: white;
    box-shadow: 0px 2px 10px 1px rgb(0 0 0 / 10%);
  }

  .ant-layout-header{
    height: 54px;
    line-height: 54px;
  }
</style>