CREATE USER 'hermes'@'localhost' IDENTIFIED WITH caching_sha2_password BY 'hermes';
GRANT USAGE ON *.* TO 'hermes'@'localhost';
ALTER USER 'hermes'@'localhost' REQUIRE NONE WITH MAX_QUERIES_PER_HOUR 0 MAX_CONNECTIONS_PER_HOUR 0 MAX_UPDATES_PER_HOUR 0 MAX_USER_CONNECTIONS 0;

GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, LOCK TABLES ON `hermes`.* TO 'hermes'@'localhost';
ALTER USER 'hermes'@'localhost' ;
