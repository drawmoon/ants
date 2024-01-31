# ants

Model-oriented SQL generator for Go.

Distribution of users by city:

```json
{
    "items": [
        {
            "id": 1,
            "alias": "userCount",
            "type": 0,
            "arithmetic": 2
        },
        {
            "id": 2,
            "alias": "city",
            "type": 1
        }
    ]
}
```

Output SQL:

```sql
SELECT
    COUNT(t1.user_id) AS userCount,
    t2.city
FROM
    users t1
LEFT JOIN
    cities t2
ON
    t1.city_id = t2.id
GROUP BY
    t2.city;
```
