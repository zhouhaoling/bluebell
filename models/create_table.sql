drop table if exists 'user';

CREATE TABLE 'user' (
    'id' bigint(20) not null auto_increment,
    'user_id' bigint(20) not null,
    'username' VARCHAR(64) collate utf8mb4_general_ci not null,
    'password' VARCHAR(64) COLLATE utf8mb4_general_ci not null,
    'email' VARCHAR(64) COLLATE utf8mb4_general_ci,
    'gender' TINYINT(4) not null DEFAULT '0',
    'create_time' TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    'update_time' TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY ('id'),
    UNIQUE KEY 'idx_username' ('username') USING BTREE,
    UNIQUE KEY 'idx_user_id' ('user_id') USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

drop table if exists 'community';
CREATE TABLE community (
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction`  varchar(256) collate utf8mb4_general_ci not null,
    `create_time` timestamp not null default current_timestamp,
    `update_time` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

insert into community values ('1', '1', 'Go', 'Golang', '2023-10-1 08:10:10', '2023-10-1 08:10:10');
insert into community values ('2', '2', 'leetcode', '刷题', '2023-10-1 08:10:16', '2023-10-1 08:10:16');
insert into community values ('3', '3', 'Java', 'Java learn', '2023-10-1 08:20:10', '2023-10-1 08:20:10');
insert into community values ('4', '4', 'Game', '游戏', '2023-10-1 08:20:13', '2023-10-1 08:20:13');

DROP TABLE IF EXISTS `post`;
create table `post` (
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null comment '帖子id',
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题',
    `content` varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null comment '作者用户id',
    `community_id` bigint(20) not null comment '所属社区',
    `status` tinyint(4) not null default '1' comment '帖子状态',
    `create_time` timestamp null default current_timestamp comment '创建时间',
    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;