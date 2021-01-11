const NewsChain = artifacts.require("NewsChain");

module.exports = function (deployer) {
    deployer.deploy(NewsChain);
};
