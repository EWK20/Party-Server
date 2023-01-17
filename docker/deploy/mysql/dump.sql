CREATE TABLE `table` (
  id INT NOT NULL auto_increment,
  capacity INT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE guest (
  id INT NOT NULL auto_increment,
  `name` VARCHAR(100) NOT NULL,
  table_id INT,
  accompanying_guests INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY(table_id) REFERENCES `table`(id)
);

INSERT INTO `` (`id`,`capacity`) VALUES (1,2);
INSERT INTO `` (`id`,`capacity`) VALUES (2,10);
INSERT INTO `` (`id`,`capacity`) VALUES (3,5);
INSERT INTO `` (`id`,`capacity`) VALUES (4,2);
INSERT INTO `` (`id`,`capacity`) VALUES (5,4);
INSERT INTO `` (`id`,`capacity`) VALUES (6,4);
INSERT INTO `` (`id`,`capacity`) VALUES (7,10);
INSERT INTO `` (`id`,`capacity`) VALUES (8,15);
INSERT INTO `` (`id`,`capacity`) VALUES (9,6);
INSERT INTO `` (`id`,`capacity`) VALUES (10,20);

INSERT INTO `getground`.`guest` (`name`, `table_id`, `acompanying_guests`) VALUES ('Echez', '2', '8');
INSERT INTO `getground`.`guest` (`name`, `table_id`, `acompanying_guests`) VALUES ('John', '6', '2');
INSERT INTO `getground`.`guest` (`name`, `table_id`, `acompanying_guests`) VALUES ('Sara', '10', '16');
INSERT INTO `getground`.`guest` (`name`, `table_id`, `acompanying_guests`) VALUES ('Hannah', '7', '6');

