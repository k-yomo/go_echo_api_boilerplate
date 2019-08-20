CREATE DATABASE IF NOT EXISTS go_echo_api_boilerplate_development;
CREATE DATABASE IF NOT EXISTS go_echo_api_boilerplate_test;

CREATE USER 'mysql'@'localhost' IDENTIFIED BY 'mysql';
GRANT ALL ON *.* TO 'mysql'@'%';
