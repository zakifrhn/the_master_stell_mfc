CREATE TABLE public.employee (
	id_employee uuid NOT NULL DEFAULT gen_random_uuid(),
	name_employee varchar(255) NOT NULL,
	address_employee varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT employee_pk PRIMARY KEY (id_employee)
);

