/*
 Navicat Premium Data Transfer

 Source Server         : dev-192.168.70.211
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.70.211:3306
 Source Schema         : password_manager

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 09/01/2024 13:31:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
                         `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'菜单名称\'',
                         `name_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'前端路径name\'',
                         `is_show` tinyint(1) NULL DEFAULT 2 COMMENT '\'状态(1隐藏/2显示, 默认正常)\'',
                         `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'菜单图标\'',
                         `path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'菜单访问路径\'',
                         `sort` int(3) NULL DEFAULT 0 COMMENT '\'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)\'',
                         `parent_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '\'父菜单编号(编号为0时表示根菜单)\'',
                         `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'创建人\'',
                         `component` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'前端路径\'',
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `idx_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '2024-01-02 19:57:27.000', '2024-01-02 19:57:29.000', NULL, '系统管理', '', 2, 'el-icon-video-play', '/mall', 0, 0, NULL, '');
INSERT INTO `menu` VALUES (2, '2024-01-02 19:58:49.000', '2024-01-02 19:58:52.000', NULL, '用户管理', 'user', 2, 'el-icon-user', '/user', 0, 1, NULL, 'User');
INSERT INTO `menu` VALUES (3, '2024-01-02 19:59:35.000', '2024-01-02 19:59:37.000', NULL, '菜单管理', 'menu', 2, 'el-icon-location', '/menu', 0, 1, NULL, 'Menu');
INSERT INTO `menu` VALUES (4, '2024-01-02 20:00:51.000', '2024-01-02 20:00:53.000', NULL, '角色管理', 'role', 2, 'el-icon-setting', '/role', 0, 1, NULL, 'Role');
INSERT INTO `menu` VALUES (5, '2024-01-02 20:00:51.000', '2024-01-02 20:00:53.000', NULL, '日志管理', 'log', 2, 'el-icon-setting', '/log', 0, 1, NULL, 'Log');

SET FOREIGN_KEY_CHECKS = 1;


/*
 Navicat Premium Data Transfer

 Source Server         : dev-192.168.70.211
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.70.211:3306
 Source Schema         : password_manager

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 09/01/2024 13:31:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for relation_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `relation_role_menu`;
CREATE TABLE `relation_role_menu`  (
                                       `role_id` bigint(20) UNSIGNED NOT NULL,
                                       `menu_id` bigint(20) UNSIGNED NOT NULL,
                                       PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
                                       INDEX `fk_relation_role_menu_menu`(`menu_id`) USING BTREE,
                                       CONSTRAINT `fk_relation_role_menu_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                       CONSTRAINT `fk_relation_role_menu_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of relation_role_menu
-- ----------------------------
INSERT INTO `relation_role_menu` VALUES (1, 1);
INSERT INTO `relation_role_menu` VALUES (1, 2);
INSERT INTO `relation_role_menu` VALUES (1, 3);
INSERT INTO `relation_role_menu` VALUES (1, 4);
INSERT INTO `relation_role_menu` VALUES (1, 5);

SET FOREIGN_KEY_CHECKS = 1;


/*
 Navicat Premium Data Transfer

 Source Server         : dev-192.168.70.211
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.70.211:3306
 Source Schema         : password_manager

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 09/01/2024 13:31:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
                         `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'角色名称\'',
                         `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'角色描述\'',
                         `status` tinyint(1) NULL DEFAULT 1 COMMENT '\'用户状态(正常/禁用, 默认正常)\'',
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `idx_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '2023-12-18 17:38:53.000', '2023-12-28 14:45:25.921', NULL, 'admin', '管理员', 1);
INSERT INTO `role` VALUES (2, '2023-12-21 16:46:37.000', '2023-12-21 16:46:40.000', NULL, 'read', '只读用户', 1);

SET FOREIGN_KEY_CHECKS = 1;


/*
 Navicat Premium Data Transfer

 Source Server         : dev-192.168.70.211
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.70.211:3306
 Source Schema         : password_manager

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 09/01/2024 13:31:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `uid` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '\'用戶uid\'',
                         `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'用户名\'',
                         `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'用户密码\'',
                         `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'手机号码\'',
                         `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'邮箱\'',
                         `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'用户昵称\'',
                         `avatar` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500' COMMENT '\'用户头像\'',
                         `status` tinyint(1) NULL DEFAULT 1 COMMENT '\'用户状态(正常/禁用, 默认正常)\'',
                         `role_id` bigint(20) UNSIGNED NULL DEFAULT 1 COMMENT '\'角色id外键\'',
                         `dept_id` bigint(20) UNSIGNED NULL DEFAULT 1 COMMENT '\'部门id外键\'',
                         `create_by` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '\'创建来源\'',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `uk_username`(`username`) USING BTREE,
                         INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE,
                         INDEX `fk_dept_users`(`dept_id`) USING BTREE,
                         INDEX `fk_role_users`(`role_id`) USING BTREE,
                         UNIQUE INDEX `username`(`username`) USING BTREE,
                         CONSTRAINT `fk_dept_users` FOREIGN KEY (`dept_id`) REFERENCES `dept` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                         CONSTRAINT `fk_role_users` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2024-01-09 11:19:19.000', '2024-01-09 11:19:22.000', NULL, 'admin', 'admin', '25285442ebc7d3a0c20047e01d341c31', '18438613802', '123@qq.com', 'admin', 'https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', 1, 1, 1, NULL);
INSERT INTO `user` VALUES (2, '2024-01-09 11:31:41.338', '2024-01-09 11:31:41.338', NULL, '', 'test', '25285442ebc7d3a0c20047e01d341c31', '12345678910', '', '', 'https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', 1, 1, 1, '');

SET FOREIGN_KEY_CHECKS = 1;
