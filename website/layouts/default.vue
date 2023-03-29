<template>
  <div>
    <client-only placeholder="loading..">
      <AppHeader/>
        <slot />
      <AppFooter/>
   </client-only>
  </div>
</template>

<style>
body {
  background-color: #f8fafc;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

.content-wrapper{
  margin: 30px 50px
}

</style>
