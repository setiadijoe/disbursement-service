CREATE TABLE flip_disbursement (
	id int8 NOT NULL,
	amount float8 NOT NULL,
	status varchar(100) NOT NULL,
	"timestamp" timestamp NOT NULL,
	bank_code varchar(50) NOT NULL,
	account_number varchar(50) NOT NULL,
	beneficiary_name varchar(50) NOT NULL,
	remark varchar(50) NOT NULL,
	receipt varchar(100) NULL,
	time_served timestamp NULL,
	fee float8 NOT NULL,
	CONSTRAINT pk_id PRIMARY KEY (id)
);

CREATE TABLE flip_disbursement_history (
	id int8 NOT NULL,
	amount float8 NOT NULL,
	status varchar(100) NOT NULL,
	"timestamp" timestamp NOT NULL,
	bank_code varchar(50) NOT NULL,
	account_number varchar(50) NOT NULL,
	beneficiary_name varchar(50) NOT NULL,
	remark varchar(50) NOT NULL,
	receipt varchar(100) NULL,
	time_served timestamp NULL,
	fee float8 NOT NULL,
	CONSTRAINT pk_flip_disbursement_history_id PRIMARY KEY (id)
);
