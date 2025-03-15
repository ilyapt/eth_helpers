//SPDX-License-Identifier: MIT
pragma solidity >=0.8.17 <0.9.0;

interface ensReverseRegistrar {
    function node(address) external view returns (bytes32);
}
interface ensRegistry {
    function resolver(bytes32) external view returns (ensResolver);
}
interface ensResolver {
    function name(bytes32) external view returns (string memory);
}

contract EnsReverseHelper {
    ensReverseRegistrar constant reverseRegistrator = ensReverseRegistrar(0xa58E81fe9b61B5c3fE2AFD33CF304c454AbFc7Cb);
    ensRegistry constant registry = ensRegistry(0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e);

    function name(address addr) public view returns(string memory result) {
        bytes32 node = reverseRegistrator.node(addr);
        ensResolver resolver = registry.resolver(node);
        if (address(resolver) != address(0x0)){
            result = resolver.name(node);
        }
    }

    function names(address[] calldata addr) public view returns(string[] memory) {
        string[] memory result = new string[](addr.length);
        for(uint i = 0; i < addr.length; i++) {
            bytes32 node = reverseRegistrator.node(addr[i]);
            ensResolver resolver = registry.resolver(node);
            if (address(resolver) != address(0x0)){
                result[i] = resolver.name(node);
            }
        }
        return result;
    }
}
