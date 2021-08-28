CREATE TABLE `shixun_famous_teacher` (
                                              `famous_teacher_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '名师id',
                                              `famous_teacher_photo` varchar(255) DEFAULT NULL COMMENT '名师头像',
                                              `famous_teacher_name` varchar(64) DEFAULT NULL COMMENT '名师姓名',
                                              `famous_teacher_mobile` varchar(11) DEFAULT NULL COMMENT '名师手机号',
                                              `famous_teacher_institution` varchar(255) DEFAULT NULL COMMENT '名师单位/机构',
                                              `famous_teacher_job_title` varchar(255) DEFAULT NULL COMMENT '名师职称',
                                              `famous_teacher_job_title_prove` varchar(255) DEFAULT NULL COMMENT '名师职称证明',
                                              `famous_teacher_introduce` text COMMENT '名师介绍',
                                              `famous_teacher_create_time` int(11) DEFAULT NULL COMMENT '名师创建时间',
                                              `famous_teacher_update_time` int(11) DEFAULT NULL COMMENT '名师更新时间',
                                              `famous_teacher_ishow` tinyint(2) DEFAULT '1' COMMENT '是否显示0-否，1-是',
                                              PRIMARY KEY (`famous_teacher_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='师训名师表';

#INSERT INTO tablesname VALUES( value1,value2)
