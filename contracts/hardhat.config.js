require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-ledger");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    hardhat: {
      forking: {
        url: `https://`,
        blockNumber: 22000000,
      },
    },
    mainnet: {
      url: 'https://',
      ledgerAccounts: ['0x145c300340f323c9f88ce88123e818a7cebeb597'],
    }
  },
};
