<template>
  <div>
    <loading :active.sync="isLoading"
             :can-cancel="false"
             :is-full-page="true" :color="colorCode" :loader="loaderStyle"></loading>
  </div>
</template>

<script>
import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';

export default {
  name: "WaitingLoader",
  components: {Loading},
  props: {
    isLoading: {
      default: false
    }
  },
  data() {
    return {
      colorCode: "#3f71f4",
      loaderStyle: "dots"
    }
  }
}
</script>
