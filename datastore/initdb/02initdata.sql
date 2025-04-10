-- TODO inserts are NOT incrementing primary keys.
INSERT INTO public.client(name) VALUES ('AT&T');
INSERT INTO public.client(name, address) VALUES ('Dr. Tom', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name) VALUES ('ACME inc.');
INSERT INTO public.client(name) VALUES ('AutoHouse llc');

-- TODO Punting for now. when we want some test data will need to adjust for uuid
-- INSERT INTO public.order(order_id, client_id, date_submited) VALUES (1, 1, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, date_submited) VALUES (2, 2, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, date_submited) VALUES (3, 3, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, date_submited) VALUES (4, 4, '2-10-2020');

-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (1, 1, 1);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (2, 2, 2);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (3, 3, 3);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (4, 4, 4);

-- INSERT INTO public.product(name) VALUES ('Wire');
-- INSERT INTO public.product(name) VALUES ('Pencil');
-- INSERT INTO public.product(name) VALUES ('Apple');
-- INSERT INTO public.product(name) VALUES ('Basket');

-- INSERT INTO public.user(login_name) VALUES ('Mary');
-- INSERT INTO public.user(login_name) VALUES ('Tom');
-- INSERT INTO public.user(login_name) VALUES ('John');
-- INSERT INTO public.user(login_name) VALUES ('Sarah');