# 4 order api

SMS and Call auth with sms.ru provider

## db scripts (DDL)

```sql
create table public.users
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    phone      text,
    name       varchar(100),
    birthdate  date
);

alter table public.users
    owner to postgres;

create unique index idx_users_phone
    on public.users (phone);

create index idx_users_deleted_at
    on public.users (deleted_at);

create table public.products
(
    id          bigserial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    name        varchar(100),
    description varchar(500),
    images      text[],
    price       numeric(10, 2)
);

alter table public.products
    owner to postgres;

create index idx_products_deleted_at
    on public.products (deleted_at);

create unique index idx_products_name
    on public.products (name);

create table public.carts
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id    bigint
        constraint fk_carts_user
            references public.users
);

alter table public.carts
    owner to postgres;

create index idx_carts_deleted_at
    on public.carts (deleted_at);

create table public.cart_items
(
    cart_id    bigint
        constraint fk_cart_items_cart
            references public.carts,
    product_id bigint
        constraint fk_cart_items_product
            references public.products,
    count      bigint default 1
);

alter table public.cart_items
    owner to postgres;

create unique index cart_item
    on public.cart_items (cart_id, product_id);

```