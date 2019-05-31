package com.ibk.spring;

import com.ibk.spring.blockchain.BlockchainUtil;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan("com.ibk.spring.*")
public class App {
    public static void main(String[] args) {
        SpringApplication.run(App.class, args);
    }

    @Bean(name = "blockchainUtil")
    public BlockchainUtil blockchainUtil() throws Exception{
        BlockchainUtil blockchainUtil = new BlockchainUtil();
        return blockchainUtil;
    }
}