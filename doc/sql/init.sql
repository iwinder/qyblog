-- 媒体库类型
INSERT INTO  qy_file_lib_type (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,identifier,`type`,created_by,updated_by) VALUES
('file-lib-type-a25nPd',0,1,'null','2022-10-10 15:50:35.648000000','2022-10-10 15:50:35.649000000',NULL,0,'本地',1,'SYSTEM',1,1),
('file-lib-type-lMkoMK',0,1,'null','2022-10-10 15:53:24.560000000','2022-10-10 15:53:24.562000000',NULL,0,'七牛云OSS',2,'SYSTEM',1,1),
('file-lib-type-qPpqOA',0,2,'null','2022-10-10 15:56:56.712000000','2022-10-23 09:57:03.200000000',NULL,0,'阿里云OSS',3,'SYSTEM',1,1);
-- 媒体库配置
INSERT INTO  qy_file_lib_config (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,access_key,secret_key,bucket,prefix,`domain`,endpoint,type_id,created_by,updated_by) VALUES
('file-lib-config-n4MG3B',0,1,'null','2022-10-24 14:48:12.414000000','2022-10-24 14:48:12.415000000',NULL,0,'','','','/home/wind/Work/program/common/domain','','',1,1,1),
('file-lib-config-NO7V3P',0,1,'null','2022-10-24 14:55:14.436000000','2022-10-29 17:57:47.915000000',NULL,0,'GS7m9SZJG','iMTyY','zz','','http://qytest.windcoder.com','',2,1,1);
-- 菜单分组
INSERT INTO  qy_sys_menus_agent (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,ftype,created_by,updated_by) VALUES
('menus-j82B8q',0,1,'null','2022-11-12 06:01:33.941000000','2022-11-12 06:01:33.943000000',NULL,0,'顶部菜单','SYSTEM',1,1),
('menus-Do0q8E',0,1,'null','2022-11-12 06:03:55.362000000','2022-11-12 06:03:55.362000000',NULL,0,'页脚菜单','SYSTEM',1,1);

-- 账户信息
INSERT INTO  qy_sys_user (instance_id,sort,status_flag,extend_shadow,created_at,updated_at, deleted_at,username,nickname,avatar,password,email,phone,admin_flag,deleted_flag,created_by,updated_by) VALUES
('user-y43D1p',1,2,'','2022-09-25 06:40:16.526000000','2022-10-07 17:55:23.812000000', NULL,'admin','哈哈','','$2a$10$5waXIiNaqAtsOHt0BhedveYPe9VR5VQGtf/59mi4gM03KP7Ze3HrW','111','',0,0,1,1);

