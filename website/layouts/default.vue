<template>
  <div>
    <n-config-provider :theme="darkTheme">
      <client-only placeholder="loading..">
        <AppHeader/>
          <slot />
        <AppFooter/>
    </client-only>
   </n-config-provider>
  </div>
</template>

<script>
import { lightTheme } from 'naive-ui'
export default defineComponent({
  setup() {
    return {
      lightTheme
    }
  }
})
</script>

<style>
body {
  background-color: #f8fafc;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

.content-wrapper{
  margin: 30px 50px
}

</style>
