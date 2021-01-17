<template>
  <div class="mt-2 ml-2">
    <div class="col-12">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">{{ title }}</h5>
          <a v-bind:href="viewId" class="btn btn-primary">
            Details
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "NewsItem",
  props: {
    id: String,
    title: String
  },
  data() {
    return {
      viewId: String
    }
  },
  mounted() {
    this.viewId = '/#/view?id=' + this.id;
  }
}
</script>
