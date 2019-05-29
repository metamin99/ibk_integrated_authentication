# Mysql Config
## **1. DB 생성**<br/>
Create Database users;
<br/><br/>

## **2. table 생성**<br />

### **은행 user 생성 sql**
create table users.bank_users(<br>
id INT(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,<br>
userId varchar(50),<br>
userName varchar(50)<br>
userPwd varchar(50)<br>);<br>

<br>

### **저축은행 user 생성 sql**
create table users.saving_users(<br>
id INT(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,<br>
userId varchar(50),<br>
userName varchar(50),<br>
userPwd varchar(50)<br>);

<br>

### **투자증권 user 생성 sql**
create table users.invest_users(<br>
id INT(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,<br>
userId varchar(50),<br>
userName varchar(50),<br>
userPwd varchar(50)<br>);

<br><br><br>

## **3. table 삭제**<br />
DROP TABLE users.bank_users;
DROP TABLE users.saving_users;
DROP TABLE users.invest_users;

<br><br><br>

## **4. data 삽입**<br />
INSERT INTO users.bank_users(userId,userName,userPwd) VALUES('ibk01','홍길','1q2w3e4r@');