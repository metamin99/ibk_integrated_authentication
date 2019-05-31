package com.ibk.spring.blockchain;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
@Setter
@Getter
public class UserVo {
    private int id;
    private String userId;
    private String userName;
    private String userPwd;
    private String lastLoginTime;
    private String limitLoginTime;
}