-- 账户-角色
INSERT INTO qy_sys_user_role (user_id,role_id,created_at,updated_at) VALUES
(1,1,'2022-10-07 17:11:24.204000000','2022-10-07 17:11:24.204000000');
-- 角色
INSERT INTO qy_sys_role (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,identifier,created_by,updated_by) VALUES
(1,'role-wBYeBx',0,1,'null','2022-09-26 13:56:13.766000000','2022-10-07 01:59:50.807000000',NULL,0,'超级管理员','',1,1);
-- api 分组
INSERT INTO qy_sys_api_group (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,name,identifier) VALUES
(1,'apis-group-On58E6',0,1,'null',1,1,'2022-10-06 06:12:18.902000000','2022-10-06 06:12:50.396000000',NULL,0,'用户管理','sys-user'),
(2,'apis-group-aXYdnA',0,1,'null',1,1,'2022-10-06 06:13:31.499000000','2022-10-06 06:13:31.500000000',NULL,0,'角色管理','sys_role'),
(3,'apis-group-5n30Xx',0,1,'null',1,1,'2022-10-06 06:13:45.907000000','2022-10-06 06:13:45.907000000',NULL,0,'API管理','sys_apis'),
(4,'apis-group-wEdgj6',0,1,'null',1,1,'2022-10-06 06:15:31.503000000','2022-10-06 06:15:31.504000000',NULL,0,'Api分组管理','sys_apis_group'),
(5,'apis-group-5jb9X0',0,1,'null',1,1,'2022-10-06 06:16:02.781000000','2022-10-06 06:16:02.781000000',NULL,0,'后端菜单管理','sys_menusAdmin'),
(7,'apis-group-xEZPX5',0,1,'null',1,1,'2022-12-17 15:26:05.426000000','2022-12-17 15:26:05.427000000',NULL,0,'站点管理','sys_site_config'),
(8,'apis-group-oXW3Xd',0,1,'null',1,1,'2022-12-24 11:34:57.752000000','2022-12-24 11:34:57.753000000',NULL,0,'评论管理','sys_comment'),
(9,'apis-group-AEpqnM',0,1,'null',1,1,'2022-12-24 12:28:07.437000000','2022-12-24 12:28:07.455000000',NULL,0,'媒体库管理','file-lib'),
(10,'apis-group-3X1GE6',0,1,'null',1,1,'2022-12-24 12:35:58.866000000','2022-12-24 12:55:16.038000000',NULL,0,'友链管理','link-friends'),
(11,'apis-group-JX0aj0',0,1,'null',1,1,'2022-12-24 12:38:50.160000000','2022-12-24 12:55:23.363000000',NULL,0,'短链接管理','link-shorts');
INSERT INTO qy_sys_api_group (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,name,identifier) VALUES
(12,'apis-group-ojl4jB',0,1,'null',1,1,'2022-12-24 12:55:51.773000000','2022-12-24 12:55:51.776000000',NULL,0,'导航管理','link-menus'),
(13,'apis-group-KXJenO',0,1,'null',1,1,'2022-12-24 13:24:10.521000000','2022-12-24 13:24:10.523000000',NULL,0,'导航分类管理','link-menus-agent'),
(14,'apis-group-vEOAX3',0,1,'null',1,1,'2022-12-24 13:30:58.492000000','2022-12-24 13:30:58.493000000',NULL,0,'标签管理','blog-tags'),
(15,'apis-group-OENlXk',0,1,'null',1,1,'2022-12-24 13:49:54.979000000','2022-12-24 13:49:54.980000000',NULL,0,'分类管理','blog-category'),
(16,'apis-group-REAOX6',0,1,'null',1,1,'2022-12-24 13:58:36.745000000','2022-12-24 13:58:36.746000000',NULL,0,'文章管理','blog-article');
-- api信息
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(1,'api-dB7Zgz',0,0,'null',1,1,'2022-09-04 18:26:27.820000000','2022-10-06 07:38:54.866000000',NULL,0,'/api/admin/v1/user','POST','新增用户',1,'用户管理','sys-user'),
(2,'api-4VkLvX',0,0,'null',1,1,'2022-09-04 18:26:39.711000000','2022-10-06 07:39:01.230000000',NULL,0,'/api/admin/v1/user/{id}','PUT','更新用户',1,'用户管理','sys-user'),
(3,'api-xBE8g9',0,0,'null',1,1,'2022-09-04 18:26:49.794000000','2022-10-06 07:39:08.432000000',NULL,0,'/api/admin/v1/user/{id}','DELETE','删除用户',1,'用户管理','sys-user'),
(4,'api-JVqdB8',0,0,'null',1,1,'2022-09-04 18:27:33.198000000','2022-10-06 07:41:16.352000000',NULL,0,'/api/admin/v1/user/{id}','GET','获取用户详情',1,'用户管理','sys-user'),
(5,'api-dgNavD',0,0,'null',1,1,'2022-09-04 18:28:02.896000000','2022-10-06 07:39:20.689000000',NULL,0,'/api/admin/v1/user','GET','批量获取用户',1,'用户管理','sys-user'),
(6,'api-wv8lvY',0,0,'null',1,1,'2022-09-04 18:28:17.489000000','2022-10-06 07:41:54.195000000',NULL,0,'/api/admin/v1/user','DELETE','批量删除用户',1,'用户管理','sys-user'),
(7,'api-MgxABn',0,0,'null',1,1,'2022-09-04 18:35:04.571000000','2022-10-06 07:39:26.350000000',NULL,0,'/api/admin/v1/user/myInfo','GET','获取用户个人信息',1,'用户管理','sys-user'),
(8,'api-rVMwgz',0,0,'null',1,1,'2022-09-04 18:36:54.288000000','2022-10-06 07:39:38.678000000',NULL,0,'/api/admin/v1/role','POST','创建角色',2,'角色管理','sys_role'),
(9,'api-oB3DVZ',0,0,'null',1,1,'2022-09-04 18:38:10.228000000','2022-10-06 07:39:47.005000000',NULL,0,'/api/admin/v1/role/{id}','PUT','更新角色',2,'角色管理','sys_role'),
(10,'api-mBR5gZ',0,0,'null',1,1,'2022-09-04 18:38:25.633000000','2022-10-06 07:39:54.190000000',NULL,0,'/api/admin/v1/role/{id}','DELETE','删除角色',2,'角色管理','sys_role');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(11,'api-oVnmBY',0,0,'null',1,1,'2022-09-04 18:38:36.086000000','2022-10-06 07:40:02.769000000',NULL,0,'/api/admin/v1/role/{id}','GET','获取角色详情',2,'角色管理','sys_role'),
(12,'api-4gK6vL',0,0,'null',1,1,'2022-09-04 18:39:22.965000000','2022-10-06 07:40:11.045000000',NULL,0,'/api/admin/v1/role','GET','批量获取角色',2,'角色管理','sys_role'),
(13,'api-3vA1vq',0,0,'null',1,1,'2022-09-04 18:39:38.529000000','2022-10-06 07:40:20.730000000',NULL,0,'/api/admin/v1/role','DELETE','批量删除角色',2,'角色管理','sys_role'),
(14,'api-rgAegY',0,0,'null',1,1,'2022-09-04 18:41:38.516000000','2022-10-06 07:35:55.346000000',NULL,0,'/api/admin/v1/sysapi','POST','新增API',3,'API管理','sys_apis'),
(15,'api-3VJ4g1',0,0,'null',1,1,'2022-09-04 18:42:10.357000000','2022-10-06 07:29:19.277000000',NULL,0,'/api/admin/v1/sysapi/{id}','PUT','更新API',3,'API管理','sys_apis'),
(16,'api-Egkogz',0,0,'null',1,1,'2022-09-04 18:42:20.820000000','2022-10-06 07:32:58.537000000',NULL,0,'/api/admin/v1/sysapi/{id}','DELETE','删除API',3,'API管理','sys_apis'),
(17,'api-0Bwng9',0,0,'null',1,1,'2022-09-04 18:42:35.216000000','2022-10-06 07:33:11.427000000',NULL,0,'/api/admin/v1/sysapi/{id}','GET','获取API详情',3,'API管理','sys_apis'),
(18,'api-kVwjBq',0,0,'null',1,1,'2022-09-04 18:43:42.127000000','2022-10-06 07:33:17.667000000',NULL,0,'/api/admin/v1/sysapi','GET','批量获取API',3,'API管理','sys_apis'),
(19,'api-RgdPgq',0,0,'null',1,1,'2022-09-04 18:43:51.461000000','2022-10-06 07:38:14.655000000',NULL,0,'/api/admin/v1/sysapi','DELETE','批量删除API',3,'API管理','sys_apis'),
(20,'api-4gWOg0',0,0,'null',1,1,'2022-09-04 18:45:05.265000000','2022-10-06 07:33:53.400000000',NULL,0,'/api/admin/v1/menusAdmin','POST','新增菜单',5,'后端菜单管理','sys_menusAdmin');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(21,'api-ngzbga',0,0,'null',1,1,'2022-09-04 18:45:27.754000000','2022-10-06 07:37:20.944000000',NULL,0,'/api/admin/v1/menusAdmin/{id}','PUT','更新菜单',5,'后端菜单管理','sys_menusAdmin'),
(22,'api-mvpWVa',0,0,'null',1,1,'2022-09-04 18:45:42.542000000','2022-10-06 07:37:28.249000000',NULL,0,'/api/admin/v1/menusAdmin/{id}','DELETE','删除菜单',5,'后端菜单管理','sys_menusAdmin'),
(23,'api-MvOQBW',0,0,'null',1,1,'2022-09-04 18:45:55.964000000','2022-10-06 07:37:34.581000000',NULL,0,'/api/admin/v1/menusAdmin/{id}','GET','获取菜单详情',5,'后端菜单管理','sys_menusAdmin'),
(24,'api-NgkqVZ',0,0,'null',1,1,'2022-09-04 18:46:08.649000000','2022-10-06 07:38:48.707000000',NULL,0,'/api/admin/v1/menusAdmin','GET','批量获取菜单',5,'后端菜单管理','sys_menusAdmin'),
(25,'api-lvqYv0',0,0,'null',1,1,'2022-09-04 18:46:19.793000000','2022-10-06 07:38:24.657000000',NULL,0,'/api/admin/v1/menusAdmin','DELETE','批量删除菜单',5,'后端菜单管理','sys_menusAdmin'),
(26,'apis-b1Aa1P',0,1,'null',1,1,'2022-10-06 07:43:16.518000000','2022-10-06 07:43:16.521000000',NULL,0,'/api/admin/v1/sysapiGroup','POST','新增Api分组',4,'Api分组管理','sys_apis_group'),
(27,'apis-Q8R53W',0,1,'null',1,1,'2022-10-06 07:43:51.885000000','2022-10-06 07:43:51.887000000',NULL,0,'/api/admin/v1/sysapiGroup/{id}','PUT','更新Api分组',4,'Api分组管理','sys_apis_group'),
(28,'apis-E1wq8r',0,1,'null',1,1,'2022-10-06 07:44:12.689000000','2022-10-06 07:44:12.691000000',NULL,0,'/api/admin/v1/sysapiGroup','DELETE','批量删除Api分组',4,'Api分组管理','sys_apis_group'),
(29,'apis-oVQo39',0,1,'null',1,1,'2022-10-06 07:44:41.519000000','2022-10-06 07:44:41.520000000',NULL,0,'/api/admin/v1/sysapiGroup','GET','批量获取Api分组',4,'Api分组管理','sys_apis_group'),
(30,'apis-Z3zdVp',0,1,'null',1,1,'2022-10-09 13:21:13.207000000','2022-10-09 13:21:13.208000000',NULL,0,'/api/admin/v1/sysapi-tree','GET','Api树',3,'API管理','sys_apis');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(31,'apis-z35M3K',0,1,'null',1,1,'2022-10-09 13:22:20.784000000','2022-10-09 13:22:20.786000000',NULL,0,'/api/admin/v1/menusAdmin-my','GET','个人后端菜单列表',5,'后端菜单管理','sys_menusAdmin'),
(32,'apis-A8ME8k',0,1,'null',1,1,'2022-10-09 13:23:18.091000000','2022-10-09 13:23:18.091000000',NULL,0,'/api/admin/v1/role-menus','POST','更新角色菜单授权',2,'角色管理','sys_role'),
(33,'apis-a1r21l',0,1,'null',1,1,'2022-10-09 13:23:49.208000000','2022-10-09 13:23:49.209000000',NULL,0,'/api/admin/v1/role-apis','POST','更新角色Api授权',2,'角色管理','sys_role'),
(34,'apis-6VDw8w',0,1,'null',1,1,'2022-10-09 13:24:31.034000000','2022-10-09 13:24:31.035000000',NULL,0,'/api/admin/v1/user-password','POST','更新密码',1,'用户管理','sys-user'),
(35,'apis-M8Jz8y',0,1,'null',1,1,'2022-12-17 15:35:25.382000000','2022-12-17 15:35:25.385000000',NULL,0,'/api/admin/v1/user-password','POST','重置密码',1,'用户管理','sys-user'),
(36,'apis-p8y71g',0,1,'null',1,1,'2022-12-24 11:36:01.697000000','2022-12-24 11:36:01.698000000',NULL,0,'/api/admin/v1/comment','POST','新增评论',8,'评论管理','sys_comment'),
(37,'apis-n8LR1r',0,1,'null',1,1,'2022-12-24 11:36:41.127000000','2022-12-24 11:36:41.128000000',NULL,0,'/api/admin/v1/comment/content','POST','更新评论内容',8,'评论管理','sys_comment'),
(38,'apis-D19pVe',0,1,'null',1,1,'2022-12-24 11:42:15.960000000','2022-12-24 11:42:15.961000000',NULL,0,'/api/admin/v1/comment/state','POST','更新评论状态',8,'评论管理','sys_comment'),
(39,'apis-lVgN8E',0,1,'null',1,1,'2022-12-24 11:42:59.681000000','2022-12-24 11:42:59.682000000',NULL,0,'/api/admin/v1/comment','DELETE','批量删除评论',8,'评论管理','sys_comment'),
(40,'apis-b3Px3X',0,1,'null',1,1,'2022-12-24 11:43:31.113000000','2022-12-24 11:43:31.114000000',NULL,0,'/api/admin/v1/comment','GET','获取评论列表',8,'评论管理','sys_comment');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(41,'apis-03GyVJ',0,1,'null',1,1,'2022-12-24 11:43:51.929000000','2022-12-24 11:43:51.930000000',NULL,0,'/api/admin/v1/comment/total','GET','获取评论总计',8,'评论管理','sys_comment'),
(42,'apis-787K8Y',0,1,'null',1,1,'2022-12-24 11:45:12.358000000','2022-12-24 11:45:12.359000000',NULL,0,'/api/admin/v1/siteConfig','POST','新增站点配置',7,'站点管理','sys_site_config'),
(43,'apis-5VWJ8J',0,1,'null',1,1,'2022-12-24 11:45:38.917000000','2022-12-24 11:45:38.917000000',NULL,0,'/api/admin/v1/siteConfig','PUT','更新站点配置',7,'站点管理','sys_site_config'),
(44,'apis-3OPl3g',0,1,'null',1,1,'2022-12-24 11:45:59.353000000','2022-12-24 11:45:59.355000000',NULL,0,'/api/admin/v1/siteConfig','GET','获取站点配置',7,'站点管理','sys_site_config'),
(45,'apis-8xx9VD',0,1,'null',1,1,'2022-12-24 12:28:50.435000000','2022-12-24 12:28:50.436000000',NULL,0,'/api/admin/v1/fileLibType','POST','新增媒体库',9,'媒体库管理','file-lib'),
(46,'apis-8aOB3P',0,1,'null',1,1,'2022-12-24 12:29:18.678000000','2022-12-24 12:29:18.679000000',NULL,0,'/api/admin/v1/fileLibType/{id}','PUT','更新媒体库',9,'媒体库管理','file-lib'),
(47,'apis-8Avr8b',0,1,'null',1,1,'2022-12-24 12:29:56.774000000','2022-12-24 12:29:56.776000000',NULL,0,'/api/admin/v1/fileLibType','GET','获取媒体库列表',9,'媒体库管理','file-lib'),
(48,'apis-VX2k1r',0,1,'null',1,1,'2022-12-24 12:30:34.820000000','2022-12-24 12:30:34.821000000',NULL,0,'/api/admin/v1/fileLibType','DELETE','批量删除媒体库',9,'媒体库管理','file-lib'),
(49,'apis-VJm48e',0,1,'null',1,1,'2022-12-24 12:31:35.511000000','2022-12-24 12:31:35.512000000',NULL,0,'/api/admin/v1/fileLibConfig','POST','新增媒体库配置',9,'媒体库管理','file-lib'),
(50,'apis-VAlv3B',0,1,'null',1,1,'2022-12-24 12:33:23.938000000','2022-12-24 12:33:23.943000000',NULL,0,'/api/admin/v1/fileLibConfig/{typeId}','GET','获取媒体库配置',9,'媒体库管理','file-lib');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(51,'apis-Ve701E',0,1,'null',1,1,'2022-12-24 12:34:21.628000000','2022-12-24 12:34:21.631000000',NULL,0,'/api/admin/v1/fileLib/byType/{typeId}','GET','批量获取媒体列表',9,'媒体库管理','file-lib'),
(52,'apis-32dG1O',0,1,'null',1,1,'2022-12-24 12:56:34.850000000','2022-12-24 12:56:34.851000000',NULL,0,'/api/admin/v1/link','POST','新增友链',10,'友链管理','link-friends'),
(53,'apis-1J6mVB',0,1,'null',1,1,'2022-12-24 12:57:05.342000000','2022-12-24 12:57:05.343000000',NULL,0,'/api/admin/v1/link/{id}','PUT','更新友链',10,'友链管理','link-friends'),
(54,'apis-VR4Z35',0,1,'null',1,1,'2022-12-24 12:57:39.791000000','2022-12-24 12:57:39.791000000',NULL,0,'/api/admin/v1/link','DELETE','批量删除友链',10,'友链管理','link-friends'),
(55,'apis-Vzo68k',0,1,'null',1,1,'2022-12-24 12:58:33.324000000','2022-12-24 12:58:33.325000000',NULL,0,'/api/admin/v1/link','GET','友链列表',10,'友链管理','link-friends'),
(56,'apis-VpzX3a',0,1,'null',1,1,'2022-12-24 13:04:56.935000000','2022-12-24 13:04:56.936000000',NULL,0,'/api/admin/v1/shortLink','POST','新增短链接',11,'短链接管理','link-shorts'),
(57,'apis-VmEP87',0,1,'null',1,1,'2022-12-24 13:05:54.049000000','2022-12-24 13:05:54.051000000',NULL,0,'/api/admin/v1/shortLink/{id}','PUT','更新短链接',11,'短链接管理','link-shorts'),
(58,'apis-3pZY1q',0,1,'null',1,1,'2022-12-24 13:08:44.320000000','2022-12-24 13:10:02.851000000',NULL,0,'/api/admin/v1/shortLink/{id}','DELETE','删除短链接',11,'短链接管理','link-shorts'),
(59,'apis-8mvD1D',0,1,'null',1,1,'2022-12-24 13:09:47.452000000','2022-12-24 13:09:47.453000000',NULL,0,'/api/admin/v1/shortLink','DELETE','批量删除短链接',11,'短链接管理','link-shorts'),
(60,'apis-3ydL3p',0,1,'null',1,1,'2022-12-24 13:10:42.375000000','2022-12-24 13:10:42.375000000',NULL,0,'/api/admin/v1/shortLink','GET','短链接列表',11,'短链接管理','link-shorts');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(61,'apis-15nO3l',0,1,'null',1,1,'2022-12-24 13:25:02.337000000','2022-12-24 13:25:02.338000000',NULL,0,'/api/admin/v1/menusAgent','POST','新增导航分类',13,'导航分类管理','link-menus-agent'),
(62,'apis-8pgg8v',0,1,'null',1,1,'2022-12-24 13:25:49.674000000','2022-12-24 13:25:49.675000000',NULL,0,'/api/admin/v1/menusAgent/{id}','PUT','更新导航分类',13,'导航分类管理','link-menus-agent'),
(63,'apis-8RJWVb',0,1,'null',1,1,'2022-12-24 13:26:58.047000000','2022-12-24 13:26:58.048000000',NULL,0,'/api/admin/v1/menusAgent','DELETE','批量删除导航分类',13,'导航分类管理','link-menus-agent'),
(64,'apis-82MA8M',0,1,'null',1,1,'2022-12-24 13:27:20.964000000','2022-12-24 13:27:20.965000000',NULL,0,'/api/admin/v1/menusAgent','GET','导航分类列表',13,'导航分类管理','link-menus-agent'),
(65,'apis-1LbQV4',0,1,'null',1,1,'2022-12-24 13:28:18.706000000','2022-12-24 13:28:18.707000000',NULL,0,'/api/admin/v1/menus','POST','新增导航',12,'导航管理','link-menus'),
(66,'apis-3Wve14',0,1,'null',1,1,'2022-12-24 13:28:38.624000000','2022-12-24 13:28:38.626000000',NULL,0,'/api/admin/v1/menus/{id}','PUT','更新导航',12,'导航管理','link-menus'),
(67,'apis-8Ybj8M',0,1,'null',1,1,'2022-12-24 13:28:56.229000000','2022-12-24 13:28:56.230000000',NULL,0,'/api/admin/v1/menus','DELETE','批量删除导航',12,'导航管理','link-menus'),
(68,'apis-VyjbVy',0,1,'null',1,1,'2022-12-24 13:29:17.848000000','2022-12-24 13:29:17.849000000',NULL,0,'/api/admin/v1/menus','GET','导航列表',12,'导航管理','link-menus'),
(69,'apis-VJYn1l',0,1,'null',1,1,'2022-12-24 13:46:13.823000000','2022-12-24 13:46:13.868000000',NULL,0,'/api/admin/v1/tags','POST','新增标签',14,'标签管理','blog-tags'),
(70,'apis-8Aoa1P',0,1,'null',1,1,'2022-12-24 13:46:32.902000000','2022-12-24 13:46:32.903000000',NULL,0,'/api/admin/v1/tags/{id}','PUT','更新标签',14,'标签管理','blog-tags');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(71,'apis-8Rz53W',0,1,'null',1,1,'2022-12-24 13:46:54.008000000','2022-12-24 13:46:54.010000000',NULL,0,'/api/admin/v1/tags','DELETE','批量删除标签',14,'标签管理','blog-tags'),
(72,'apis-1wOq8r',0,1,'null',1,1,'2022-12-24 13:47:14.344000000','2022-12-24 13:47:14.346000000',NULL,0,'/api/admin/v1/tags','GET','标签列表',14,'标签管理','blog-tags'),
(73,'apis-1QBo39',0,1,'null',1,1,'2022-12-24 13:54:08.597000000','2022-12-24 13:54:08.610000000',NULL,0,'/api/admin/v1/category','POST','新增分类',15,'分类管理','blog-category'),
(74,'apis-1zedVp',0,1,'null',1,1,'2022-12-24 13:55:13.672000000','2022-12-24 13:55:13.681000000',NULL,0,'/api/admin/v1/category/{id}','PUT','更新分类',15,'分类管理','blog-category'),
(75,'apis-15PM3K',0,1,'null',1,1,'2022-12-24 13:56:57.945000000','2022-12-24 13:56:57.946000000',NULL,0,'/api/admin/v1/category','DELETE','批量删除分类',15,'分类管理','blog-category'),
(76,'apis-3M0E8k',0,1,'null',1,1,'2022-12-24 13:57:13.563000000','2022-12-24 13:57:13.565000000',NULL,0,'/api/admin/v1/category','GET','分类列表',15,'分类管理','blog-category'),
(77,'apis-Vrm21l',0,1,'null',1,1,'2022-12-24 13:59:22.429000000','2022-12-24 13:59:22.431000000',NULL,0,'/api/admin/v1/article','POST','新增文章',16,'文章管理','blog-article'),
(78,'apis-3Dbw8w',0,1,'null',1,1,'2022-12-24 13:59:46.238000000','2022-12-24 13:59:46.240000000',NULL,0,'/api/admin/v1/article/{id}','PUT','更新文章',16,'文章管理','blog-article'),
(79,'apis-3Jez8y',0,1,'null',1,1,'2022-12-24 14:00:09.654000000','2022-12-24 14:00:09.656000000',NULL,0,'/api/admin/v1/article','DELETE','批量删除文章',16,'文章管理','blog-article'),
(80,'apis-1y771g',0,1,'null',1,1,'2022-12-24 14:00:28.160000000','2022-12-24 14:00:28.161000000',NULL,0,'/api/admin/v1/article/{id}','GET','文章详情',16,'文章管理','blog-article');
INSERT INTO qycmsnew.qy_sys_api (id,instance_id,sort,status_flag,extend_shadow,created_by,updated_by,created_at,updated_at,deleted_at,deleted_flag,`path`,`method`,description,group_id,api_group,identifier) VALUES
(81,'apis-1LNR1r',0,1,'null',1,1,'2022-12-24 14:00:53.357000000','2022-12-24 14:00:53.358000000',NULL,0,'/api/admin/v1/article/initPermaLink','POST','生成文章链接',16,'文章管理','blog-article'),
(82,'apis-19YpVe',0,1,'null',1,1,'2022-12-24 14:01:12.443000000','2022-12-24 14:01:12.444000000',NULL,0,'/api/admin/v1/article','GET','文章列表',16,'文章管理','blog-article');
-- 权限信息表
INSERT INTO  qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'g','U_admin','R_1','*','','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user/{id}','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user/{id}','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user/myInfo','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role/{id}','DELETE','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role/{id}','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi/{id}','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi/{id}','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin','POST','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin/{id}','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin/{id}','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis_group','/api/admin/v1/sysapiGroup','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis_group','/api/admin/v1/sysapiGroup/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis_group','/api/admin/v1/sysapiGroup','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis_group','/api/admin/v1/sysapiGroup','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_apis','/api/admin/v1/sysapi-tree','GET','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_menusAdmin','/api/admin/v1/menusAdmin-my','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role-menus','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_role','/api/admin/v1/role-apis','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user-password','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys-user','/api/admin/v1/user-password','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment/content','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment/state','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment','GET','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_comment','/api/admin/v1/comment/total','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_site_config','/api/admin/v1/siteConfig','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_site_config','/api/admin/v1/siteConfig','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','sys_site_config','/api/admin/v1/siteConfig','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibType','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibType/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibType','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibType','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibConfig','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLibConfig/{typeId}','GET','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','file-lib','/api/admin/v1/fileLib/byType/{typeId}','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-friends','/api/admin/v1/link','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-friends','/api/admin/v1/link/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-friends','/api/admin/v1/link','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-friends','/api/admin/v1/link','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-shorts','/api/admin/v1/shortLink','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-shorts','/api/admin/v1/shortLink/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-shorts','/api/admin/v1/shortLink/{id}','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-shorts','/api/admin/v1/shortLink','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-shorts','/api/admin/v1/shortLink','GET','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus','/api/admin/v1/menus','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus','/api/admin/v1/menus/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus','/api/admin/v1/menus','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus','/api/admin/v1/menus','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus-agent','/api/admin/v1/menusAgent','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus-agent','/api/admin/v1/menusAgent/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus-agent','/api/admin/v1/menusAgent','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','link-menus-agent','/api/admin/v1/menusAgent','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-tags','/api/admin/v1/tags','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-tags','/api/admin/v1/tags/{id}','PUT','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-tags','/api/admin/v1/tags','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-tags','/api/admin/v1/tags','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-category','/api/admin/v1/category','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-category','/api/admin/v1/category/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-category','/api/admin/v1/category','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-category','/api/admin/v1/category','GET','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article/{id}','PUT','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article','DELETE','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article/{id}','GET','','');
INSERT INTO qy_sys_casbin_rule (instance_id,sort,status_flag,extend_shadow,created_at,updated_at,disabled_at,deleted_at,ptype,v0,v1,v2,v3,v4,v5) VALUES
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article/initPermaLink','POST','',''),
('def-0',0,0,NULL,NULL,NULL,NULL,NULL,'p','R_1','blog-article','/api/admin/v1/article','GET','','');
-- 后端菜单
INSERT INTO qy_sys_menus_admin (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,breadcrumbName,identifier,parent_id,icon,mtype,`path`,component,`level`,redirect,created_by,updated_by) VALUES
(1,'menus-admin-nXPgLP',1,1,'null','2022-10-04 05:37:22.593000000','2022-10-04 06:30:49.699000000',NULL,0,'仪表盘','仪表盘','dashboard',0,'DashboardOutlined',1,'/dashboard','/views/dashboard/Workplace.vue',0,NULL,1,1),
(2,'menus-admin-B53YmJ',6,1,'null','2022-10-04 06:07:40.925000000','2022-10-10 14:03:52.302000000',NULL,0,'系统管理','系统管理','system',0,'SettingOutlined',1,'/system','',0,'/system/user/list',1,1),
(3,'menus-admin-20qGr3',1,1,'null','2022-10-04 06:10:59.722000000','2022-10-06 07:10:34.334000000',NULL,0,'用户管理','用户管理','user',2,'UserOutlined',1,'/system/user','',0,'/system/user/list',1,1),
(4,'menus-admin-jmkRrK',1,1,'null','2022-10-04 09:01:11.973000000','2022-10-04 09:03:20.951000000',NULL,0,'用户列表','用户列表','user-list',3,'UserOutlined',2,'/system/user/list','/views/user/index.vue',0,NULL,1,1),
(5,'menus-admin-E5Po0o',2,1,'null','2022-10-04 09:02:54.870000000','2022-10-04 09:02:54.870000000',NULL,0,'新增用户','新增用户','user-add',3,'UserAddOutlined',2,'/system/user/add','/views/user/add.vue',0,NULL,1,1),
(6,'menus-admin-Z53V0x',3,1,'null','2022-10-04 09:04:53.602000000','2022-10-04 09:04:53.602000000',NULL,0,'编辑用户','编辑用户','user-update',3,'UserSwitchOutlined',2,'/system/user/:id','/views/user/_id.vue',0,NULL,1,1),
(7,'menus-admin-l56v5X',2,1,'null','2022-10-04 16:53:24.230000000','2022-10-06 07:10:59.341000000',NULL,0,'角色管理','角色管理','role',2,'TeamOutlined',1,'/system/role','',0,'/system/role/list',1,1),
(8,'menus-admin-lmwA57',3,1,'null','2022-10-04 16:55:23.900000000','2022-10-06 07:11:42.444000000',NULL,0,'菜单管理','菜单管理','menusAdmin',2,'MenuOutlined',1,'/system/menusAdmin','',0,'/system/menusAdmin/list',1,1),
(9,'menus-admin-OmzP5D',0,1,'null','2022-10-04 16:56:47.631000000','2022-10-04 16:56:47.632000000',NULL,0,'角色列表','角色列表','role-list',7,'TeamOutlined',2,'/system/role/list','/views/role/index.vue',0,NULL,1,1),
(10,'menus-admin-wrXj0o',1,1,'null','2022-10-04 16:59:00.508000000','2022-10-04 16:59:00.509000000',NULL,0,'菜单列表','菜单列表','menusAdmin-list',8,'MenuFoldOutlined',2,'/system/menusAdmin/list','/views/menusAdmin/index.vue',0,NULL,1,1);
INSERT INTO qy_sys_menus_admin (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,breadcrumbName,identifier,parent_id,icon,mtype,`path`,component,`level`,redirect,created_by,updated_by) VALUES
(11,'menus-admin-O0Aa0L',2,1,'null','2022-10-04 17:01:44.327000000','2022-10-04 17:01:44.329000000',NULL,0,'新增菜单','新增菜单','menusAdmin-add',8,'BarsOutlined',2,'/system/menusAdmin/add','/views/menusAdmin/add.vue',0,NULL,1,1),
(12,'menus-admin-20l3mw',3,1,'null','2022-10-04 17:02:52.994000000','2022-10-04 17:02:52.995000000',NULL,0,'编辑菜单','编辑菜单','menusAdmin-update',8,'AlignRightOutlined',2,'/system/menusAdmin/:id','/views/menusAdmin/_id.vue',0,NULL,1,1),
(13,'menus-admin-RrQW0K',4,1,'null','2022-10-04 17:07:38.870000000','2022-10-06 07:12:22.798000000',NULL,0,'Api管理','Api管理','apis',2,'ApiOutlined',1,'/system/apis','',0,'/system/apis/list',1,1),
(14,'menus-admin-7mYB0Z',1,1,'null','2022-10-04 17:09:46.845000000','2022-10-04 17:09:46.847000000',NULL,0,'Api列表','Api列表','apis-list',13,'AppstoreAddOutlined',2,'/system/apis/list','/views/apis/index.vue',0,NULL,1,1),
(15,'menus-admin-A0ygra',2,1,'null','2022-10-04 17:15:07.704000000','2022-10-04 17:15:07.706000000',NULL,0,'新增Api','新增Api','apis-add',13,'HeatMapOutlined',2,'/system/apis/add','/views/apis/add.vue',0,NULL,1,1),
(16,'menus-admin-VmK6rB',3,1,'null','2022-10-04 17:16:00.213000000','2022-10-04 17:16:00.215000000',NULL,0,'编辑Api','编辑Api','apis-update',13,'AppstoreOutlined',2,'/system/apis/:id','/views/apis/_id.vue',0,NULL,1,1),
(17,'menus-admin-6ryEro',5,1,'null','2022-10-04 17:23:12.209000000','2022-10-04 17:23:12.211000000',NULL,0,'站点管理','站点管理','site',2,'ApartmentOutlined',1,'/system/site','/views/site/index.vue',0,NULL,1,1),
(18,'menus-admin-v0z4rd',3,1,'null','2022-10-10 14:02:16.393000000','2022-11-12 11:13:44.084000000',NULL,0,'媒体管理','媒体管理','fileLibs',0,'FolderOutlined',1,'/fileLibs','',0,'/fileLib/list',1,1),
(19,'menus-admin-M5nLmv',1,1,'null','2022-10-10 14:06:12.157000000','2022-10-24 15:05:31.974000000',NULL,0,'媒体库','媒体库','fileLib',18,'CloudUploadOutlined',1,'/fileLib','',0,'/fileLib/list',1,1),
(20,'menus-admin-d59ZrV',2,1,'null','2022-10-10 14:16:58.306000000','2022-10-10 14:54:46.588000000',NULL,0,'媒体库配置','媒体库配置','fileLib-config',18,'FolderAddOutlined',1,'/fileLib/config','',0,'/fileLib/config/list',1,1);
INSERT INTO qy_sys_menus_admin (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,breadcrumbName,identifier,parent_id,icon,mtype,`path`,component,`level`,redirect,created_by,updated_by) VALUES
(21,'menus-admin-Kmznmg',1,1,'null','2022-10-10 14:25:22.361000000','2022-10-10 16:41:34.222000000',NULL,0,'媒体库列表','媒体库列表','fileLib-list',19,'FileProtectOutlined',2,'/fileLib/list','/views/fileLib/index.vue',0,'',1,1),
(22,'menus-admin-V0N8r4',1,1,'null','2022-10-10 14:56:39.277000000','2022-10-10 14:56:39.278000000',NULL,0,'媒体库配置列表','媒体库配置列表','fileLib-config-list',20,'FileAddOutlined',2,'/fileLib/config/list','/views/fileLibConfig/index.vue',0,'',1,1),
(23,'menus-admin-jr7q0v',2,1,'null','2022-10-10 16:34:49.231000000','2022-10-10 16:34:49.232000000',NULL,0,'编辑媒体库配置','编辑媒体库配置','fileLib-config-update',20,'FileAddOutlined',2,'/fileLib/config/:id','/views/fileLibConfig/_id.vue',0,'',1,1),
(24,'menus-admin-lrjd0v',4,1,'null','2022-11-06 13:35:32.357000000','2022-11-12 11:17:53.999000000',NULL,0,'链接管理','链接管理','links',0,'LinkOutlined',1,'/links','',0,'/links/menus',1,1),
(25,'menus-admin-z58l54',1,1,'null','2022-11-06 13:39:25.527000000','2022-11-06 13:39:25.529000000',NULL,0,'导航管理','导航管理','menus',24,'ApartmentOutlined',1,'/links/menus','',0,'/links/menus/list',1,1),
(26,'menus-admin-80Bemk',1,1,'null','2022-11-06 13:42:15.139000000','2022-11-06 13:42:15.163000000',NULL,0,'导航列表','导航列表','menus-list',25,'ApartmentOutlined',2,'/links/menus/list','/views/menus/index.vue',0,'',1,1),
(27,'menus-admin-p5ZM0k',2,1,'null','2022-11-06 13:44:07.685000000','2022-11-06 13:44:07.699000000',NULL,0,'导航配置','导航配置','menus-update',25,'ApartmentOutlined',2,'/links/menus/:id','/views/menus/_id.vue',0,'',1,1),
(28,'menus-admin-V5Y7m6',2,1,'null','2022-11-06 13:48:10.719000000','2022-11-06 13:48:26.223000000',NULL,0,'友链管理','友链管理','link',24,'NodeIndexOutlined',1,'/links/link','',0,'/links/link/list',1,1),
(29,'menus-admin-Dr92mJ',3,1,'null','2022-11-06 13:50:55.983000000','2022-11-06 13:50:55.986000000',NULL,0,'短连接管理','短连接管理','shortLink',24,'PaperClipOutlined',1,'/links/shortLink','',0,'/links/shortLink/list',1,1),
(30,'menus-admin-M0jb0Y',1,1,'null','2022-11-06 13:52:25.568000000','2022-11-06 13:52:25.569000000',NULL,0,'友联列表','友联列表','link-list',28,'NodeIndexOutlined',2,'/links/link/list','/views/link/index.vue',0,'',1,1);
INSERT INTO qy_sys_menus_admin (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,breadcrumbName,identifier,parent_id,icon,mtype,`path`,component,`level`,redirect,created_by,updated_by) VALUES
(31,'menus-admin-PmxQ0K',3,1,'null','2022-11-06 13:55:30.637000000','2022-11-06 13:55:30.638000000',NULL,0,'短连接列表','短连接列表','shortLink-list',29,'PaperClipOutlined',2,'/links/shortLink/list','/views/shortLink/index.vue',0,'',1,1),
(32,'menus-admin-15Q15v',2,1,'null','2022-11-12 11:16:14.234000000','2022-11-12 11:24:10.413000000',NULL,0,'文章系统','文章系统','blog',0,'ReadOutlined',1,'/blog','',0,'/blog/article',1,1),
(33,'menus-admin-p0OzmN',1,1,'null','2022-11-12 11:22:11.319000000','2022-11-12 11:22:11.321000000',NULL,0,'文章管理','文章管理','article',32,'EditOutlined',1,'/blog/article','',0,'/blog/article/list',1,1),
(34,'menus-admin-xm19mP',1,1,'null','2022-11-12 11:25:35.569000000','2022-11-12 11:25:35.571000000',NULL,0,'文章列表','文章列表','article-list',33,'EditOutlined',2,'/blog/article/list','/views/article/index.vue',0,'',1,1),
(35,'menus-admin-B0aOr8',2,1,'null','2022-11-12 11:27:03.555000000','2022-11-12 11:30:39.271000000',NULL,0,'新增文章','新增文章','article-add',33,'EditOutlined',2,'/blog/article/add','/views/article/add.vue',0,'',1,1),
(36,'menus-admin-8m7kmb',3,1,'null','2022-11-12 11:30:01.991000000','2022-11-12 11:30:01.993000000',NULL,0,'编辑文章','编辑文章','article-update',33,'EditOutlined',2,'/blog/article/:id','/views/article/_id.vue',0,'',1,1),
(37,'menus-admin-bm6w0j',2,1,'null','2022-11-12 11:32:21.830000000','2022-11-12 11:32:21.831000000',NULL,0,'页面管理','页面管理','page',32,'SnippetsOutlined',1,'/blog/page','',0,'/blog/page/list',1,1),
(38,'menus-admin-qrXpr4',1,1,'null','2022-11-12 11:34:03.124000000','2022-11-12 11:34:03.126000000',NULL,0,'页面列表','页面列表','page-list',37,'SnippetsOutlined',2,'/blog/page/list','/views/page/index.vue',0,'',1,1),
(39,'menus-admin-Mm4N5j',2,1,'null','2022-11-12 11:35:09.252000000','2022-11-12 11:35:09.253000000',NULL,0,'新增页面','新增页面','page-add',37,'SnippetsOutlined',2,'/blog/page/add','/views/page/add.vue',0,'',1,1),
(40,'menus-admin-VmWx02',3,1,'null','2022-11-12 11:37:04.126000000','2022-11-12 11:37:04.126000000',NULL,0,'编辑页面','编辑页面','page-update',37,'SnippetsOutlined',2,'/blog/page/:id','/views/page/_id.vue',0,'',1,1);
INSERT INTO qy_sys_menus_admin (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,name,breadcrumbName,identifier,parent_id,icon,mtype,`path`,component,`level`,redirect,created_by,updated_by) VALUES
(41,'menus-admin-erAyr3',3,1,'null','2022-11-12 11:41:43.526000000','2022-11-12 11:41:43.528000000',NULL,0,'分类管理','分类管理','category',32,'BlockOutlined',1,'/blog/category','',0,'/blog/category/list',1,1),
(42,'menus-admin-Q5kK5Z',1,1,'null','2022-11-12 11:43:22.217000000','2022-11-12 11:43:22.220000000',NULL,0,'分类列表','分类列表','category-list',41,'BlockOutlined',2,'/blog/category/list','/views/category/index.vue',0,'',1,1),
(43,'menus-admin-YmLJ5E',4,1,'null','2022-11-12 11:46:59.399000000','2022-11-12 11:46:59.399000000',NULL,0,'标签管理','标签管理','tags',32,'TagsOutlined',1,'/blog/tags','',0,'/blog/tags/list',1,1),
(44,'menus-admin-mq7D0p',1,1,'null','2022-11-12 11:48:37.773000000','2022-11-12 11:48:37.774000000',NULL,0,'标签列表','标签列表','tags-list',43,'TagsOutlined',2,'/blog/tags/list','/views/tags/index.vue',0,'',1,1),
(45,'menus-admin-mNNX54',5,1,'null','2022-11-12 11:52:18.517000000','2022-11-12 11:52:18.518000000',NULL,0,'评论管理','评论管理','comment',32,'CommentOutlined',1,'/blog/comment','',0,'/blog/comment/list',1,1),
(46,'menus-admin-53OYmJ',1,1,'null','2022-11-12 11:53:26.874000000','2022-11-12 11:53:26.875000000',NULL,0,'评论列表','评论列表','comment-list',45,'CommentOutlined',2,'/blog/comment/list','/views/comment/index.vue',0,'',1,1);

