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

 Date: 12/08/2024 15:47:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_apis
-- ----------------------------
DROP TABLE IF EXISTS `system_apis`;
CREATE TABLE `system_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_by` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '\'创建来源\'',
  `path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `method` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `api_group` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_system_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_apis
-- ----------------------------
INSERT INTO `system_apis` VALUES (1, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/user/create', 'POST', '创建用户', '用户管理');
INSERT INTO `system_apis` VALUES (2, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/user/delete', 'POST', '删除用户', '用户管理');
INSERT INTO `system_apis` VALUES (3, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/user/update/:id', 'POST', '更新用户', '用户管理');
INSERT INTO `system_apis` VALUES (4, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/user/list', 'GET', '用户列表', '用户管理');
INSERT INTO `system_apis` VALUES (5, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/user/get/:id', 'GET', '用户详情', '用户管理');
INSERT INTO `system_apis` VALUES (6, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/menu/create', 'POST', '创建菜单', '菜单管理');
INSERT INTO `system_apis` VALUES (7, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/menu/delete/:id', 'POST', '删除菜单', '菜单管理');
INSERT INTO `system_apis` VALUES (8, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/menu/update/:id', 'POST', '更新菜单', '菜单管理');
INSERT INTO `system_apis` VALUES (9, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/menu/list', 'GET', '菜单列表', '菜单管理');
INSERT INTO `system_apis` VALUES (10, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/menu/get/:id', 'GET', '菜单详情', '菜单管理');
INSERT INTO `system_apis` VALUES (11, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/role/create', 'POST', '创建角色', '角色管理');
INSERT INTO `system_apis` VALUES (12, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/role/delete/:id', 'POST', '删除角色', '角色管理');
INSERT INTO `system_apis` VALUES (13, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/role/update/:id', 'POST', '更新角色', '角色管理');
INSERT INTO `system_apis` VALUES (14, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/role/list', 'GET', '角色列表', '角色管理');
INSERT INTO `system_apis` VALUES (15, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/role/get/:id', 'GET', '角色详情', '角色管理');
INSERT INTO `system_apis` VALUES (16, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/create', 'POST', '创建路由', '路由管理');
INSERT INTO `system_apis` VALUES (17, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/delete/:id', 'POST', '删除路由', '路由管理');
INSERT INTO `system_apis` VALUES (18, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/update/:id', 'POST', '更新路由', '路由管理');
INSERT INTO `system_apis` VALUES (19, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/list', 'GET', '路由列表', '路由管理');
INSERT INTO `system_apis` VALUES (20, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/get/:id', 'GET', '路由详情', '路由管理');
INSERT INTO `system_apis` VALUES (21, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/apis/get/group', 'GET', '路由组', '路由管理');
INSERT INTO `system_apis` VALUES (22, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/ldap/create', 'POST', '创建LDAP', 'LDAP管理');
INSERT INTO `system_apis` VALUES (23, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/ldap/list', 'GET', 'LDAP列表', 'LDAP管理');
INSERT INTO `system_apis` VALUES (24, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/ldap/ping', 'POST', 'LDAP测试', 'LDAP管理');
INSERT INTO `system_apis` VALUES (25, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/rbac/create', 'POST', '创建授权', '授权管理');
INSERT INTO `system_apis` VALUES (26, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/rbac/delete/*', 'POST', '删除授权', '授权管理');
INSERT INTO `system_apis` VALUES (27, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/rbac/update/*', 'POST', '更新授权', '授权管理');
INSERT INTO `system_apis` VALUES (28, '2024-08-08 17:27:05.000', '2024-08-08 17:27:05.000', NULL, 'admin', '/sys/rbac/list', 'GET', '授权列表', '授权管理');

SET FOREIGN_KEY_CHECKS = 1;
