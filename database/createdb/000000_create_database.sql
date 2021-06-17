/*******************************************************************************
 * 
 *                          создание базы данных
 * 
 ******************************************************************************/
-- 0. database create
SELECT 'create database if not exists plantbook_admin with owner plantbook_admin template template1;'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'plantbook_admin')\gexec
-- trigger for db-build +1