CREATE TABLE ticket_purchase (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	purchase_order_id UUID NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT ticket_purchase_id_pk PRIMARY KEY (id),
	CONSTRAINT ticket_purchase_id_fk FOREIGN KEY (user_id)
            REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT ticket_purchase_purchase_order_id_fk FOREIGN KEY (purchase_orders_id)
            REFERENCES purchase_orders (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE ticket_purchase IS 'Storage the head of the ticket_purchase';