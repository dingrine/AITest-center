CREATE TABLE `exam_result_detail` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键标识',
    `exam_id` BIGINT NOT NULL COMMENT '考试ID，关联exam_info表',
    `total_score` TINYINT COMMENT '总分',
    `financial_knowledge_score` TINYINT COMMENT '金融业务知识分',
    `regulatory_compliance_score` TINYINT COMMENT '法规合规分',
    `risk_management_score` TINYINT COMMENT '安全与风控分',
    `trade_conduct_score` TINYINT COMMENT '交易执行与市场行为分',
    `investor_ethics_score` TINYINT COMMENT '投资者保护与职业道德分',
    `is_passed` BOOLEAN COMMENT '考试是否通过',
    `total_questions` INT COMMENT '题目总数',
    `correct_answers` INT COMMENT '正确回答题目数',
    `incorrect_answers` INT COMMENT '错误回答题目数',
    `passed_at` DATETIME COMMENT '通过认证时间',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (`id`),
    FOREIGN KEY (`exam_id`) REFERENCES exam_info(`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试结果详情表';