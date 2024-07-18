/*
 Navicat Premium Data Transfer

 Source Server         : me
 Source Server Type    : MySQL
 Source Server Version : 80036 (8.0.36)
 Source Host           : localhost:3306
 Source Schema         : micro_user

 Target Server Type    : MySQL
 Target Server Version : 80036 (8.0.36)
 File Encoding         : 65001

 Date: 18/07/2024 23:20:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int NOT NULL COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`config_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '初始化密码 123456', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', '深色主题theme-dark，浅色主题theme-light', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (4, '账号自助-验证码开关', 'sys.account.captchaEnabled', 'true', 'Y', '是否开启验证码功能（true开启，false关闭）', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', '是否开启注册用户功能（true开启，false关闭）', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (6, '用户登录-黑名单列表', 'sys.login.blackIPList', '', 'Y', '设置登录IP黑名单限制，多个匹配项以;分隔，支持匹配（*通配、网段）', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL COMMENT '部门id',
  `parent_id` bigint DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '部门名称',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (100, 0, '0', '若依科技', 0, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (101, 100, '0,100', '深圳总公司', 1, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (102, 100, '0,100', '长沙分公司', 2, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (103, 101, '0,100,101', '研发部门', 1, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (104, 101, '0,100,101', '市场部门', 2, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (105, 101, '0,100,101', '测试部门', 3, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (106, 101, '0,100,101', '财务部门', 4, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (107, 101, '0,100,101', '运维部门', 5, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (108, 102, '0,100,102', '市场部门', 1, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (109, 102, '0,100,102', '财务部门', 2, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (47373459289804800, 101, '0,100,101', '测试', 6, '', '', '', '0', '', '2024-05-11 01:25:12.861', '', '2024-05-11 01:40:37.025', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint NOT NULL COMMENT '字典编码',
  `dict_sort` int DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`dict_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', '性别男', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', '性别女', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', '性别未知', 'admin', '2024-04-22 08:36:41.000', '', '2024-05-14 12:59:54.334', NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', '显示菜单', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', '隐藏菜单', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', '正常状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', '停用状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', '正常状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', '停用状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', '默认分组', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', '系统分组', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', '系统默认是', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', '系统默认否', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', '通知', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', '公告', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', '正常状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', '关闭状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (18, 99, '其他', '0', 'sys_oper_type', '', 'info', 'N', '0', '其他操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (19, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', '新增操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (20, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', '修改操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (21, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', '删除操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (22, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', '授权操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (23, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', '导出操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (24, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', '导入操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (25, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', '强退操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (26, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', '生成操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (27, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', '清空操作', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (28, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', '正常状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (29, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', '停用状态', 'admin', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (48635534888669184, 0, 'ces', 'ces', 'sys_test', '', 'danger', 'N', '0', '测试', '', '2024-05-14 13:00:15.125', '', '2024-05-14 13:00:59.000', '2024-05-14 13:49:15.714');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (48659447857287168, 0, '测试', 'test001', 'sys_test', '', 'default', 'N', '0', '', '', '2024-05-14 14:35:16.420', '', '2024-05-14 14:35:16.418', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint NOT NULL COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`dict_id`),
  UNIQUE KEY `dict_type` (`dict_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '用户性别列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '菜单状态列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '系统开关列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '任务状态列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '任务分组列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '系统是否列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '通知类型列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '通知状态列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '操作类型列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '登录状态列表', '2024-04-22 08:36:41.000', '', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `remark`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (48617958791974912, '测试', 'sys_test', '0', '', 'ces', '2024-05-14 11:50:24.661', '', '2024-05-14 13:01:55.344', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint NOT NULL COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '3' COMMENT '计划执行错误策略（1立即执行 2执行一次 3放弃执行）',
  `concurrent` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (1, '系统默认（无参）', 'DEFAULT', 'ryTask.ryNoParams', '0/10 * * * * ?', '3', '1', '1', 'admin', '2024-04-22 08:36:41', '', NULL, '');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (2, '系统默认（有参）', 'DEFAULT', 'ryTask.ryParams(\'ry\')', '0/15 * * * * ?', '3', '1', '1', 'admin', '2024-04-22 08:36:41', '', NULL, '');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (3, '系统默认（多参）', 'DEFAULT', 'ryTask.ryMultipleParams(\'ry\', true, 2000L, 316.50D, 100)', '0/20 * * * * ?', '3', '1', '1', 'admin', '2024-04-22 08:36:41', '', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log` (
  `job_log_id` bigint NOT NULL COMMENT '任务日志ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '日志信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '异常信息',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='定时任务调度日志表';

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor` (
  `info_id` bigint NOT NULL COMMENT '访问ID',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '用户账号',
  `ipaddr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`),
  KEY `idx_sys_logininfor_s` (`status`),
  KEY `idx_sys_logininfor_lt` (`login_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint DEFAULT '0' COMMENT '父菜单ID',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '组件路径',
  `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '路由参数',
  `is_frame` int DEFAULT '1' COMMENT '是否为外链（0是 1否）',
  `is_cache` int DEFAULT '0' COMMENT '是否缓存（0缓存 1不缓存）',
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '菜单状态（0正常 1停用）',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1, '系统管理', 0, 1, 'system', NULL, '', 1, 0, 'M', '0', '0', '', 'system', 'admin', '2024-04-22 08:36:40.000', '', '2024-05-09 01:53:12.330', NULL, '系统管理目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (2, '系统监控', 0, 2, 'monitor', NULL, '', 1, 0, 'M', '0', '0', '', 'monitor', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '系统监控目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (3, '系统工具', 0, 3, 'tool', NULL, '', 1, 0, 'M', '0', '0', '', 'tool', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '系统工具目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (4, '官网', 0, 4, 'http://ruoyi.vip', NULL, '', 0, 0, 'M', '0', '0', '', 'guide', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '若依官网地址');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (100, '用户管理', 1, 1, 'user', 'system/user/index', '', 1, 0, 'C', '0', '0', 'system:user:list', 'user', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '用户管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (101, '角色管理', 1, 2, 'role', 'system/role/index', '', 1, 0, 'C', '0', '0', 'system:role:list', 'peoples', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '角色管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (102, '菜单管理', 1, 3, 'menu', 'system/menu/index', '', 1, 0, 'C', '0', '0', 'system:menu:list', 'tree-table', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '菜单管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (103, '部门管理', 1, 4, 'dept', 'system/dept/index', '', 1, 0, 'C', '0', '0', 'system:dept:list', 'tree', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '部门管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (104, '岗位管理', 1, 5, 'post', 'system/post/index', '', 1, 0, 'C', '0', '0', 'system:post:list', 'post', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '岗位管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (105, '字典管理', 1, 6, 'dict', 'system/dict/index', '', 1, 0, 'C', '0', '0', 'system:dict:list', 'dict', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '字典管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (106, '参数设置', 1, 7, 'config', 'system/config/index', '', 1, 0, 'C', '0', '0', 'system:config:list', 'edit', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '参数设置菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (107, '通知公告', 1, 8, 'notice', 'system/notice/index', '', 1, 0, 'C', '0', '0', 'system:notice:list', 'message', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '通知公告菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (108, '日志管理', 1, 9, 'log', '', '', 1, 0, 'M', '0', '0', '', 'log', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '日志管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (109, '在线用户', 2, 1, 'online', 'monitor/online/index', '', 1, 0, 'C', '0', '0', 'monitor:online:list', 'online', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '在线用户菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (110, '定时任务', 2, 2, 'job', 'monitor/job/index', '', 1, 0, 'C', '0', '0', 'monitor:job:list', 'job', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '定时任务菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (111, '数据监控', 2, 3, 'druid', 'monitor/druid/index', '', 1, 0, 'C', '0', '0', 'monitor:druid:list', 'druid', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '数据监控菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (112, '服务监控', 2, 4, 'server', 'monitor/server/index', '', 1, 0, 'C', '0', '0', 'monitor:server:list', 'server', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '服务监控菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (113, '缓存监控', 2, 5, 'cache', 'monitor/cache/index', '', 1, 0, 'C', '0', '0', 'monitor:cache:list', 'redis', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '缓存监控菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (114, '缓存列表', 2, 6, 'cacheList', 'monitor/cache/list', '', 1, 0, 'C', '0', '0', 'monitor:cache:list', 'redis-list', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '缓存列表菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (115, '表单构建', 3, 1, 'build', 'tool/build/index', '', 1, 0, 'C', '0', '0', 'tool:build:list', 'build', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '表单构建菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (116, '代码生成', 3, 2, 'gen', 'tool/gen/index', '', 1, 0, 'C', '0', '0', 'tool:gen:list', 'code', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '代码生成菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (117, '系统接口', 3, 3, 'swagger', 'tool/swagger/index', '', 1, 0, 'C', '0', '0', 'tool:swagger:list', 'swagger', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '系统接口菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', '', 1, 0, 'C', '0', '0', 'monitor:operlog:list', 'form', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '操作日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', '', 1, 0, 'C', '0', '0', 'monitor:logininfor:list', 'logininfor', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '登录日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1000, '用户查询', 100, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:user:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1001, '用户新增', 100, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:user:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1002, '用户修改', 100, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:user:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1003, '用户删除', 100, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:user:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1004, '用户导出', 100, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:user:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1005, '用户导入', 100, 6, '', '', '', 1, 0, 'F', '0', '0', 'system:user:import', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1006, '重置密码', 100, 7, '', '', '', 1, 0, 'F', '0', '0', 'system:user:resetPwd', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1007, '角色查询', 101, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:role:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1008, '角色新增', 101, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:role:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1009, '角色修改', 101, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1010, '角色删除', 101, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:role:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1011, '角色导出', 101, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:role:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1012, '菜单查询', 102, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1013, '菜单新增', 102, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1014, '菜单修改', 102, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1015, '菜单删除', 102, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:menu:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1016, '部门查询', 103, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1017, '部门新增', 103, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1018, '部门修改', 103, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1019, '部门删除', 103, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:dept:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1020, '岗位查询', 104, 1, '', '', '', 1, 0, 'F', '0', '0', 'system:post:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1021, '岗位新增', 104, 2, '', '', '', 1, 0, 'F', '0', '0', 'system:post:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1022, '岗位修改', 104, 3, '', '', '', 1, 0, 'F', '0', '0', 'system:post:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1023, '岗位删除', 104, 4, '', '', '', 1, 0, 'F', '0', '0', 'system:post:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1024, '岗位导出', 104, 5, '', '', '', 1, 0, 'F', '0', '0', 'system:post:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1025, '字典查询', 105, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1026, '字典新增', 105, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1027, '字典修改', 105, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1028, '字典删除', 105, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1029, '字典导出', 105, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:dict:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1030, '参数查询', 106, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1031, '参数新增', 106, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1032, '参数修改', 106, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1033, '参数删除', 106, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1034, '参数导出', 106, 5, '#', '', '', 1, 0, 'F', '0', '0', 'system:config:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1035, '公告查询', 107, 1, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1036, '公告新增', 107, 2, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1037, '公告修改', 107, 3, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1038, '公告删除', 107, 4, '#', '', '', 1, 0, 'F', '0', '0', 'system:notice:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1039, '操作查询', 500, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1040, '操作删除', 500, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1041, '日志导出', 500, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1042, '登录查询', 501, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1043, '登录删除', 501, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1044, '日志导出', 501, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1045, '账户解锁', 501, 4, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:unlock', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1046, '在线查询', 109, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1047, '批量强退', 109, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:batchLogout', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1048, '单条强退', 109, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:online:forceLogout', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1049, '任务查询', 110, 1, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1050, '任务新增', 110, 2, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:add', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1051, '任务修改', 110, 3, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1052, '任务删除', 110, 4, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1053, '状态修改', 110, 5, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:changeStatus', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1054, '任务导出', 110, 6, '#', '', '', 1, 0, 'F', '0', '0', 'monitor:job:export', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1055, '生成查询', 116, 1, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:query', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1056, '生成修改', 116, 2, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:edit', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1057, '生成删除', 116, 3, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:remove', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1058, '导入代码', 116, 4, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:import', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1059, '预览代码', 116, 5, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:preview', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (1060, '生成代码', 116, 6, '#', '', '', 1, 0, 'F', '0', '0', 'tool:gen:code', '#', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`, `remark`) VALUES (46655658472247296, '测试', 0, 4, 'test', '', '', 1, 0, 'M', '0', '1', '', '404', '', '2024-05-09 01:52:55.805', '', '2024-05-28 10:10:34.678', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `notice_id` int NOT NULL COMMENT '公告ID',
  `notice_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告标题',
  `notice_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` longblob COMMENT '公告内容',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
BEGIN;
INSERT INTO `sys_notice` (`notice_id`, `notice_title`, `notice_type`, `notice_content`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (1, '温馨提醒：2018-07-01 若依新版本发布啦', '2', 0xE696B0E78988E69CACE58685E5AEB9, '0', 'admin', '2024-04-22 08:36:41', '', NULL, '管理员');
INSERT INTO `sys_notice` (`notice_id`, `notice_title`, `notice_type`, `notice_content`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (2, '维护通知：2018-07-01 若依系统凌晨维护', '1', 0xE7BBB4E68AA4E58685E5AEB9, '0', 'admin', '2024-04-22 08:36:41', '', NULL, '管理员');
COMMIT;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint NOT NULL COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '模块标题',
  `business_type` int DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '请求方式',
  `operator_type` int DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '返回参数',
  `status` int DEFAULT '0' COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  `cost_time` bigint DEFAULT '0' COMMENT '消耗时间',
  PRIMARY KEY (`oper_id`),
  KEY `idx_sys_oper_log_bt` (`business_type`),
  KEY `idx_sys_oper_log_s` (`status`),
  KEY `idx_sys_oper_log_ot` (`oper_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (1, 'ceo', '董事长', 1, '0', '', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (2, 'se', '项目经理', 2, '0', '', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (3, 'hr', '人力资源', 3, '0', '', 'admin', '2024-04-22 08:36:39.000', '', NULL, NULL);
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (4, 'user', '普通员工', 4, '0', '', 'admin', '2024-04-22 08:36:40.000', '', NULL, NULL);
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (5, 'post_test', '测试', 5, '1', '1', '', '2024-05-11 13:19:23.285', '', '2024-05-11 13:19:39.038', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `menu_check_strictly` tinyint(1) DEFAULT '1' COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) DEFAULT '1' COMMENT '部门树选择项是否关联显示',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=46139014180245505 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remark`, `deleted_at`) VALUES (1, '超级管理员', 'admin', 1, '1', 1, 1, '0', 'admin', '2024-04-22 08:36:40.000', '', NULL, '超级管理员', NULL);
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remark`, `deleted_at`) VALUES (2, '普通角色', 'common', 2, '2', 1, 1, '0', 'admin', '2024-04-22 08:36:40.000', '', NULL, '普通角色', NULL);
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remark`, `deleted_at`) VALUES (46102722679672832, '测试', 'test1', 3, '1', 1, 1, '0', '', '2024-05-07 13:15:45.636', '', '2024-05-07 13:23:23.551', '', '2024-05-07 13:30:54.083');
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remark`, `deleted_at`) VALUES (46137415466749952, '测试', 'test', 3, '1', 1, 1, '0', '', '2024-05-07 15:33:37.021', '', '2024-05-09 11:40:09.646', '', NULL);
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`, `remark`, `deleted_at`) VALUES (46139014180245504, '测试2', 'test2', 4, '1', 1, 1, '0', '', '2024-05-07 15:39:58.207', '', '2024-05-07 15:39:58.204', '', '2024-05-07 15:40:04.023');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (2, 100, NULL);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (2, 101, NULL);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (2, 105, NULL);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (46137415466749952, 100, '2024-05-09 11:40:09.657');
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (46137415466749952, 101, '2024-05-09 11:40:09.657');
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (46137415466749952, 103, '2024-05-09 11:40:09.657');
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`, `deleted_at`) VALUES (46137415466749952, 104, '2024-05-09 11:40:09.657');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 2, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 3, '2024-05-09 02:01:14.561');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 4, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 100, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 101, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 102, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 103, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 104, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 105, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 106, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 107, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 108, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 109, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 110, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 111, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 112, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 113, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 114, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 115, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 116, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 117, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 500, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 501, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1000, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1001, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1002, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1003, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1004, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1005, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1006, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1007, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1008, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1009, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1010, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1011, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1012, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1013, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1014, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1015, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1016, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1017, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1018, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1019, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1020, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1021, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1022, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1023, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1024, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1025, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1026, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1027, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1028, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1029, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1030, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1031, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1032, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1033, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1034, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1035, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1036, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1037, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1038, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1039, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1040, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1041, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1042, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1043, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1044, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1045, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1046, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1047, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1048, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1049, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1050, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1051, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1052, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1053, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1054, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1055, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1056, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1057, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1058, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1059, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (2, 1060, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 100, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1000, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1001, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1002, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1003, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1004, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1005, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1006, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 101, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1007, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1008, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1009, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1010, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1011, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 102, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1012, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1013, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1014, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1015, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 103, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1016, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1017, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1018, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1019, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 104, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1020, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1021, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1022, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1023, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1024, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 105, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1025, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1026, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1027, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1028, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1029, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 106, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1030, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1031, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1032, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1033, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1034, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 107, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1035, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1036, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1037, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1038, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 108, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 500, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1039, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1040, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1041, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 501, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1042, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1043, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1044, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1045, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 2, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 109, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1046, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1047, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1048, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 110, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1049, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1050, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1051, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1052, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1053, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1054, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 111, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 112, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 113, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 114, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 3, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 115, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 116, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1055, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1056, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1057, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1058, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1059, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1060, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 117, '2024-05-07 15:40:43.754');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 100, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1000, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1001, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1002, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1003, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1004, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1005, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1006, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 101, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1007, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1008, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1009, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1010, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1011, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 102, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1012, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1013, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1014, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1015, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 103, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1016, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1017, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1018, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1019, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 104, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1020, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1021, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1022, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1023, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1024, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 105, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1025, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1026, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1027, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1028, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1029, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 106, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1030, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1031, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1032, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1033, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1034, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 107, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1035, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1036, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1037, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1038, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 108, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 500, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1039, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1040, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1041, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 501, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1042, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1043, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1044, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1045, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 2, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 109, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1046, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1047, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1048, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 110, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1049, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1050, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1051, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1052, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1053, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1054, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 111, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 112, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 113, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 114, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 3, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 115, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 116, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1055, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1056, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1057, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1058, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1059, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 1060, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46139014180245504, 117, '2024-05-07 15:40:04.028');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 100, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1000, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1001, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1002, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1003, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1004, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1005, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1006, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 101, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1007, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1008, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1009, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1010, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1011, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 102, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1012, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1013, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1014, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1015, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 103, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1016, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1017, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1018, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1019, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 104, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1020, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1021, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1022, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1023, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1024, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 105, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1025, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1026, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1027, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1028, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1029, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 106, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1030, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1031, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1032, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1033, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1034, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 107, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1035, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1036, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1037, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1038, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 108, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 500, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1039, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1040, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1041, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 501, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1042, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1043, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1044, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1045, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 2, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 109, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1046, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1047, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1048, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 110, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1049, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1050, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1051, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1052, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1053, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1054, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 111, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 112, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 113, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 114, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 3, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 115, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 116, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1055, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1056, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1057, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1058, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1059, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1060, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 117, '2024-05-08 01:11:57.658');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 100, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1000, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1001, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1002, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1003, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1004, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1005, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1006, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 101, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1007, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1008, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1009, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1010, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1011, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 102, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1012, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1013, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1014, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1015, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 103, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1016, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1017, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1018, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1019, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 104, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1020, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1021, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1022, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1023, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1024, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 105, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1025, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1026, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1027, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1028, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1029, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 106, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1030, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1031, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1032, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1033, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1034, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 107, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1035, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1036, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1037, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1038, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 108, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 500, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1039, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1040, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1041, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 501, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1042, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1043, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1044, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1045, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 2, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 109, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1046, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1047, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1048, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 110, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1049, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1050, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1051, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1052, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1053, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1054, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 111, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 112, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 113, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 114, '2024-05-08 01:15:41.681');
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 2, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 100, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1000, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1001, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1002, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1003, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1004, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1005, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1006, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 101, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1007, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1008, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1009, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1010, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1011, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 102, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1012, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1013, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1014, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1015, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 103, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1016, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1017, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1018, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1019, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 104, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1020, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1021, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1022, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1023, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1024, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 105, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1025, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1026, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1027, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1028, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1029, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 106, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1030, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1031, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1032, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1033, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1034, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 107, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1035, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1036, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1037, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1038, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 108, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 500, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1039, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1040, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1041, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 501, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1042, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1043, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1044, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1045, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 109, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1046, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1047, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1048, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 110, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1049, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1050, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1051, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1052, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1053, NULL);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`, `deleted_at`) VALUES (46137415466749952, 1054, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `dept_id` bigint DEFAULT NULL COMMENT '部门ID',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '用户邮箱',
  `phone_number` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '密码',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`user_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phone_number`, `sex`, `avatar`, `password`, `status`, `del_flag`, `login_ip`, `login_date`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (1, 103, 'admin', 'admin', '00', 'ry@163.com', '15888888888', '1', '', '$2a$10$HJkbovt22k4h4TsrYrlpaeKz1Fa08E05eBBrtJCC8Oog2UDO6VqBq', '0', '0', '127.0.0.1', '2024-04-24 19:40:09', '管理员', 'admin', '2024-04-24 19:40:09', '', NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phone_number`, `sex`, `avatar`, `password`, `status`, `del_flag`, `login_ip`, `login_date`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (2, 105, 'PTJ', '浦同矫', '00', 'ry@qq.com', '15666666666', '1', '', '$2a$10$HJkbovt22k4h4TsrYrlpaeKz1Fa08E05eBBrtJCC8Oog2UDO6VqBq', '0', '0', '127.0.0.1', '2024-04-24 19:40:09', '测试员', 'admin', '2024-04-24 19:40:09', '', NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phone_number`, `sex`, `avatar`, `password`, `status`, `del_flag`, `login_ip`, `login_date`, `remark`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_at`) VALUES (45049925926391808, 101, 'test', '测试', '00', '15391493301@163.com', '15391493301', '0', '', '$2a$10$Pyv6lOGKzOcodExba7hI2.FVzAnEqeipZDVwpEouxlfQ/h/RlWila', '0', '0', '', '2024-05-04 15:32:19', '测试', '', '2024-05-04 15:32:19', '', '2024-05-09 10:41:01', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (1, 1, NULL);
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (2, 2, NULL);
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 2, '2024-05-08 16:17:39.786');
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 3, '2024-05-08 16:17:39.786');
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-08 16:17:39.786');
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 2, '2024-05-09 10:41:01.427');
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-09 10:41:01.427');
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 2, NULL);
INSERT INTO `sys_user_post` (`user_id`, `post_id`, `deleted_at`) VALUES (45049925926391808, 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 1, NULL);
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 2, NULL);
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-08 16:17:39.783');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 11:59:36.133');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 11:59:34.041');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 12:00:02.246');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 11:59:58.368');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 12:02:39.629');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 12:02:39.629');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 12:09:13.547');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 12:09:13.547');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 12:12:21.437');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 12:12:21.437');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 12:15:29.375');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 12:15:29.375');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 14:39:29.429');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 14:39:27.507');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-08 16:13:17.287');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 16:13:17.287');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-08 22:25:43.036');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-08 22:25:52.020');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 2, '2024-05-08 22:25:52.020');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-08 23:01:57.701');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 10:44:09.315');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 10:44:09.315');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-08 23:01:57.701');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-09 10:40:54.080');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-09 10:41:01.424');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:41:01.424');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 2, '2024-05-09 10:41:01.424');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 1, '2024-05-09 11:03:34.071');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:44:18.835');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 10:55:52.794');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 10:55:52.794');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:56:58.660');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 10:57:12.668');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 10:57:12.668');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:57:29.452');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 10:57:42.227');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 10:57:42.227');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:57:46.710');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 10:58:25.975');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 10:58:25.975');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 10:58:29.715');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 11:00:47.994');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 11:00:47.994');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 11:02:04.294');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (1, 46137415466749952, '2024-05-09 11:03:04.924');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (2, 46137415466749952, '2024-05-09 11:03:04.924');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, '2024-05-09 11:03:34.071');
INSERT INTO `sys_user_role` (`user_id`, `role_id`, `deleted_at`) VALUES (45049925926391808, 46137415466749952, NULL);
COMMIT;

-- ----------------------------
-- Function structure for GetDeptSubordinates
-- ----------------------------
DROP FUNCTION IF EXISTS `GetDeptSubordinates`;
delimiter ;;
CREATE FUNCTION `GetDeptSubordinates`(cur_id INT)
 RETURNS text CHARSET utf8mb4
  DETERMINISTIC
BEGIN
    DECLARE output TEXT;

    WITH RECURSIVE Subordinates AS (
        SELECT dept_id, parent_id
        FROM sys_dept
        WHERE dept_id = cur_id
        UNION ALL
        SELECT t.dept_id, t.parent_id
        FROM sys_dept t
                 JOIN Subordinates s ON t.parent_id = s.dept_id
    )
    SELECT GROUP_CONCAT(dept_id) INTO output
    FROM Subordinates;

    RETURN output;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
