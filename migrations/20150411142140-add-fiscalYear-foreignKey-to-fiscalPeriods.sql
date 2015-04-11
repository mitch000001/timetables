-- +migrate Up
SET client_min_messages = 'warning';

ALTER TABLE public.fiscal_periods ADD COLUMN fiscal_year_id integer;

-- +migrate Down
ALTER TABLE public.fiscal_periods DROP COLUMN fiscal_year_id;
