CREATE DATABASE IF NOT EXISTS go_echo_boilerplate_development;
CREATE DATABASE IF NOT EXISTS go_echo_boilerplate_test;

CREATE USER 'mysql'@'localhost' IDENTIFIED BY 'mysql';
GRANT ALL ON *.* TO 'mysql'@'%';
