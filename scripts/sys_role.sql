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

 Date: 12/08/2024 15:47:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_by` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '\'创建来源\'',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'角色名称\'',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'角色描述\'',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '\'用户状态(正常/禁用, 默认正常)\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '2024-08-12 15:13:00.291', '2024-08-12 15:13:00.291', NULL, 'admin', 'dev', '开发', 1);
INSERT INTO `sys_role` VALUES (2, '2024-08-12 15:13:13.556', '2024-08-12 15:13:13.556', NULL, 'admin', 'test', '测试', 1);
INSERT INTO `sys_role` VALUES (3, '2024-08-12 15:13:28.250', '2024-08-12 15:13:28.250', NULL, 'admin', 'admin', '管理员', 1);

SET FOREIGN_KEY_CHECKS = 1;
