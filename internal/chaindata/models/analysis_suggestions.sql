CREATE TABLE `exam_info1` (
    `id` INT AUTO_INCREMENT NOT NULL COMMENT '考试ID，关联模型表'
    `detail_id` BIGINT NOT NULL COMMENT '关联结果详情表id',
    `strength_analysis` TEXT COMMENT '优势分析',
    `improvement_suggestions` TEXT COMMENT '改善建议',

    PRIMARY KEY (`id`),
    -- 添加外键约束（确保关联表存在）
    FOREIGN KEY (`detail_id`) 
        REFERENCES exam_result_detail(`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='考试总表';
