CREATE TABLE IF NOT EXISTS `users`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username`   varchar(20)         NOT NULL DEFAULT '' COMMENT '用户名',
    `password`   varchar(64)         NOT NULL DEFAULT '' COMMENT '密码',
    `phone`      bigint(20) unsigned          DEFAULT NULL COMMENT '手机号',
    `created_at` timestamp           NOT NULL COMMENT '创建时间',
    `creator`    bigint(20) unsigned NOT NULL COMMENT '创建人',
    `updated_at` timestamp           NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `reviser`    bigint(20) unsigned NOT NULL COMMENT '修改人',
    `is_deleted` tinyint(1)          NOT NULL DEFAULT '0' COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;