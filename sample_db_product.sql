-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: sample_db
-- ------------------------------------------------------
-- Server version	8.0.36-0ubuntu0.20.04.1

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
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `img_url` varchar(250) NOT NULL,
  `speccs` varchar(250) NOT NULL,
  `price` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (13,'http://localhost:8080/uploads/apple.jpg','iPhone 15 and iPhone 15 Plus. Dynamic Island. 48MP Main camera with 2x Telephoto. All-day battery life.','79999','iPhone'),(14,'http://localhost:8080/uploads/motorola.jpg',' 108 MP ultra high-res camera system​. Capture photos with breathtaking detail even in the most challenging light using 9x larger ultra pixels.','49999','Motorola'),(15,'http://localhost:8080/uploads/realme.jpg','Powerful Octa-core Chipset. Processor ; 18W | 5000mAh. Charging & Battery ; 17.13cm (6.74\'\'). Display ; 4GB+128GB | 6GB+128GB I 6GB+64GB. Storage & RAM.','39990','Realme'),(16,'http://localhost:8080/uploads/poco.png','Processor. Qualcomm Snapdragon 732G | Octa Core | 2.3 GHz ; Rear Camera. 64MP + 13MP + 2MP + 2MP ; Front Camera. 20MP ; Display.','28950','POCO'),(17,'http://localhost:8080/uploads/vivo.jpg','Experience the power of 5G with the vivo V29. Featuring Qualcomm Snapdragon 778G processor, 4600mAh battery supported 80W fast charging, and more.','25949','Vivo'),(18,'http://localhost:8080/uploads/redmi.jpg','SuperPower.SuperNote! Redmi Note 13 Pro 5G with the all newdesign packed in a double sided glassbody in 3 elegant shades.200MP Camera with 4X Zoom.','24999','Redmi');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-12 17:50:29
