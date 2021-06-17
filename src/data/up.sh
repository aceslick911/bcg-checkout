cp -f ../test/database.db ./
jq -r "[.[] | \
{ \
  ID: .ID,\
  created_at: .created_at,\
  updated_at: .updated_at,\
  condition_qty: .condition.qty,\
  condition_item: .condition.item,\
  discount_type: .discount.type, \
  discount_item: .discount.details.item, \
  discount_value: .discount.details.value \
 }\
]" ./discount_rules.json > discounts.json

 sqlitebiter -o database.db file --format json ./products.json ./discounts.json