-- 角色-api
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,1,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,2,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,3,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,4,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,5,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,6,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,7,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,8,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,9,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,10,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,11,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,12,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,13,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,14,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,15,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,16,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,17,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,18,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,19,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,20,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,21,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,22,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,23,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,24,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,25,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,26,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,27,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,28,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,29,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,30,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,31,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,32,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,33,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,34,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,35,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,36,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,37,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,38,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,39,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,40,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,41,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,42,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,43,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,44,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,45,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,46,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,47,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,48,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,49,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,50,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,51,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,52,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,53,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,54,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,55,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,56,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,57,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,58,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,59,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,60,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,61,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,62,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,63,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,64,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,65,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,66,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,67,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,68,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,69,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,70,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,71,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,72,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,73,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,74,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,75,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,76,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,77,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,78,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,79,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,80,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
INSERT INTO qy_sys_role_api (role_id,api_id,created_at,updated_at) VALUES
(1,81,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000'),
(1,82,'2022-12-24 14:02:53.127000000','2022-12-24 14:02:53.127000000');
-- 角色-菜单
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(1,1,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,2,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,3,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,4,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,5,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,6,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,7,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,8,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,9,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,10,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000');
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(1,11,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,12,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,13,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,14,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,15,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,16,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,17,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,18,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,19,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,20,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000');
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(1,21,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,22,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,23,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,24,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,25,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,26,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,27,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,28,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,29,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,30,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000');
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(1,31,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,32,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,33,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,34,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,35,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,36,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,37,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,38,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,39,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,40,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000');
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(1,41,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,42,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,43,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,44,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,45,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(1,46,'2022-11-12 11:54:03.530000000','2022-11-12 11:54:03.530000000'),
(2,1,'2022-10-07 17:12:10.342000000','2022-10-07 17:12:10.342000000'),
(2,4,'2022-10-07 17:12:10.342000000','2022-10-07 17:12:10.342000000'),
(3,1,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000'),
(3,3,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000');
INSERT INTO qy_sys_role_menus (role_id,menus_id,created_at,updated_at) VALUES
(3,4,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000'),
(3,5,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000'),
(3,6,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000'),
(3,7,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000'),
(3,9,'2022-10-07 14:48:53.556000000','2022-10-07 14:48:53.556000000');

-- 站点配置
INSERT INTO qy_sys_site_config (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,config_key,config_value,config_name,config_tip,ftype,created_by,updated_by) VALUES
(1,'site-config-oQ01Pw',0,1,'null','2022-10-30 15:58:34.591000000','2022-12-12 17:33:43.245000000',NULL,0,'site_name','Windcoder','站点名称','',1,1,1),
(2,'site-config-KrYErv',0,1,'null','2022-10-30 16:02:34.243000000','2022-12-12 17:33:43.252000000',NULL,0,'site_url','http://test.windcoder.com/','站点地址','',1,1,1),
(3,'site-config-XQ6wPz',0,1,'null','2022-10-30 16:03:06.576000000','2022-12-12 17:33:43.254000000',NULL,0,'site_key','Java,Web,前端,编程','关键词','主页关键词元信息，SEO友好(多个关键词以英文逗号分隔)',1,1,1),
(4,'site-config-kPjzPv',0,1,'null','2022-10-30 16:03:56.918000000','2022-12-12 17:33:43.256000000',NULL,0,'site_description','专注web前端与java程序开发，与web和Java开发爱好者分享素材与案例','站点描述','主页描述元信息，SEO友好',1,1,1),
(5,'site-config-vrJZrn',0,1,'null','2022-10-30 16:04:22.671000000','2022-12-12 17:33:43.258000000',NULL,0,'site_icon','https://windcoder.com/wp-content/uploads/2015/08/favicon-1.gif','Favicon','站点 ICON 图标',1,1,1),
(6,'site-config-BPljRG',0,1,'null','2022-10-30 16:04:50.469000000','2022-12-12 17:33:43.260000000',NULL,0,'site_logo','https://windcoder.com/wp-content/uploads/2017/02/logo_vift.png','网站Logo','请上传PNG图片作为网站Logo',1,1,1),
(7,'site-config-gRe6Qj',0,1,'null','2022-10-30 16:05:25.004000000','2022-12-12 17:33:43.263000000',NULL,0,'site_small_logo','https://windcoder.com/wp-content/uploads/2017/02/logo_vif_small.png','网站小Logo','请提供PNG图片作为网站小Logo',1,1,1),
(8,'site-config-JQ1grw',0,1,'null','2022-10-30 16:07:37.441000000','2022-12-12 17:33:43.266000000',NULL,0,'site_default_file_lib','2','文件保存位置','用于选择图片等上传到本地还是七牛等第三方平台',1,1,1),
(9,'site-config-aPqWRz',0,1,'null','2022-10-30 16:08:08.461000000','2022-12-12 17:33:43.269000000',NULL,0,'site_beian','冀ICP备14001026号','页脚备案文字','工信部域名备案号码',1,1,1),
(10,'site-config-4Qk3QD',0,1,'null','2022-10-30 16:08:32.197000000','2022-12-12 17:33:43.271000000',NULL,0,'site_open_date','2013-07-17','建站日期','网站开放的日期, 使用`YYYY-mm-dd`格式',1,1,1);
INSERT INTO qy_sys_site_config (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,config_key,config_value,config_name,config_tip,ftype,created_by,updated_by) VALUES
(11,'site-config-VRm9Rk',0,1,'null','2022-10-30 16:09:07.108000000','2022-12-12 17:33:43.273000000',NULL,0,'site_comment_flag','1','允许评论','站点评论总开关',2,1,1),
(12,'site-config-aP6dRd',0,1,'null','2022-10-30 16:09:32.650000000','2022-12-12 17:33:43.275000000',NULL,0,'site_home_undisplay_cats','','不显示的分类ID列表','分类ID数字之间用英文逗号分隔, 如果留空将展示所有分类到文章',2,1,1),
(13,'site-config-Nrjbr8',0,1,'null','2022-12-10 20:31:04.118000000','2022-12-12 17:33:43.277000000',NULL,0,'site_head_script_code','console.log("你好");','页头自定义JavaScript代码','页面头部加载的自定义代码，位于head标签结束前',2,1,1),
(14,'site-config-bQxlQO',0,1,'null','2022-10-30 16:10:29.110000000','2022-12-12 17:33:43.280000000',NULL,0,'site_head_code','body{
background:#f5f6f8
}
.content .markdown-body h3, .content .markdown-body h4, .content .markdown-body h5 {
padding-left: 10px;
border-left: 2px solid #f1c40f;
line-height: 1;
font-weight: 700;
margin-bottom: 20px;
}
.content  .markdown-body details{
background:#faad14;
}


','页头自定义CSS代码','页面头部加载的自定义代码，位于head标签结束前',2,1,1),
(15,'site-config-1PD7Q4',0,1,'null','2022-10-30 16:10:43.721000000','2022-12-12 17:33:43.283000000',NULL,0,'site_foot_code','

<!--百度统计-->
<script>
var _hmt = _hmt || [];
(function() {
var hm = document.createElement("script");
hm.src = "https://hm.baidu.com/hm.js?9a74e514dc0118daedd178f4c8da0201";
var s = document.getElementsByTagName("script")[0];
s.parentNode.insertBefore(hm, s);
})();
</script>
<script type="text/javascript" src="//js.users.51.la/20525743.js"></script>
<!--百度自动推送-->
<script>

(function(){
var bp = document.createElement(''script'');
var curProtocol = window.location.protocol.split('':'')[0];
if (curProtocol === ''https'') {
bp.src = ''https://zz.bdstatic.com/linksubmit/push.js'';
}
else {
bp.src = ''http://push.zhanzhang.baidu.com/push.js'';
}
var s = document.getElementsByTagName("script")[0];
s.parentNode.insertBefore(bp, s);
})();

</script>
<!--百度自动推送End-->




','页脚自定义代码','页面底部加载的自定义代码，位于body标签结束前',2,1,1),
(16,'site-config-ErXDr8',0,1,'null','2022-10-30 16:10:57.683000000','2022-12-12 17:33:43.286000000',NULL,0,'site_sider_code','<div><p>当前处于试运行期间，可能存在不稳定情况，敬请见谅。</p>
<a href="https://wj.qq.com/s2/7758573/07b0/" target="_blank">欢迎点击此处反馈访问过程中出现的问题</a></div>','侧边栏自定义代码','侧边栏加载的自定义代码，位于body标签结束前',2,1,1),
(17,'site-config-mrkVPe',0,1,'null','2022-10-30 16:11:29.487000000','2022-12-12 17:33:43.289000000',NULL,0,'site_desc','','个人描述','个人一句话说明',3,1,1),
(18,'site-config-vQpoRw',0,1,'null','2022-10-30 16:11:54.531000000','2022-12-12 17:33:43.292000000',NULL,0,'site_qq','2641914215','网站QQ','站点服务专属QQ号码',3,1,1),
(19,'site-config-dRJOQk',0,1,'null','2022-10-30 16:12:13.614000000','2022-12-12 17:33:43.294000000',NULL,0,'site_qq_group','8d9f2d4dbab7b99146c96f43a01962e56f6dcaa99ab1d350dfde46e2b41f31e1','网站QQ群 ID Key','站点专属服务QQ群加群链接的ID Key, 非群号, 至`http://shang.qq.com`获取',3,1,1),
(20,'site-config-VPOaR6',0,1,'null','2022-10-30 16:12:36.267000000','2022-12-12 17:33:43.296000000',NULL,0,'site_weibo','52xj','网站微博名','站点服务专属微博用户名',3,1,1);
INSERT INTO qy_sys_site_config (id,instance_id,sort,status_flag,extend_shadow,created_at,updated_at,deleted_at,deleted_flag,config_key,config_value,config_name,config_tip,ftype,created_by,updated_by) VALUES
(21,'site-config-NP62Pk',0,1,'null','2022-10-30 16:12:56.513000000','2022-12-12 17:33:43.300000000',NULL,0,'site_weixin_qr','https://windcoder.com/wp-content/uploads/2019/03/qrcode_weixin.jpg','网站微信','站点服务专属微信号的二维码图片(可以是公众号二维码)',3,1,1),
(22,'site-config-7PyePK',0,1,'null','2022-10-30 16:13:20.722000000','2022-12-12 17:33:43.303000000',NULL,0,'site_mailme_id','windcoderz@foxmail.com','QQ邮我按钮ID','QQ邮我链接内的ID字段, 访问`http://open.mail.qq.com`获取',3,1,1),
(23,'site-config-vQK5ra',0,1,'null','2022-10-30 16:13:48.563000000','2022-12-12 17:33:43.305000000',NULL,0,'site_pay_flag','2','启用赞赏','赞赏开关，开启后展示收款二维码',4,1,1),
(24,'site-config-pQvkPw',0,1,'null','2022-10-30 16:14:12.582000000','2022-12-12 17:33:43.307000000',NULL,0,'site_weixin_pay_qr','','网站微信收款二维码','用于网站收集打赏等的微信收款二维码图片(赞助站长小工具等使用)',4,1,1),
(25,'site-config-5PV0R1',0,1,'null','2022-10-30 16:14:29.770000000','2022-12-12 17:33:43.310000000',NULL,0,'site_alipay_pay_qr','','网站支付宝收款二维码','用于网站收集打赏等的支付宝收款二维码图片(赞助站长小工具等使用)',4,1,1);
