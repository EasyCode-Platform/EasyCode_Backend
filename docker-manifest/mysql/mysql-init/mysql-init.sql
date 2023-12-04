-- 初始化 ec_backend 数据库
CREATE DATABASE IF NOT EXISTS ec_backend;

-- 创建用户
CREATE USER 'ec_backend'@'%' IDENTIFIED BY 'scut2023';
GRANT ALL PRIVILEGES ON ec_backend.* TO 'ec_backend'@'%';
FLUSH PRIVILEGES;

-- 选择 ec_backend 数据库
USE ec_backend;

--- 创建 apps 表对应 App 结构体
CREATE TABLE IF NOT EXISTS `apps` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uid` BINARY(16) NOT NULL,
    `team_id` INT,
    `name` TEXT,
    `component_id` TEXT NOT NULL,
    `config` JSON,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
    );

-- 创建 app_v2 表对应 AppV2 结构体
CREATE TABLE IF NOT EXISTS `app_data` (
    `aid` BINARY(16) NOT NULL PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
    );

-- 创建 tables 表对应 Table 结构体
CREATE TABLE IF NOT EXISTS `tables` (
    `tid` BINARY(16) NOT NULL PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `app_aid` BINARY(16),
    FOREIGN KEY (`app_aid`) REFERENCES `app_v2`(`aid`)
    );