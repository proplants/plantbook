-- alter to 25
ALTER TABLE public.users ALTER COLUMN name_user TYPE varchar(25) USING name_user::varchar;
ALTER TABLE public.users ALTER COLUMN email_addr TYPE varchar(25) USING email_addr::varchar;
ALTER TABLE public.users ALTER COLUMN first_name TYPE varchar(25) USING first_name::varchar;
ALTER TABLE public.users ALTER COLUMN last_name TYPE varchar(25) USING last_name::varchar;
ALTER TABLE public.users ALTER COLUMN phone_number TYPE varchar(25) USING phone_number::varchar;
--
ALTER TABLE public.user_plants ALTER COLUMN name_user_plant TYPE varchar(25) USING name_user_plant::varchar;
