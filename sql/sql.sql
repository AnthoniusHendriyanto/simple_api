CREATE DATABASE `customers` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- customers.customers definition

CREATE TABLE `customers` (
    `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
    `customerId` varchar(100) NOT NULL,
    `phoneNumber` varchar(100) NOT NULL,
    `status` tinyint NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `customers_UN` (`customerId`),
    KEY `customers_FK` (`status`),
    CONSTRAINT `customers_FK` FOREIGN KEY (`status`) REFERENCES `customers_status` (`status`)
);

-- customers.customers_status definition
CREATE TABLE `customers_status` (
    `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
    `status` tinyint NOT NULL,
    `description` char(32) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `customers_status_UN` (`status`)
);

INSERT INTO customers.customers_status (status, description) VALUES (0, "Inactive");
INSERT INTO customers.customers_status (status, description) VALUES (1, "Active");