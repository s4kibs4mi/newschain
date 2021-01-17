<template>
  <div class="mt-3 body-bg" id="container">
    <WaitingLoader :isLoading="isLoading"/>
    <div class="form ml-5 mr-5">
      <div class="form-group">
        <input class="form-control" type="text" v-model="newsTitle" placeholder="News Title">
      </div>
      <div class="form-group">
        <textarea placeholder="News Details" v-model="newsDetails" id="news-details"></textarea>
      </div>
      <div class="form-group">
        <button class="btn btn-primary" v-on:click="onUpdate">Update</button>
      </div>
    </div>
  </div>
</template>

<script>
import WaitingLoader from "@/components/WaitingLoader";
import Blockchain from "@/utils/blockchain";
import EasyMDE from "easymde";

export default {
  name: "NewsUpdate",
  components: {WaitingLoader},
  data() {
    return {
      isLoading: false,
      newsTitle: '',
      newsDetails: '',
      account: undefined,
      easyMde: undefined,
      contract: undefined,
      postId: undefined
    }
  },
  mounted() {
    this.isLoading = true;
    this.easyMde = new EasyMDE({element: document.getElementById('news-details')});
    this.postId = this.$route.query.id;

    Blockchain.getWeb3Client().eth.requestAccounts()
        .then(accounts => {
          this.contract = Blockchain.getContract(Blockchain.getWeb3Client());
          this.account = accounts[0];

          this.contract.methods.isAuthorRegistered().call({
            from: this.account,
          }).then(result => {
            if (!result) {
              return this.$router.push('profile');
            }

            this.getPost();
          }).catch(console.log)
        }).catch(console.log);
  },
  methods: {
    onUpdate: function () {
      this.newsDetails = this.easyMde.value();

      if (!this.isFieldsValid()) {
        alert('All fields are required');
        return
      }

      this.isLoading = true;

      Blockchain.getWeb3Client().eth.getTransactionCount(this.account)
          .then(count => {
            let now = Math.round((new Date()).getTime() / 1000);

            let txPld = {
              "from": this.account,
              "to": Blockchain.getContractId(),
              "value": Blockchain.getWeb3Client().utils.toHex(Blockchain.getWeb3Client().utils.toWei("0.001", "ether")),
              "gas": 21000,
              "nonce": count,
              "data": this.contract.methods.updatePost(this.postId, this.newsTitle, this.newsDetails, now).encodeABI(),
            };

            Blockchain.getWeb3Client().eth.sendTransaction(txPld)
                .then(result => {
                  this.isLoading = false;
                  console.log(result);
                  this.$router.push('view?id=' + this.postId);
                }).catch(err => {
              console.log(err)
              this.isLoading = false;
            });
          }).catch(err => {
        console.log(err);
        this.isLoading = false;
      });
    },
    isFieldsValid: function () {
      return this.newsTitle.length > 0 && this.newsDetails.length > 0;
    },
    getPost: function () {
      Blockchain.getWeb3Client().eth.requestAccounts()
          .then(accounts => {
            this.contract = Blockchain.getContract(Blockchain.getWeb3Client());
            this.account = accounts[0];

            this.contract.methods.getPost(this.postId).call({
              from: this.account,
            }).then(result => {
              this.isLoading = false;

              this.newsTitle = result.title;
              this.newsDetails = result.details;
              this.easyMde.value(this.newsDetails);
              console.log(result);
            }).catch(console.log)
          }).catch(console.log);
    }
  }
}
</script>
