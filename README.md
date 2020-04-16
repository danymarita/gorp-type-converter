How to Run
1. Create postgresql database
2. execute sql
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS plans (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id int NOT NULL,
    campaigns json NOT NULL DEFAULT '{"data": []}',
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);
-- generate dummy data
INSERT INTO public.plans (user_id, campaigns) 
VALUES (60, '{"data":[{"category_id":2,"net_amount":120000},{"category_id":4,"net_amount":60000}]}');
```
3. Run **go run main.go**