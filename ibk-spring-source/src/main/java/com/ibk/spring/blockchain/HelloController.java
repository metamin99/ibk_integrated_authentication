package com.ibk.spring.blockchain;

import java.util.ArrayList;

import javax.annotation.Resource;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import lombok.extern.slf4j.Slf4j;

@RestController
@Slf4j
public class HelloController {


    @Resource(name ="blockchainUtil")
    BlockchainUtil blockchainUtil;

    @RequestMapping("/")
    String list() throws Exception {
        log.info("------------------QueryAllUsers");
        String functionName = "QueryAllUsers";
        ArrayList<String> params = new ArrayList<>();

        return blockchainUtil.queryChaincode(functionName, params);
    }

    @RequestMapping("/id")
    String list(String userId) throws Exception {
        log.info("------------------QueryUserByUserId");
        String functionName = "QueryUserByUserId";
        ArrayList<String> params = new ArrayList<>();
        params.add(userId);

        return blockchainUtil.queryChaincode(functionName, params);
    }

    @RequestMapping("/insert")
    void insert(UserVo userVo) throws Exception {
        log.info("------------------CreateUser");
        String functionName = "CreateUser";
        ArrayList<String> params = new ArrayList<>();
        params.add(userVo.getUserId());
        params.add(userVo.getUserName());
        params.add(userVo.getUserPwd());

        blockchainUtil.invokeChaincode(functionName, params);
    }
}
