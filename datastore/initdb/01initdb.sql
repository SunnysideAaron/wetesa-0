DROP TABLE IF EXISTS public.client;

CREATE TABLE IF NOT EXISTS public.client
(
    client_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    address character varying COLLATE pg_catalog."default" NULL,
    CONSTRAINT client_pk PRIMARY KEY (client_id)
)

-- TODO why is this here? What does it do? For some reason if I comment it out the tables don't get created.
TABLESPACE pg_default;

DROP TABLE IF EXISTS public.order;

-- TODO date_submitted needs to be a date type will have to research and do research on UTC and server time.
CREATE TABLE IF NOT EXISTS public.order
(
    order_id uuid NOT NULL,
    client_id uuid NOT NULL,
    submitted_date character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT order_pk PRIMARY KEY (order_id)
)

TABLESPACE pg_default;

DROP TABLE IF EXISTS public.order_product;

CREATE TABLE IF NOT EXISTS public.order_product
(
    order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    amount integer NOT NULL
)

TABLESPACE pg_default;

DROP TABLE IF EXISTS public.product;

CREATE TABLE IF NOT EXISTS public.product
(
    product_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT product_pk PRIMARY KEY (product_id)
)

TABLESPACE pg_default;

DROP TABLE IF EXISTS public.user;

CREATE TABLE IF NOT EXISTS public.user
(
    user_id uuid NOT NULL DEFAULT gen_random_uuid(),
    login_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

-- TODO do dependancies / joins after all tables created