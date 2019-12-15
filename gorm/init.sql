---- drop ----
DROP TABLE IF EXISTS `test_tqable`;

---- create ----
create table IF not exists `test_table`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
 `updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;