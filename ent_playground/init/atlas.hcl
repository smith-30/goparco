table "users" {
  schema = schema.main
  column "age" {
    null = false
    type = integer
  }
  column "name" {
    null    = false
    type    = varchar(255)
    default = "unknown"
  }
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}
