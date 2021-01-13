<template>
  <div>
    <Header/>
    <RouterView :key="$route.fullPath"></RouterView>
  </div>
</template>

<script>

import Header from "@/components/Header";

export default {
  name: 'App',
  components: {Header},
  mounted() {

  }
}
</script>
