package com.ibk.spring.blockchain;

import java.util.Set;

import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

/**
 * BlockchainUser
 */
@Data
@Setter
@Getter
public class BlockchainUser implements User{

    private String name;
    private String affiliation;
    private Enrollment enrollment;
    private String enrollmentSecret;
    private String account;
    private Set<String> roles;
    private String mspId;
}