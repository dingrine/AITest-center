CREATE TABLE `exam_info` (
    `id` INT AUTO_INCREMENT NOT NULL COMMENT '考试ID，关联模型表',
    `detrader_id` INT NOT NULL COMMENT '模型ID',
    `detrader_name` VARCHAR(100) NOT NULL COMMENT '模型名称（冗余字段，便于查询）',
    `evaluation_system` TINYINT NOT NULL COMMENT '评测体系（0:ProTest、1:TradingTest）',
    `region_cert_type` TINYINT NOT NULL COMMENT '地区认证类型（0:MAT-CN、1:MAT-HK、2:MAT-US）',
    `exam_count` INT,
    `passed_count` INT NOT NULL DEFAULT 0,
    `latest_exam_name` VARCHAR(100) NOT NULL COMMENT '最近一次考试的名称',
    `latest_exam_status` TINYINT NOT NULL DEFAULT 0,
    `latest_certified_at` DATETIME COMMENT '最近一次通过考试的时间',
    `certification_status` TINYINT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_detrader_eval_region` (`detrader_id`, `evaluation_system`, `region_cert_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='考试总表';
