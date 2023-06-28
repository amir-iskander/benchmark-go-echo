-- 001_create_users_table.up.sql

CREATE TABLE users (
  id SERIAL PRIMARY KEY,  
  name VARCHAR(100),
  email VARCHAR(100)
);

`001_create_users_table.down.sql`:

```sql 
DROP TABLE users;
```

