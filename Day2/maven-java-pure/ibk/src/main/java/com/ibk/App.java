package com.ibk;

/**
 * Hello world!
 *
 */
public class App {
    public static void main(String[] args) {
        //1. 필요한 변수 설정
        HFCAClient hfcaClient;
        HFClient hfClient;
        Channel channel;

        final String CHANNEL_NAME = "ibkchannel";
        final String CHAINCODE_NAME = "ibk";

        final String CA_NAME = "ca.ibk.com";
        final String CA_URL = "http://192.168.56.10:7054";
        final String CA_ADMIN_ID = "admin";
        final String CA_ADMIN_PASSWORD = "adminpw";
        final String PEER_ID = "peer0.org.ibk.com";
        final String PEER_URL = "grpc://192.168.56.10:7051";
        final String ORDERER_ID = "orderer.ibk.com";
        final String ORDERER_URL = "grpc://192.168.56.10:7050";
        final String USER_NAME = "admin";
        final String USER_ORG = "org";
        final String USER_MSP_ID = "OrgMSP";

        //2. fabric client 생성
        CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
        hfcaClient = HFCAClient.createNewInstance(CA_NAME, CA_URL, new Properties());
        hfClient = HFClient.createNewInstance();
        hfcaClient.setCryptoSuite(cryptoSuite);
        hfClient.setCryptoSuite(cryptoSuite);

        //3. admin 컨택
        Enrollment adminEnrollment = hfcaClient.enroll(CA_ADMIN_ID, CA_ADMIN_PASSWORD);
        BlockchainUser blockchainUser = new BlockchainUser();
        blockchainUser.setName(USER_NAME);
        blockchainUser.setAffiliation(USER_ORG);
        blockchainUser.setMspId(USER_MSP_ID);
        blockchainUser.setEnrollment(adminEnrollment);
        hfClient.setUserContext(blockchainUser);


    }
}
