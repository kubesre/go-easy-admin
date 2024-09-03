/*
 Navicat Premium Data Transfer

 Source Server         : dev-192.168.70.211
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.70.211:3306
 Source Schema         : go-easy-admin-v2

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 12/08/2024 14:39:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_ldap
-- ----------------------------
DROP TABLE IF EXISTS `sys_ldap`;
CREATE TABLE `sys_ldap`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_by` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '\'创建来源\'',
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `dn` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ou` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `filter` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `mapping` json NULL COMMENT '\'属性映射\'',
  `status` tinyint(1) NULL DEFAULT 2 COMMENT '\'状态(正常/禁用, 默认禁用)\'',
  `ssl` tinyint(1) NULL DEFAULT 2 COMMENT '\'状态(正常/禁用, 默认禁用)\'',
  `admin_user` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_ldap_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_ldap
-- ----------------------------
INSERT INTO `sys_ldap` VALUES (1, '2024-08-07 01:23:36.831', '2024-08-07 01:30:45.501', NULL, 'admin', '192.168.10.12:389', 'ou=user,dc=ienglish,dc=cn', 'UwFaWxMHnDwNMXgJ', 'ou=user,dc=ienglish,dc=cn', '(&(objectClass=organizationalPerson)(uid=%s))', '{\"email\": \"mail\", \"phone\": \"telephoneNumber\", \"username\": \"cn\", \"nick_name\": \"sn\"}', 1, 2, 'cn=admin,dc=ienglish,dc=cn');

SET FOREIGN_KEY_CHECKS = 1;
