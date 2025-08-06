CREATE TABLE `exam_qa` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '问答唯一标识',
    `exam_id` BIGINT NOT NULL COMMENT '考试Id，关联exam_info表',
    `question_code` VARCHAR(100) NOT NULL COMMENT '编号',
    `question` VARCHAR(255) NOT NULL COMMENT '问题',
    `answer` VARCHAR(255) NOT NULL COMMENT '标准回答',
    `model_answer` VARCHAR(255) NOT NULL COMMENT '模型回答',
    `analysis` VARCHAR(255) COMMENT '分析建议',
    `question_created_at` DATETIME COMMENT '问题创建时间',
    `answer_created_at` DATETIME COMMENT '回答创建时间',
    `qa_type` VARCHAR(100) COMMENT '类别',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (`id`),
    -- 添加外键约束（确保关联表存在）
    FOREIGN KEY (`exam_id`) 
        REFERENCES exam_info(`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考试问答记录表';