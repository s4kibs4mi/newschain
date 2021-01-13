<template>
  <div id="container" class="body-bg">
    <WaitingLoader :isLoading="isLoading"/>
    <div v-if="!isAuthorRegistered" class="form ml-2 mr-2 mt-2">
      <div class="form-group">
        <input class="form-control" v-model="name" placeholder="Your name"/>
      </div>
      <div class="form-group">
        <input class="form-control" v-model="title" placeholder="Your title"/>
      </div>
      <div class="form-group">
        <input class="form-control" v-model="email" placeholder="Your email"/>
      </div>
      <div class="form-group">
        <button class="btn btn-primary" v-on:click="onCreateAccount">Create Account</button>
      </div>
    </div>
    <div v-else class="form ml-2 mr-2 mt-2">
      <div class="form-group">
        <input class="form-control" v-model="name" placeholder="Your name"/>
      </div>
      <div class="form-group">
        <input class="form-control" v-model="title" placeholder="Your title"/>
      </div>
      <div class="form-group">
        <input class="form-control" v-model="email" placeholder="Your email"/>
      </div>
      <div class="form-group">
        <button class="btn btn-primary" v-on:click="onUpdateAccount">Update Account</button>
      </div>
    </div>
  </div>
</template>

<script>
import Blockchain from "@/utils/blockchain";
import WaitingLoader from "@/components/WaitingLoader";

export default {
  name: "AuthorLogin",
  components: {WaitingLoader},
  data() {
    return {
      isLoading: false,
      isAuthorRegistered: false,
      contract: undefined,
      account: undefined,

      name: '',
      title: '',
      email: ''
    }
  },
  mounted() {
    this.isLoading = true;
    Blockchain.getWeb3Client().eth.requestAccounts()
        .then(accounts => {
          this.contract = Blockchain.getContract(Blockchain.getWeb3Client());
          this.account = accounts[0];

          this.contract.methods.isAuthorRegistered().call({
            from: this.account.address,
          }).then(result => {
            this.isLoading = false;
            this.isAuthorRegistered = result;
          }).catch(console.log)
        }).catch(console.log);
  },
  methods: {
    onCreateAccount: function () {
      if (!this.isFieldsValid()) {
        alert('All fields are required');
        return
      }

      this.isLoading = true;

      Blockchain.getWeb3Client().eth.getTransactionCount(this.account)
          .then(count => {
            let txPld = {
              "from": this.account,
              "to": Blockchain.getContractId(),
              "value": Blockchain.getWeb3Client().utils.toHex(Blockchain.getWeb3Client().utils.toWei("0.001", "ether")),
              "gas": 21000,
              "nonce": count,
              "data": this.contract.methods.createAuthor(this.name, this.title, this.email, Date.now()).encodeABI(),
            };

            Blockchain.getWeb3Client().eth.sendTransaction(txPld)
                .then(result => {
                  this.isLoading = false;
                  console.log(result);
                  location.reload();
                }).catch(err => {
              console.log(err)
              this.isLoading = false;
            });
          }).catch(err => {
        console.log(err);
        this.isLoading = false;
      });
    },
    onUpdateAccount: function () {

    },
    isFieldsValid: function () {
      return this.name.length > 0 && this.email.length > 0 && this.title.length > 0;
    }
  }
}
</script>
