<template>
  <div class="mt-3 ml-5 mr-5">
    <div>
      <div class="row">
        <h2>{{ newsTitle }}</h2>
      </div>
      <label>{{ newsCreatedAt }}</label>
      <hr/>
      <div class="row">
        <vue-simple-markdown :source="newsDetails"/>
      </div>
    </div>
  </div>
</template>

<script>

import Blockchain from "@/utils/blockchain";

export default {
  name: "NewsView",
  components: {},
  data() {
    return {
      newsTitle: '',
      newsDetails: '',
      newsCreatedAt: 'CreatedAt: ',
      account: undefined
    }
  },
  mounted() {
    let postId = this.$route.query.id;

    Blockchain.getWeb3Client().eth.requestAccounts()
        .then(accounts => {
          this.contract = Blockchain.getContract(Blockchain.getWeb3Client());
          this.account = accounts[0];

          this.contract.methods.getPost(postId).call({
            from: this.account,
          }).then(result => {
            this.isLoading = false;

            this.newsTitle = result.title;
            this.newsDetails = result.details;
            let dateTime = new Date(result.created_at * 1000);
            this.newsCreatedAt = dateTime.toString();

            console.log(result);
          }).catch(console.log)
        }).catch(console.log);
  }
}
</script>
