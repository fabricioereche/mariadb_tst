CREATE TABLE `beer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL DEFAULT '',
  `type` int(11) NOT NULL,
  `style` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `beer_id_IDX` (`id`) USING BTREE
);