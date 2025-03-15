const { expect } = require('chai');

describe('EnsReverseHelper', function () {
    let helper;

    before(async function () {
        helper = await ethers.getContractFactory('EnsReverseHelper')
          .then(helper => helper.deploy());
    });

    it('single address query', async function () {
        expect(await helper.name('0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13'))
          .to.equal('jaredfromsubway.eth');
    });

    it('multi address query', async function () {
        expect(await helper.names([
          '0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13',
          '0x1f2F10D1C40777AE1Da742455c65828FF36Df387',
          '0x623ac7E26C774Dbc646e0d4aC12110567CEf579F',
          '0x225f137127d9067788314bc7fcc1f36746a3c3B5',
        ])).to.deep.equal(['jaredfromsubway.eth','','','luc.eth']);
    });
});
