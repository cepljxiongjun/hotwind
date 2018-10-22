CREATE TABLE public.user (
  id serial PRIMARY KEY ,
  username varchar(255) DEFAULT '' NOT NULL,
  email varchar(255) DEFAULT '' NOT NULL,
  password_hash char(32) DEFAULT '' NOT NULL,
  created_at timestamp NOT  NULL,
  updated_at timestamp NOT  NULL
);
comment on column public.user.username is '用户名';
comment on column public.user.email is '邮箱';
comment on column public.user.password_hash is '密码';
comment on column public.user.created_at is '创建时间';
comment on column public.user.updated_at is '更新时间';