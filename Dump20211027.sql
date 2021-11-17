CREATE DATABASE  IF NOT EXISTS `hedgex` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `hedgex`;
-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: hedgex
-- ------------------------------------------------------
-- Server version	8.0.27-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `burn`
--

DROP TABLE IF EXISTS `burn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `burn` (
  `transaction` varchar(80) NOT NULL COMMENT '交易哈希',
  `account` varchar(45) DEFAULT NULL COMMENT '回收流动性的账户',
  `amount` bigint DEFAULT NULL COMMENT '销毁流动性代币数量',
  `block` bigint DEFAULT NULL,
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `block` (`block`),
  KEY `amount` (`amount`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='销毁流动性记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `burn`
--

LOCK TABLES `burn` WRITE;
/*!40000 ALTER TABLE `burn` DISABLE KEYS */;
/*!40000 ALTER TABLE `burn` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `explosive`
--

DROP TABLE IF EXISTS `explosive`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `explosive` (
  `transaction` varchar(80) NOT NULL COMMENT '交易哈希',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '爆仓账户',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '区块高度',
  `amount` int NOT NULL DEFAULT '0' COMMENT '爆仓量，单位为合约“张数”',
  `price` bigint NOT NULL DEFAULT '0' COMMENT '爆仓价格',
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `block` (`block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='爆仓记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `explosive`
--

LOCK TABLES `explosive` WRITE;
/*!40000 ALTER TABLE `explosive` DISABLE KEYS */;
/*!40000 ALTER TABLE `explosive` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `interest`
--

DROP TABLE IF EXISTS `interest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `interest` (
  `transaction` varchar(80) NOT NULL COMMENT '交易哈希',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '被收取利息的账户',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '区块高度',
  `amount` int NOT NULL DEFAULT '0' COMMENT '收取的利息量，保证金数量',
  `price` bigint NOT NULL DEFAULT '0' COMMENT '持仓价格',
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `block` (`block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='利息收取记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interest`
--

LOCK TABLES `interest` WRITE;
/*!40000 ALTER TABLE `interest` DISABLE KEYS */;
/*!40000 ALTER TABLE `interest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mint`
--

DROP TABLE IF EXISTS `mint`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mint` (
  `transaction` varchar(80) NOT NULL COMMENT '交易哈希',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '提供流动性的账户',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '产生的流动性代币数量',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '产生的流动性代币数量',
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `block` (`block`),
  KEY `amount` (`amount`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='增加流动性记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mint`
--

LOCK TABLES `mint` WRITE;
/*!40000 ALTER TABLE `mint` DISABLE KEYS */;
/*!40000 ALTER TABLE `mint` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recharge`
--

DROP TABLE IF EXISTS `recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `recharge` (
  `transaction` varchar(80) NOT NULL DEFAULT '' COMMENT '交易哈希，唯一主键',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '账户地址',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '充值数量，保证金',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '区块高度',
  PRIMARY KEY (`transaction`),
  KEY `block` (`block`),
  KEY `amount` (`amount`),
  KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recharge`
--

LOCK TABLES `recharge` WRITE;
/*!40000 ALTER TABLE `recharge` DISABLE KEYS */;
/*!40000 ALTER TABLE `recharge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trade`
--

DROP TABLE IF EXISTS `trade`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `trade` (
  `transaction` varchar(80) NOT NULL COMMENT '交易哈希',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '开平仓账号',
  `direction` smallint NOT NULL DEFAULT '0' COMMENT '交易方向，开多、开空、平多、平空，分别为1，-1，-2，2',
  `amount` int unsigned NOT NULL DEFAULT '0' COMMENT '交易数量，合约张数',
  `price` bigint NOT NULL DEFAULT '0' COMMENT '交易价格',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '区块高度',
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `direction` (`direction`),
  KEY `block` (`block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户交易记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trade`
--

LOCK TABLES `trade` WRITE;
/*!40000 ALTER TABLE `trade` DISABLE KEYS */;
/*!40000 ALTER TABLE `trade` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `account` varchar(45) NOT NULL COMMENT '用户地址，格式为：0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e',
  `contract` varchar(45) NOT NULL COMMENT '所属合约',
  `margin` bigint NOT NULL DEFAULT '0' COMMENT '保证金数量',
  `lposition` int NOT NULL DEFAULT '0' COMMENT '多仓持仓量',
  `lprice` bigint NOT NULL DEFAULT '0' COMMENT '多仓持仓价',
  `sposition` int NOT NULL DEFAULT '0' COMMENT '空仓持仓量',
  `sprice` bigint NOT NULL DEFAULT '0' COMMENT '空仓持仓价',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '最后一次更新的区块高度',
  PRIMARY KEY (`account`,`contract`),
  KEY `block` (`block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表，用来存储合约的用户数据，随着链上区块高度的增加而更新';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `withdraw`
--

DROP TABLE IF EXISTS `withdraw`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `withdraw` (
  `transaction` varchar(80) NOT NULL DEFAULT '' COMMENT '交易哈希，唯一主键',
  `account` varchar(45) NOT NULL DEFAULT '' COMMENT '提现账户地址',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '提现数量，保证金',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '区块高度',
  PRIMARY KEY (`transaction`),
  KEY `account` (`account`),
  KEY `block` (`block`),
  KEY `amount` (`amount`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='提现记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `withdraw`
--

LOCK TABLES `withdraw` WRITE;
/*!40000 ALTER TABLE `withdraw` DISABLE KEYS */;
/*!40000 ALTER TABLE `withdraw` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-27 18:40:14
