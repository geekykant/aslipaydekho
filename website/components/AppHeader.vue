<template>
  <nav id="main-nav">
    <n-space justify="space-between">
      <div>
          <span><nuxt-link to="/">Home</nuxt-link></span>
          <span><nuxt-link to="/about">About</nuxt-link></span>
          <span><nuxt-link to="/posts/">All Posts</nuxt-link></span>
          <span><nuxt-link to="/posts/edit">Edit Post</nuxt-link></span>
      </div>
      <div style="vertical-align: middle; display: flex;">
        <span>👋 {{ todaysDate }}</span>
        <user-login-logout-dropdown/>
      </div>
    </n-space>
  </nav>
</template>

<script lang="ts">
let websiteTitle: string = "AsliPayDekho";
useHead({
  titleTemplate: (titleChunk) => {
    return titleChunk ? `${titleChunk} - ${websiteTitle}` : websiteTitle;
  },
});

export default{
  data() {
    return {
      todaysDate: new Date().toLocaleTimeString()
    }
  },
  mounted(){
    setInterval(() => {
      this.todaysDate = new Date().toLocaleTimeString()
    }, 1000)
  }
}

</script>

<style scoped>
  *{
    background-color: white;
  }

  a {
    color: #546e7a
  }

  a:hover {
    color:  #0a8a68;
  }

  .router-link-exact-active {
    color: #12b488;
    text-decoration: underline;
    font-weight: 700;
  }

  nav {
    padding: 10px 50px;
    border-bottom: 2px solid #eeeeee;
  }

  span {
    margin: 0 15px;
    font-size: 14px;
  }
</style>