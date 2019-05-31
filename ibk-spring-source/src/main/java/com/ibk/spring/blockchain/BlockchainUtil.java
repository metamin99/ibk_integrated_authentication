package com.ibk.spring.blockchain;

import java.util.ArrayList;
import java.util.Collection;
import java.util.LinkedList;
import java.util.Properties;

import static java.nio.charset.StandardCharsets.UTF_8;

import javax.annotation.PostConstruct;

import org.hyperledger.fabric.sdk.ChaincodeID;
import org.hyperledger.fabric.sdk.ChaincodeResponse.Status;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.QueryByChaincodeRequest;
import org.hyperledger.fabric.sdk.TransactionProposalRequest;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.springframework.stereotype.Component;

/**
 * BlockchainUtil
 */
@Component
public class BlockchainUtil {

    // 초기화
    @PostConstruct
    public void init() throws Exception {
        this.initFabricConfiguration();
        this.createDefaultUser();
        this.connectChannel();
    }

    // fabric peer에게 연결할 인스턴스
    HFClient hFClient;

    // ca에 연결할 인스턴스
    HFCAClient hFCAClient;

    // channel 인스턴스
    Channel channel;

    // ca 변수 설정
    final String CA_NAME = "ca.ibk.com";
    final String CA_URL = "http://192.168.56.10:7054";

    // chaincodeID
    final String CHAINCODE_ID = "ibk";

    /**
     * @apiNote 1. 초기화 함수
     * @throws Exception
     */
    public void initFabricConfiguration() throws Exception {
        CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
        // hfclient 인스턴스 생성
        this.hFClient = HFClient.createNewInstance();

        // hfcaclient 인스턴스 생성
        this.hFCAClient = HFCAClient.createNewInstance(CA_NAME, CA_URL, new Properties());

        // crypto suite(암호화캡슐) 설정 - hfclient
        this.hFClient.setCryptoSuite(cryptoSuite);

        // crypto suite(암호화캡슐) 설정 - hfcaclient
        this.hFCAClient.setCryptoSuite(cryptoSuite);

    }

    /**
     * @apiNote 2. Create DefaultUser
     * @throws Exception
     */
    public void createDefaultUser() throws Exception {
        // CA한테 Enrollment 요청
        BlockchainUser blockchainUser = new BlockchainUser();
        Enrollment enrollment = this.hFCAClient.enroll("admin", "adminpw");
        blockchainUser.setName("admin");
        blockchainUser.setAffiliation("org");
        blockchainUser.setMspId("OrgMSP");
        blockchainUser.setEnrollment(enrollment);
        this.hFClient.setUserContext(blockchainUser);
    }

    /**
     * @apiNote 3. Connect Channel
     * @throws Exception
     */
    public void connectChannel() throws Exception {
        String peer_name = "peer2.org.ibk.com";
        String peer_url = "grpc://192.168.56.10:9051";

        String orderer_name = "orderer.ibk.com";
        String orderer_url = "grpc://192.168.56.10:7050";

        String channel_name = "ibkchannel";

        // peer 설정 완료
        Peer peer = hFClient.newPeer(peer_name, peer_url);
        Orderer orderer = hFClient.newOrderer(orderer_name, orderer_url);

        this.channel = this.hFClient.newChannel(channel_name);

        this.channel.addPeer(peer);
        this.channel.addOrderer(orderer);

        this.channel.initialize();
    }

    /**
     * @apiNote util - queryChaincode
     * @throws Exception
     */
    public String queryChaincode(String functionName, ArrayList<String> params) throws Exception {
        ChaincodeID chaincodeID = ChaincodeID.newBuilder().setName(CHAINCODE_ID).build();
        QueryByChaincodeRequest queryByChaincodeRequest = hFClient.newQueryProposalRequest();

        queryByChaincodeRequest.setChaincodeID(chaincodeID);
        queryByChaincodeRequest.setFcn(functionName);
        queryByChaincodeRequest.setArgs(params);
        queryByChaincodeRequest.setProposalWaitTime(10000);

        Collection<ProposalResponse> res = channel.queryByChaincode(queryByChaincodeRequest);

        // reponse collect
        String stringResponse = "";
        for (ProposalResponse pres : res) {
            stringResponse = new String(pres.getChaincodeActionResponsePayload());
        }
        return stringResponse;
    }

    /**
     * @apiNote util - invokeChaincode
     * @throws Exception
     */
    public void invokeChaincode(String functionName, ArrayList<String> params) throws Exception {
        ChaincodeID chaincodeID = ChaincodeID.newBuilder().setName(CHAINCODE_ID).build();
        TransactionProposalRequest transactionProposalRequest = hFClient.newTransactionProposalRequest();

        transactionProposalRequest.setChaincodeID(chaincodeID);
        transactionProposalRequest.setFcn(functionName);
        transactionProposalRequest.setArgs(params);
        transactionProposalRequest.setProposalWaitTime(10000);

        Collection<ProposalResponse> res = channel.sendTransactionProposal(transactionProposalRequest,
                channel.getPeers());

        // reponse collect
        Collection<ProposalResponse> success = new LinkedList<>();
        Collection<ProposalResponse> fail = new LinkedList<>();

        for (ProposalResponse pres : res) {
            if (pres.getStatus() == Status.SUCCESS) {
                success.add(pres);
            } else {
                fail.add(pres);
            }
        }

        if (fail.size() > 0) {
            ProposalResponse firstTransactionProposalResponse = fail.iterator().next();
            throw new Exception("Not enough endorsers for:" + fail.size() + " endorser error: "
                    + firstTransactionProposalResponse.getMessage() + ". Was verified: "
                    + firstTransactionProposalResponse.isVerified());
        }

        ProposalResponse resp = success.iterator().next();
        byte[] x = resp.getChaincodeActionResponsePayload();

        String resultAsString;
        if (x != null) {
            resultAsString = new String(x, UTF_8);
        }
        channel.sendTransaction(success);
    }

}