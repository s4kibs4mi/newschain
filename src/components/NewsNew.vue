<template>
  <div class="mt-3">
    <div class="form ml-5 mr-5">
      <div class="form-group">
        <input class="form-control" type="text" v-model="newsTitle" placeholder="News Title">
      </div>
      <div class="form-group">
        <textarea placeholder="News Details" v-model="newsDetails"></textarea>
      </div>
      <div class="form-group">
        <button class="btn btn-primary" v-on:click="onSave">Save Key</button>
      </div>
    </div>
  </div>
</template>

<script>
import Blockchain from "@/utils/blockchain";

export default {
  name: "NewsNew",
  data() {
    return {
      newsTitle: '',
      newsDetails: '',
      loggedAccount: undefined,
    }
  },
  mounted() {
    this.loggedAccount = Blockchain.getAccountByPrivateKey(Blockchain.getPrivateKey());
    Blockchain.setDefaultAccount(this.loggedAccount);
  },
  methods: {
    onSave: function () {
      if (this.loggedAccount === undefined) {
        return
      }

      let contract = Blockchain.getContract();
      contract.methods.createAuthor('Sakib', 'None', 'sakib@liveklass.io', 1).call().then(console.log)

      alert('Post saved');
    }
  }
}
</script>
