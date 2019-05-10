# 192_168_9_230_player 

## t_account_seal_off 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
## t_player_game 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
## t_player_game_total 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
## t_player_id 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
## t_player_packsack 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
## t_player_welffare 

COLUMN_NAME | COLUMN_TYPE | COLUMN_KEY | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_COMMENT | EXTRA
---|---|---|---|---|---|---|
id | bigint(20) | PRI |  | NO | 数据递增ID | auto_increment | 
playerID | bigint(20) | MUL |  | NO | 玩家ID |  | 
welffareID | bigint(20) |  |  | NO | 任务ID |  | 
finishTimes | int(11) |  |  | YES | 完成次数 |  | 
drawTimes | int(11) |  |  | YES | 领取次数 |  | 
canDrawTimes | int(11) |  |  | YES | 可领取次数 |  | 
createTime | datetime |  |  | YES | 创建时间 |  | 
createBy | varchar(64) |  |  | YES | 创建人 |  | 
updateTime | datetime |  |  | YES | 更新时间 |  | 
updateBy | varchar(64) |  |  | YES | 更新人 |  | 


