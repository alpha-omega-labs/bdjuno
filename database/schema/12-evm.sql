-- CREATE TABLE evm_params
-- (
--     one_row_id     BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
--     evm_denom      TEXT   NOT NULL,
--     enable_create  BOOLEAN   NOT NULL,
--     enable_call    BOOLEAN   NOT NULL,
--     extra_eips     NUMERIC[]   NOT NULL,
--     chain_config   JSONB   NOT NULL,
--     height         BIGINT  NOT NULL,
--     CHECK (one_row_id)
-- );

CREATE TABLE evm_params
(
    one_row_id     BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params   JSONB   NOT NULL,
    height         BIGINT  NOT NULL,
    CHECK (one_row_id)
);
