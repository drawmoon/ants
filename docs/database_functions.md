# Database Functions

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- [Statement](#statement)
  - [Page](#page)
- [Function](#function)
  - [Avg](#avg)
  - [Count](#count)
  - [Sum](#sum)
  - [Max](#max)
  - [Min](#min)
  - [Median](#median)
  - [Percentile](#percentile)
  - [Round](#round)
  - [Convert](#convert)
  - [ToChar](#tochar)
  - [Rpad](#rpad)
  - [Now](#now)
  - [Date](#date)
  - [DateAdd](#dateadd)
  - [DateSub](#datesub)
  - [DateTrunc](#datetrunc)
  - [DateFormat](#dateformat)
  - [Extract](#extract)
  - [RawSql](#rawsql)
  - [User](#user)
- [Expression](#expression)
  - [Column](#column)
  - [Comparison](#comparison)
  - [Logical](#logical)
  - [Not](#not)
  - [In](#in)
  - [IsNull](#isnull)
  - [Like](#like)
  - [Between](#between)
  - [CaseWhen](#casewhen)
  - [Empty](#empty)
- [Value](#value)
  - [String](#string)
  - [Integer](#integer)
  - [Numeric](#numeric)
  - [Interval](#interval]
  - [DateTime](#datetime)
  - [BindVariable](#bindvariable)

## Statement

### Page

>

- Syntax

```sql
PAGE(statement, skip, limit)
```

- Arguments

| Name | Description |
| --- | --- |
| statement | sql statement |
| skip | skip number |
| limit | limit number |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql

```

- Output

```sql
# PostgreSQL
from t limit 0 offset 5
# Version major > 9
from t offset 0 rows fetch next 5 rows only

# MySQL
from t limit 0,5

# Sql Server
from t order by id offset 0 rows fetch next 5 rows only

# Oracle
select "_t1".* from (select ROWNUM as "_", "_t".* from ((from t) "_t")) "_t1" where rownum >= 0 and rownum <= 5
# Version major > 10
from t offset 0 rows fetch next 5 rows only
```

## Function

### Avg

>

- Syntax

```sql
AVG([ALL | DISTINCT] expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be calculated. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
AVG(abc)
AVG(ALL abc)      # all
AVG(DISTINCT abc) # distinct
```

- Output

```sql
avg(abc)
avg(all abc)      # all
avg(distinct abc) # distinct
```

### Count

>

- Syntax

```sql
COUNT([ALL | DISTINCT] expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be counted. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
COUNT(abc)
COUNT(ALL abc)      # all
COUNT(DISTINCT abc) # distinct
```

- Output

```sql
count(abc)
count(all abc)      # all
count(distinct abc) # distinct
```

### Sum

>

- Syntax

```sql
SUM([ALL | DISTINCT] expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be summed. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
SUM(abc)
SUM(ALL abc)      # all
SUM(DISTINCT abc) # distinct
```

- Output

```sql
sum(abc)
sum(all abc)      # all
sum(distinct abc) # distinct
```

### Max

>

- Syntax

```sql
MAX([ALL | DISTINCT] expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be evaluated. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
MAX(abc)
MAX(ALL abc)      # all
MAX(DISTINCT abc) # distinct
```

- Output

```sql
max(abc)
max(all abc)      # all
max(distinct abc) # distinct
```

### Min

>

- Syntax

```sql
MIN([ALL | DISTINCT] expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be evaluated. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
MIN(abc)
MIN(ALL abc)      # all
MIN(DISTINCT abc) # distinct
```

- Output

```sql
min(abc)
min(all abc)      # all
min(distinct abc) # distinct
```

### Median

>

- Syntax

```sql
MEDIAN(expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be calculated |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | × |
| Sql Server | × |
| Oracle | × |

- Example

```sql
MEDIAN(abc)
```

- Output

```sql
# Postgres
percentile_cont(0.5) within group (order by abc)
```

### Percentile

>

- Syntax

```sql
PERCENTILE[_DISC | CONT](expr, num, [ASC | DESC])
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be calculated |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | × |
| Sql Server | × |
| Oracle | × |

- Example

```sql
PERCENTILE(abc, 0.25)
PERCENTILE(abc, 0.25, ASC)  # ASC
PERCENTILE(abc, 0.25, DESC) # DESC
PERCENTILE_DIST(abc, 0.25)  # DIST
```

- Output

```sql
# Postgres
percentile_cont(0.25) within group (order by abc)
percentile_cont(0.25) within group (order by abc asc)  # ASC
percentile_cont(0.25) within group (order by abc desc) # DESC
percentile_dist(0.25) within group (order by abc)      # DIST

# Oracle
PERCENTILE_CONT(0.25) WITHIN GROUP (ORDER BY abc)
PERCENTILE_CONT(0.25) WITHIN GROUP (ORDER BY abc ASC)  # ASC
PERCENTILE_CONT(0.25) WITHIN GROUP (ORDER BY abc DESC) # DESC
PERCENTILE_DIST(0.25) WITHIN GROUP (ORDER BY abc)      # DIST
```

### Round

>

- Syntax

```sql
ROUND(expr, [integer])
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be calculated |
| integer | The number of decimal places to be rounded to |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
ROUND(abc)
```

- Output

```sql
round(abc)
```

### Convert

>

- Syntax

```sql
expr::type
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be converted |
| type | The type to be converted to |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
NOW()::text
```

- Output

```sql
# PostgreSQL
now()::varchar

# MySQL
convert(now(), 'char')

# Sql Server
convert(varchar, getdate(), 121)

# Oracle
TO_CHAR(SYSTIMESTAMP)
```

### ToChar

>

- Syntax

```sql
TO_CHAR(expr, [fmt])
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The string to be converted. |
| fmt | The format to use for the output. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
TO_CHAR(NOW(), 'yyyy-MM-dd')
```

- Output

```sql
# PostgreSQL
now()::varchar

# MySQL
convert(now(), 'char')

# Sql Server
convert(varchar, getdate(), 121)

# Oracle
TO_CHAR(SYSTIMESTAMP)
```

### Rpad

>

- Syntax

```sql
RPAD(expr, integer, text)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The string to be padded. |
| integer | The length of the output string. |
| text | The text to be used for padding. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | × |
| MySQL | × |
| Sql Server | × |
| Oracle | × |

- Example

```sql
RPAD('2022', 8, '01')
```

- Output

```sql
# PostgreSQL
rpad('2022', 8, '01')

# MySQL
rpad('2022', 8, '01')

# Sql Server
left('2022' + replicate('01', 4), 8)
```

### Now

>

- Syntax

```sql
NOW()
```

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
NOW()
```

- Output

```sql
# PostgreSQL
now()

# MySQL
now()

# Sql Server
getdate()

# Oracle
SYSTIMESTAMP
```

### Date

>

- Syntax

```sql
DATE(expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be converted to a date. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
DATE(NOW())
```

- Output

```sql
# PostgreSQL
date(now())

# MySQL
date(now())

# Sql Server
convert(date, getdate())

# Oracle
TRUNC(SYSTIMESTAMP, 'DAY')
```

### DateAdd

>

- Syntax

```sql
<YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND>_ADD(interval, integer)
```

- Arguments

| Name | Description |
| --- | --- |
| interval | The interval to add. |
| integer | The number of intervals to add. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
YEAR_ADD(NOW(), 1)    # Year
QUARTER_ADD(NOW(), 1) # Quarter
MONTH_ADD(NOW(), 1)   # Month
WEEK_ADD(NOW(), 1)    # Week
DAY_ADD(NOW(), 1)     # Day
HOUR_ADD(NOW(), 1)    # Hour
MINUTE_ADD(NOW(), 1)  # Minute
SECOND_ADD(NOW(), 1)  # Second
```

- Output

```sql
# PostgreSQL
now() + interval '1 year'    # Year
now() + interval '1 quarter' # Quarter
now() + interval '1 month'   # Month
now() + interval '1 week'    # Week
now() + interval '1 day'     # Day
now() + interval '1 hour'    # Hour
now() + interval '1 minute'  # Minute
now() + interval '1 second'  # Second

# MySQL
date_add(now(), interval 1 year)    # Year
date_add(now(), interval 1 quarter) # Quarter
date_add(now(), interval 1 month)   # Month
date_add(now(), interval 1 week)    # Week
date_add(now(), interval 1 day)     # Day
date_add(now(), interval 1 hour)    # Hour
date_add(now(), interval 1 minute)  # Minute
date_add(now(), interval 1 second)  # Second

# Oracle
SYSTIMESTAMP + INTERVAL '1' YEAR   # Year
SYSTIMESTAMP + INTERVAL '3' MONTH  # Quarter
SYSTIMESTAMP + INTERVAL '1' MONTH  # Month
SYSTIMESTAMP + INTERVAL '7' DAY    # Week
SYSTIMESTAMP + INTERVAL '1' DAY    # Day
SYSTIMESTAMP + INTERVAL '1' HOUR   # Hour
SYSTIMESTAMP + INTERVAL '1' MINUTE # Minute
SYSTIMESTAMP + INTERVAL '1' SECOND # Second
```

### DateSub

>

- Syntax

```sql
<YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND>_SUB(interval, integer)
```

- Arguments

| Name | Description |
| --- | --- |
| interval | The interval to subtract from the current date. |
| integer | The number of intervals to subtract. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
YEAR_SUB(NOW(), 1)    # Year
QUARTER_SUB(NOW(), 1) # Quarter
MONTH_SUB(NOW(), 1)   # Month
WEEK_SUB(NOW(), 1)    # Week
DAY_SUB(NOW(), 1)     # Day
HOUR_SUB(NOW(), 1)    # Hour
MINUTE_SUB(NOW(), 1)  # Minute
SECOND_SUB(NOW(), 1)  # Second
```

- Output

```sql
# PostgreSQL
now() - interval '1 year'    # Year
now() - interval '1 quarter' # Quarter
now() - interval '1 month'   # Month
now() - interval '1 week'    # Week
now() - interval '1 day'     # Day
now() - interval '1 hour'    # Hour
now() - interval '1 minute'  # Minute
now() - interval '1 second'  # Second

# MySQL
date_sub(now(), interval 1 year)    # Year
date_sub(now(), interval 1 quarter) # Quarter
date_sub(now(), interval 1 month)   # Month
date_sub(now(), interval 1 week)    # Week
date_sub(now(), interval 1 day)     # Day
date_sub(now(), interval 1 hour)    # Hour
date_sub(now(), interval 1 minute)  # Minute
date_sub(now(), interval 1 second)  # Second

# Oracle
SYSTIMESTAMP - INTERVAL '1' YEAR   # Year
SYSTIMESTAMP - INTERVAL '3' MONTH  # Quarter
SYSTIMESTAMP - INTERVAL '1' MONTH  # Month
SYSTIMESTAMP - INTERVAL '7' DAY    # Week
SYSTIMESTAMP - INTERVAL '1' DAY    # Day
SYSTIMESTAMP - INTERVAL '1' HOUR   # Hour
SYSTIMESTAMP - INTERVAL '1' MINUTE # Minute
SYSTIMESTAMP - INTERVAL '1' SECOND # Second
```

### DateTrunc

>

- Syntax

```sql
DATE_TRUNC(interval, fmt)
```

- Arguments

| Name | Description |
| --- | --- |
| interval | The interval to truncate to. |
| fmt | The format to use for the output. |

- Format

| Name | Description |
| --- | --- |
| Year | 从年份开始截断 |
| Quarter | 当前季度的第一天 |
| Month | 从月份开始截断 |
| Week | 当前周的第一天 |
| Day | 从日期开始截断 |
| Hour | 从小时开始截断 |
| Minute | 从分钟开始截断 |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | × |

- Example

```sql
DATE_TRUNC(NOW(), Year);    # Year
DATE_TRUNC(NOW(), Quarter); # Quarter
DATE_TRUNC(NOW(), Month);   # Month
DATE_TRUNC(NOW(), Week);    # Week
DATE_TRUNC(NOW(), Day);     # Day
DATE_TRUNC(NOW(), Hour);    # Hour
DATE_TRUNC(NOW(), Minute);  # Minute
```

- Output

```sql
# PostgreSQL
date_trunc('year', now());    # Year
date_trunc('quarter', now()); # Quarter
date_trunc('month', now());   # Month
date_trunc('week', now());    # Week
date_trunc('day', now());     # Day
date_trunc('hour', now());    # Hour
date_trunc('minute', now());  # Minute

# MySQL
timestamp(date_format(now(), '%Y-01-01'))                                 # Year
timestamp(concat(year(now()), '-', ceil(month(now()) / 3) * 3 - 2, '-1')) # Quarter
timestamp(date_format(now(), '%Y-%m-01'))                                 # Month
timestamp(date(date_sub(now(), interval weekday(now()) day)))             # Week
timestamp(date_format(now(), '%Y-%m-%d'))                                 # Day
timestamp(date_format(now(), '%Y-%m-%d %H:00:00'))                        # Hour
timestamp(date_format(now(), '%Y-%m-%d %H:%i:00'))                        # Minute

# Sql Server
convert(datetime, format(getdate(), 'yyyy-01-01'))                                                           # Year
convert(datetime, concat(year(getdate()), '-', (ceiling(datepart(month, getdate()) / 3) + 1) * 3 - 2, '-1')) # Quarter
convert(datetime, format(getdate(), 'yyyy-MM-01'))                                                           # Month
convert(datetime, format(dateadd(day, -(datepart(weekday, getdate()) - 2), getdate()), 'yyyy-MM-dd'))        # Week
convert(datetime, format(getdate(), 'yyyy-MM-dd'))                                                           # Day
convert(datetime, format(getdate(), 'yyyy-MM-dd hh:00:00'))                                                  # Hour
convert(datetime, format(getdate(), 'yyyy-MM-dd hh:mi:00'))                                                  # Minute

# Oracle
TRUNC(SYSTIMESTAMP, 'YEAR')  # Year
TRUNC(SYSTIMESTAMP, 'Q')     # Quarter
TRUNC(SYSTIMESTAMP, 'MONTH') # Month
TRUNC(SYSTIMESTAMP, 'IW')    # Week
TRUNC(SYSTIMESTAMP, 'DD')    # Day
TRUNC(SYSTIMESTAMP, 'HH')    # Hour
TRUNC(SYSTIMESTAMP, 'MI')    # Minute
```

### DateFormat

>

- Syntax

```sql

```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | × |
| MySQL | × |
| Sql Server | × |
| Oracle | × |

- Example

```sql

```

- Output

```sql

```

### Extract

>

- Syntax

```sql
<YEAR | QUARTER | MONTH | WEEK | WEEKDAY | DAY | HOUR | MINUTE | SECOND>(expr)
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to extract from. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | × |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | × |

- Example

```sql
YEAR(NOW())    # Year
QUARTER(NOW()) # Quarter
MONTH(NOW())   # Month
WEEK(NOW())    # Week
WEEKDAY(NOW()) # WeekDay
DAY(NOW())     # Day
HOUR(NOW())    # Hour
MINUTE(NOW())  # Minute
SECOND(NOW())  # Second
```

- Output

```sql
# PostgreSQL
extract(year from now())    # Year
extract(quarter from now()) # Quarter
extract(month from now())   # Month
extract(week from now())    # Week
extract(dow from now())     # WeekDay
extract(day from now())     # Day
extract(hour from now())    # Hour
extract(minute from now())  # Minute
extract(second from now())  # Second

# MySQL
year(now())    # Year
quarter(now()) # Quarter
month(now())   # Month
week(now())    # Week
weekday(now()) # WeekDay
day(now())     # Day
hour(now())    # Hour
minute(now())  # Minute
second(now())  # Second

# Sql Server
datepart(year, getdate())     # Year
datepart(quarter, getdate())  # Quarter
datepart(month, getdate())    # Month
datepart(week, getdate())     # Week
datepart(weekday, getdate())  # WeekDay
datepart(day, getdate())      # Day
datepart(hour, getdate())     # Hour
datepart(minute, getdate())   # Minute
datepart(second, getdate())   # Second

# Oracle
extract(year from systimestamp)        # Year
to_number(to_char(systimestamp, 'q'))  # Quarter
extract(month from systimestamp)       # Month
to_number(to_char(systimestamp, 'ww')) # Week
to_number(to_char(systimestamp, 'd'))  # WeekDay
extract(day from systimestamp)         # Day
extract(hour from systimestamp)        # Hour
extract(minute from systimestamp)      # Minute
extract(second from systimestamp)      # Second
```

### RawSql

> RawSql 直通函数可用于將 SQL 运算表达式直接传送到数据库，而不由 U# 进行解析。

- Syntax

```sql
RAW_SQL('statement')
```

- Arguments

| Name | Description |
| --- | --- |
| statement | sql statement |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
RAW_SQL('date(now())')
```

- Output

```sql
date(now())
```

### User

>

- Syntax

```sql

```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql

```

- Output

```sql

```

## Expression

### Column

>

- Syntax

```sql
expr
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc
```

- Output

```sql
abc
```

### Comparison

>

- Syntax

```sql
expr <= | <> | < | > | <= | >= > expr
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to compare |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc = 1
```

- Output

```sql
abc = 1
```

### Logical

>

- Syntax

```sql
expr <AND | OR> expr
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to compare |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc = 1 AND bcd = 2
```

- Output

```sql
abc = 1 and bcd = 2
```

### Not

>

- Syntax

```sql
NOT expr
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
NOT abc = 1
```

- Output

```sql
not abc = 1
```

### In

>

- Syntax

```sql
column_name [NOT] IN (value1, value2, ...)
```

- Arguments

| Name | Description |
| --- | --- |
| column_name | The column name |
| value1, value2,... | The value to compare |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc IN (1, 2, 3)
```

- Output

```sql
abc in (1, 2, 3)
```

### IsNull

>

- Syntax

```sql
expr IS [NOT] NULL
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The value to compare |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc IS NULL
```

- Output

```sql
abc is null
```

### Like

>

- Syntax

```sql
expr [NOT] LIKE pattern
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The value to compare |
| pattern | The pattern to compare |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc LIKE 'abc%'
```

- Output

```sql
abc like 'abc%'
```

### Between

>

- Syntax

```sql
expr [NOT] BETWEEB low and higt
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be compared |
| low | The lower bound |
| higt | The upper bound |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
abc BETWEEB 3 and 10
```

- Output

```sql
abc between 3 and 10
```

### CaseWhen

>

- Syntax

```sql
CASE WHEN expr THEN result1 [WHEN expr THEN result2]...[ELSE resultN] END
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be evaluated. |
| result1 | The result to be returned if the expression is true. |
| result2 | The result to be returned if the expression is true and the previous result is false. |
|... |... |
| resultN | The result to be returned if none of the previous results are true. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
CASE abc WHEN 'low' THEN 1 WHEN 'high' THEN 2 ELSE 0 END
```

- Output

```sql
case abc when 'low' then 1 when 'high' then 2 else 0 end
```

### Empty

>

- Syntax

```sql
:EMPTY
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc = 1 AND :EMPTY
```

- Output

```sql
abc = 1
```

## Value

### String

>

- Syntax

```sql
'literal'
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
'hello world'
```

- Output

```sql
'hello world'
```

### Integer

>

- Syntax

```sql
whole number
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
1000
```

- Output

```sql
1000
```

### Numeric

>

- Syntax

```sql
float number
```

- Arguments

| Name | Description |
| --- | --- |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
1.001
```

- Output

```sql
1.001
```

### Interval

>

- Syntax

```sql
INTERVAL expr YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The expression to be evaluated. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
INTERVAL '1' YEAR
INTERVAL '1' QUARTER
INTERVAL '1' MONTH
INTERVAL '1' WEEK
INTERVAL '1' DAY
INTERVAL '1' HOUR
INTERVAL '1' MINUTE
INTERVAL '1' SECOND
```

- Output

```sql
# PostgreSQL
interval '1 year'    # Year
interval '1 quarter' # Quarter
interval '1 month'   # Month
interval '1 week'    # Week
interval '1 day'     # Day
interval '1 hour'    # Hour
interval '1 minute'  # Minute
interval '1 second'  # Second

# MySQL
interval 1 year    # Year
interval 1 quarter # Quarter
interval 1 month   # Month
interval 1 week    # Week
interval 1 day     # Day
interval 1 hour    # Hour
interval 1 minute  # Minute
interval 1 second  # Second

# Oracle
INTERVAL '1' YEAR   # Year
INTERVAL '3' MONTH  # Quarter
INTERVAL '1' MONTH  # Month
INTERVAL '7' DAY    # Week
INTERVAL '1' DAY    # Day
INTERVAL '1' HOUR   # Hour
INTERVAL '1' MINUTE # Minute
INTERVAL '1' SECOND # Second
```

### DateTime

>

- Syntax

```sql
#datetime#
```

- Arguments

| Name | Description |
| --- | --- |
| datetime | The datetime value. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | ✓ |
| Oracle | ✓ |

- Example

```sql
#2023-04-23T16:00:00.000Z#
```

- Output

```sql
2023-04-24 00:00:00.000
```

### BindVariable

>

- Syntax

```sql
:expr
```

- Arguments

| Name | Description |
| --- | --- |
| expr | The bind variable name. |

- Support

| Database | Support |
| --- | --- |
| PostgreSQL | ✓ |
| MySQL | ✓ |
| Sql Server | × |
| Oracle | ✓ |

- Example

```sql
abc = :abc // when abc maps a value of 123
```

- Output

```sql
abc = 123
```
