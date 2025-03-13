-- TODO inserts are NOT incrementing primary keys.
INSERT INTO public.client(client_id, name) VALUES (1, 'AT&T');
INSERT INTO public.client(client_id, name, address) VALUES (2, 'Dr. Tom', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(client_id, name) VALUES (3, 'ACME inc.');
INSERT INTO public.client(client_id, name) VALUES (4, 'AutoHouse llc');

INSERT INTO public.order(order_id, client_id, date_submited) VALUES (1, 1, '2-10-2020');
INSERT INTO public.order(order_id, client_id, date_submited) VALUES (2, 2, '2-10-2020');
INSERT INTO public.order(order_id, client_id, date_submited) VALUES (3, 3, '2-10-2020');
INSERT INTO public.order(order_id, client_id, date_submited) VALUES (4, 4, '2-10-2020');

INSERT INTO public.order_product(order_id, product_id, amount) VALUES (1, 1, 1);
INSERT INTO public.order_product(order_id, product_id, amount) VALUES (2, 2, 2);
INSERT INTO public.order_product(order_id, product_id, amount) VALUES (3, 3, 3);
INSERT INTO public.order_product(order_id, product_id, amount) VALUES (4, 4, 4);

INSERT INTO public.product(product_id, name) VALUES (1, 'Wire');
INSERT INTO public.product(product_id, name) VALUES (2, 'Pencil');
INSERT INTO public.product(product_id, name) VALUES (3, 'Apple');
INSERT INTO public.product(product_id, name) VALUES (4, 'Basket');

INSERT INTO public.user(user_id, login_name) VALUES (1, 'Mary');
INSERT INTO public.user(user_id, login_name) VALUES (2, 'Tom');
INSERT INTO public.user(user_id, login_name) VALUES (3, 'John');
INSERT INTO public.user(user_id, login_name) VALUES (4, 'Sarah');