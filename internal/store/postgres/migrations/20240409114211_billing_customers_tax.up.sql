ALTER TABLE billing_customers ADD COLUMN IF NOT EXISTS tax jsonb DEFAULT '{}'::jsonb;