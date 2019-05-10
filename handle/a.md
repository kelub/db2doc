# t_account_seal_off
COLUMN_COMMENT | COLUMN_DEFAULT | COLUMN_KEY | COLUMN_NAME | COLUMN_TYPE | IS_NULLABLE
---|---|---|---|---|---|
数据递增ID |  | PRI | id | bigint(20) | NO | 
玩家ID |  | UNI | playerID | bigint(20) | NO | 
封停开始时间 |  |  | startTime | datetime | YES | 
封停截止时间 |  |  | endTime | datetime | YES | 
封停账号事由 |  |  | reason | varchar(256) | YES | 
封停天数 | 0 |  | days | int(11) | NO | 


