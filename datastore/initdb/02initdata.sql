INSERT INTO public.client(name) VALUES ('AT&T');
INSERT INTO public.client(name, address) VALUES ('Dr. Tom', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name) VALUES ('ACME inc.');
INSERT INTO public.client(name) VALUES ('AutoHouse llc');
INSERT INTO public.client(name, address) VALUES ('Dairy Queen', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Taco Bell', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('TJ Max', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('DC Dental', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Death Star', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Verizon', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('T-Mobile', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Burger King', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Dr. Pepper', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Office Max', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Blockbuster', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Toys R Us', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Sears', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Bloomingdale', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('7-11', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('Cold Stone', '124 SW 15 St, Fargo, ND 99229');
INSERT INTO public.client(name, address) VALUES ('NASA', '124 SW 15 St, Fargo, ND 99229');

-- TODO Punting for now. when we want some test data will need to adjust for uuid
-- INSERT INTO public.order(order_id, client_id, submitted_date) VALUES (1, 1, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, submitted_date) VALUES (2, 2, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, submitted_date) VALUES (3, 3, '2-10-2020');
-- INSERT INTO public.order(order_id, client_id, submitted_date) VALUES (4, 4, '2-10-2020');

-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (1, 1, 1);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (2, 2, 2);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (3, 3, 3);
-- INSERT INTO public.order_product(order_id, product_id, amount) VALUES (4, 4, 4);

INSERT INTO public.product(name) VALUES ('Wire');
INSERT INTO public.product(name) VALUES ('Pencil');
INSERT INTO public.product(name) VALUES ('Apple');
INSERT INTO public.product(name) VALUES ('Basket');

INSERT INTO public.user(login_name) VALUES ('Mary');
INSERT INTO public.user(login_name) VALUES ('Tom');
INSERT INTO public.user(login_name) VALUES ('John');
INSERT INTO public.user(login_name) VALUES ('Sarah');