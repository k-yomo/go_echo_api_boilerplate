## Migration Usage

### Generate migration files
```
// it will create up & down migration files
make gen_migration name=NAME
```

- example
```
make gen_migration name=create_users

// two migration files below will be created
// 20190625011352_create_users.up.sql 
// 20190625011352_create_users.down.sql
```

### Apply up migrations
- It will apply all pending migrations
```
make up_migrate
```

```
// force execution
make up_migrate arg=-f
```

### Apply down migrations
- It will rollback migration by 1 step.
```
make down_migrate
```

```
// force execution
make down_migrate arg=-f
```

### Drop all migrations
- It will drop all migrations
```
make drop_migrate
```

```
// force execution
make drop_migrate arg=-f
```

### Generate models
- It will generate models from the DB schema
```
make gen_modles
```
