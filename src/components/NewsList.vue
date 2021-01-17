<template>
  <div>
    <WaitingLoader :isLoading="isLoading"/>
    <NewsItem v-for="p in posts" :id="p.id" v-bind:key="p.id" :title="p.title"/>
  </div>
</template>

<script>
import axios from 'axios';
import NewsItem from "@/components/NewsItem";
import WaitingLoader from "@/components/WaitingLoader";

export default {
  name: "NewsList",
  components: {NewsItem, WaitingLoader},
  data() {
    return {
      isLoading: false,
      posts: []
    }
  },
  mounted() {
    this.listPosts();
  },
  methods: {
    listPosts: function () {
      this.isLoading = true;
      axios.get('http://localhost:9090/v1/' + 'events?page=1&limit=3', {}).then(resp => {
        this.posts = resp.data.data;
        this.isLoading = false;
      }).catch(err => {
        console.log(err);
      });
    },
  }
}
</script>
