<template>
  <div>
    <div>
      <WaitingLoader :isLoading="isLoading"/>
      <NewsItem v-for="p in posts" :id="p.id" v-bind:key="p.id" :title="p.title"/>
    </div>
    <div style="text-align: center" class="mt-2">
      <button class="btn btn-primary" v-on:click="onPrev">Prev</button>&nbsp;<button v-on:click="onNext"
                                                                                     class="btn btn-primary">Next
    </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import NewsItem from "@/components/NewsItem";
import WaitingLoader from "@/components/WaitingLoader";
import Blockchain from "@/utils/blockchain";

export default {
  name: "NewsList",
  components: {NewsItem, WaitingLoader},
  data() {
    return {
      isLoading: false,
      posts: [],
      index: 1,
    }
  },
  mounted() {
    this.listPosts();
  },
  methods: {
    listPosts: function () {
      this.isLoading = true;
      axios.get(Blockchain.getAppUrl() + 'events?page=' + this.index, {}).then(resp => {
        this.posts = resp.data.data;
        this.isLoading = false;
      }).catch(err => {
        console.log(err);
      });
    },
    onNext: function () {
      this.isLoading = true;
      this.index = this.index + 1;
      this.listPosts();
    },
    onPrev: function () {
      if (this.index <= 1) {
        return
      }
      this.isLoading = true;
      this.index = this.index - 1;
      this.listPosts();
    }
  }
}
</script>
