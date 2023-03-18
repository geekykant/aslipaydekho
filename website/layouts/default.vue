<template>
  <div>
    <AppHeader/>
      <body>
        <slot />
      </body>
    <AppFooter/>
  </div>
</template>

<style>
body {
  background-color: #f8fafc;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
</style>
