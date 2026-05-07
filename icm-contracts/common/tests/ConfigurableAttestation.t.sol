pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {IMessageVerifier} from "../ITeleporterMessengerV2.sol";
import {TeleporterICMMessage} from "../ITeleporterMessengerV2.sol";
import {
    ConfigurableAttestation,
    ConfigurableAttestationVerifier
} from "../ConfigurableAttestation.sol";

contract ConfigurableAttestationTest is Test {
    function testVerification() public {
        Accept verifier1 = new Accept();
        Accept verifier2 = new Accept();
        Accept verifier3 = new Accept();
        Reject verifier4 = new Reject();

        TeleporterICMMessage memory message;
        address[] memory v = new address[](3);

        v[0] = address(verifier1);
        v[1] = address(verifier2);
        v[2] = address(verifier4);
        ConfigurableAttestationVerifier confVerifier1 =
            new ConfigurableAttestationVerifier(ConfigurableAttestation.mOfNConfig(2, 3), v);
        assert(confVerifier1.verifyMessage(message));

        v[0] = address(verifier1);
        v[1] = address(verifier4);
        v[2] = address(verifier3);
        ConfigurableAttestationVerifier confVerifier2 =
            new ConfigurableAttestationVerifier(ConfigurableAttestation.mOfNConfig(2, 3), v);
        assert(confVerifier2.verifyMessage(message));

        v[0] = address(verifier4);
        v[1] = address(verifier2);
        v[2] = address(verifier3);
        ConfigurableAttestationVerifier confVerifier3 =
            new ConfigurableAttestationVerifier(ConfigurableAttestation.mOfNConfig(2, 3), v);
        assert(confVerifier3.verifyMessage(message));

        v[0] = address(verifier1);
        v[1] = address(verifier2);
        v[2] = address(verifier3);
        ConfigurableAttestationVerifier confVerifier4 =
            new ConfigurableAttestationVerifier(ConfigurableAttestation.mOfNConfig(2, 3), v);
        assert(confVerifier4.verifyMessage(message));
    }

    function testMalformedConfig() public {
        Accept verifier1 = new Accept();
        Accept verifier2 = new Accept();

        address[] memory v = new address[](2);
        v[0] = address(verifier1);
        v[1] = address(verifier2);

        // there are at most 2^2 = 4 outcomes and thus the config can be at most 2^4 - 1 = 15
        vm.expectRevert("Malformed configuration");
        new ConfigurableAttestationVerifier(16, v);
    }

    function testNonVerification() public {
        Accept verifier1 = new Accept();
        Accept verifier2 = new Accept();
        Reject verifier3 = new Reject();

        TeleporterICMMessage memory message;
        address[] memory v = new address[](3);

        v[0] = address(verifier1);
        v[1] = address(verifier2);
        v[2] = address(verifier3);
        uint8[] memory subsets = new uint8[](2);
        subsets[0] = 5;
        subsets[1] = 7;
        ConfigurableAttestationVerifier confVerifier1 = new ConfigurableAttestationVerifier(
            ConfigurableAttestation.subsetsToConfig(subsets), v
        );
        assert(!confVerifier1.verifyMessage(message));
    }

    function testConfigFunctions() public pure {
        uint256 expected = 232;
        assertEq(expected, ConfigurableAttestation.mOfNConfig(2, 3));
        uint8[] memory subsets = new uint8[](4);
        subsets[0] = 3;
        subsets[1] = 5;
        subsets[2] = 6;
        subsets[3] = 7;
        assertEq(expected, ConfigurableAttestation.subsetsToConfig(subsets));
    }
}

contract Accept is IMessageVerifier {
    function verifyMessage(
        // solhint-disable-next-line no-unused-vars
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        return true;
    }
}

contract Reject is IMessageVerifier {
    function verifyMessage(
        // solhint-disable-next-line no-unused-vars
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        return false;
    }
}
