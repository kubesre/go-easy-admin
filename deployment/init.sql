INSERT INTO `menu` VALUES (1, '2023-12-01 21:56:29.655', '2023-12-01 21:56:29.655', NULL, 'public', 'icon', '/api/v1/ping', 0, 0, '');
INSERT INTO `menu` VALUES (2, '2023-12-02 18:27:16.255', '2023-12-02 18:27:16.255', NULL, '系统管理', 'icon', '/api/v1/admin', 0, 0, '');
INSERT INTO `menu` VALUES (3, '2023-12-02 18:27:51.255', '2023-12-02 18:27:51.255', NULL, '用户管理', 'icon', '/api/v1/user', 0, 2, '');
INSERT INTO `menu` VALUES (4, '2023-12-02 18:28:06.467', '2023-12-02 18:28:06.467', NULL, '部门管理', 'icon', '/api/v1/dept', 0, 2, '');
INSERT INTO `menu` VALUES (5, '2023-12-05 12:41:42.000', '2023-12-05 12:41:44.000', NULL, '菜单管理', 'icon', '/api/v1/menu', 0, 2, NULL);
INSERT INTO `menu` VALUES (6, '2023-12-05 12:42:32.000', '2023-12-05 12:42:34.000', NULL, '审计管理', 'icon', '/api/v1/log', 0, 2, NULL);
INSERT INTO `menu` VALUES (7, '2023-12-05 12:43:41.000', '2023-12-05 12:43:45.000', NULL, '角色管理', 'icon', '/api/v1/role', 0, 2, NULL);
INSERT INTO `menu` VALUES (8, '2023-12-05 12:44:22.000', '2023-12-05 12:44:24.000', NULL, '授权管理', 'icon', '/api/v1/policy', 0, 2, NULL);

INSERT INTO `role` VALUES (1, '2023-12-01 20:05:17.000', '2023-12-01 22:26:03.122', NULL, 'admin', '管理员');



INSERT INTO `dept` VALUES (1, '2023-12-01 20:05:52.000', '2023-12-01 20:05:55.000', NULL, '公司顶部部门', 0, 0);
INSERT INTO `dept` VALUES (2, '2023-12-02 19:42:29.356', '2023-12-02 22:58:36.112', NULL, '郑州研发中心', 0, 0);
INSERT INTO `dept` VALUES (3, '2023-12-02 21:46:31.487', '2023-12-02 21:46:31.487', NULL, '运维部门', 0, 2);
INSERT INTO `dept` VALUES (4, '2023-12-02 21:46:42.011', '2023-12-02 21:46:42.011', NULL, '测试部门', 0, 2);
INSERT INTO `dept` VALUES (5, '2023-12-02 21:47:31.413', '2023-12-02 21:47:31.413', NULL, '系统运维组', 0, 3);

INSERT INTO `user` VALUES (1, '2023-12-01 20:03:37.000', '2023-12-01 20:03:40.000', NULL, 'zhangxinglei', 'admin', '123456', '12345678901', NULL, NULL, 'https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', 1, 1, 1, '系统');

INSERT INTO `casbin_rule` VALUES (1, 'p', '1', '/health', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', '1', '/api/v1/login', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', '1', '/api/v1/register', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', '1', '/api/v1/userinfo/', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', '1', '/api/v1/user/list', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', '1', '/api/v1/role/', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', '1', '/api/v1/role/update/', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', '1', '/api/v1/role/bind_menu', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', '1', '/api/v1/role/del', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', '1', '/api/v1/dept/add', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', '1', '/api/v1/dept/list', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', '1', '/api/v1/dept/info/', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', '1', '/api/v1/dept/del', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', '1', '/api/v1/menu/add', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', '1', '/api/v1/menu/list', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', '1', '/api/v1/user/list', 'GET', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', '1', '/api/v1/policy/add', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', '1', '/api/v1/policy/del', 'POST', '管理员', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', '1', '/api/v1/policy/list ', 'GET', '管理员', '', '');


INSERT INTO `api_path` VALUES (1, '2023-12-05 12:47:45.000', '2023-12-05 12:47:48.000', NULL, '/health', 'GET', '健康检查', 1);
INSERT INTO `api_path` VALUES (2, '2023-12-05 10:52:30.000', '2023-12-05 10:52:32.000', NULL, '/api/v1/login', 'POST', '登录', 1);
INSERT INTO `api_path` VALUES (3, '2023-12-05 10:52:53.000', '2023-12-05 10:52:55.000', NULL, '/api/v1/register', 'POST', '注册', 1);
INSERT INTO `api_path` VALUES (4, '2023-12-05 10:53:19.000', '2023-12-05 10:53:22.000', NULL, '/api/v1/userinfo/', 'GET', '用户详情', 3);
INSERT INTO `api_path` VALUES (5, '2023-12-05 10:53:43.000', '2023-12-05 10:53:45.000', NULL, '/api/v1/user/list', 'GET', '用户列表', 3);
INSERT INTO `api_path` VALUES (6, '2023-12-05 10:54:15.000', '2023-12-05 10:54:18.000', NULL, '/api/v1/role/', 'GET', '角色详情', 7);
INSERT INTO `api_path` VALUES (7, '2023-12-05 10:55:29.000', '2023-12-05 10:55:33.000', NULL, '/api/v1/role/add', 'POST', '添加角色', 7);
INSERT INTO `api_path` VALUES (8, '2023-12-05 10:55:54.000', '2023-12-05 10:55:56.000', NULL, '/api/v1/role/update/', 'POST', '更新角色', 7);
INSERT INTO `api_path` VALUES (9, '2023-12-05 10:56:27.000', '2023-12-05 10:56:30.000', NULL, '/api/v1/role/bind_menu', 'POST', '角色菜单绑定', 7);
INSERT INTO `api_path` VALUES (10, '2023-12-05 10:56:52.000', '2023-12-05 10:56:54.000', NULL, '/api/v1/role/del', 'POST', '删除角色', 7);
INSERT INTO `api_path` VALUES (11, '2023-12-05 10:57:17.000', '2023-12-05 10:57:20.000', NULL, '/api/v1/dept/add', 'POST', '添加部门', 4);
INSERT INTO `api_path` VALUES (12, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/dept/list', 'GET', '部门列表', 4);
INSERT INTO `api_path` VALUES (13, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/dept/info/', 'GET', '部门详情', 4);
INSERT INTO `api_path` VALUES (14, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/dept/del', 'POST', '删除部门', 4);
INSERT INTO `api_path` VALUES (15, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/menu/add', 'POST', '添加菜单', 5);
INSERT INTO `api_path` VALUES (16, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/menu/list', 'GET', '菜单列表', 5);
INSERT INTO `api_path` VALUES (17, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/policy/add', 'POST', '添加策略', 8);
INSERT INTO `api_path` VALUES (18, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/policy/del', 'POST', '删除策略', 8);
INSERT INTO `api_path` VALUES (19, '2023-12-05 10:57:17.000', '2023-12-05 10:57:17.000', NULL, '/api/v1/policy/list ', 'GET', '策略列表', 8);