



DROP TABLE IF EXISTS `tbl_user`;
create table tbl_user (
                          `id` int(11) not null auto_increment,
                          `user_id` int(11) not null default 0,
                          `user_name` varchar(50) not null,
                          `user_nike` varchar(50) not null,
                          `create_at` timestamp not null default current_timestamp,
                          `update_at` timestamp not null default current_timestamp on update current_timestamp,
                          `status` int(10) default 0	comment '0表示删除,1表示正常',
                          primary key (`id`)
) engine=innodb default charset=utf8mb4;