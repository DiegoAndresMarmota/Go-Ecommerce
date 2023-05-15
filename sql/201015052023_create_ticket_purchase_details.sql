CREATE TABLE ticket_purchase_details (
	id UUID NOT NULL,
	ticket_purchase_id UUID NOT NULL,
	product_id UUID NOT NULL,
	amount INTEGER NOT NULL,
	unit_price NUMERIC(10,2) NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT ticket_purchase_details_id_pk PRIMARY KEY (id),
	CONSTRAINT ticket_purchase_details_ticket_purchase_id_fk FOREIGN KEY (ticket_purchase_id)
            REFERENCES ticket_purchase (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT ticket_purchase_details_product_id_fk FOREIGN KEY (product_id)
            REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE ticket_purchase_details IS 'Storage the details of the ticket_purchase_details';