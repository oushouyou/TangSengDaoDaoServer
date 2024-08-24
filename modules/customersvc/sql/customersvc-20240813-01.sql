-- +migrate Up

-- 客户服务群组表，每个客户一个外部服务群
create table `customersvc_group`
(
    customer_external_no VARCHAR(40) NOT NULL DEFAULT ''  COMMENT '外部客户编号',
    user_uid VARCHAR(40) NOT NULL DEFAULT ''  COMMENT '用户ID',
    group_no VARCHAR(40) NOT NULL DEFAULT ''  COMMENT '群ID',
    created_at timeStamp     not null DEFAULT CURRENT_TIMESTAMP,
    updated_at timeStamp     not null DEFAULT CURRENT_TIMESTAMP 
);
CREATE UNIQUE INDEX customer_external_no on `customersvc_group` (customer_external_no);