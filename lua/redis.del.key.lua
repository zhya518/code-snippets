local cursor="0";
local count = 0;
repeat
-- redis-cli -h 10.70.28.27 -p 4011 --eval delete-keys.lua
-- redis-cli -h 10.70.28.27 -p 4011 --raw keys "userid:*" | xargs redis-cli -h 10.70.28.27 -p 4011 --raw DEL
 local pattern = "userid:*"
 local scanResult = redis.call("SCAN", cursor, "MATCH", pattern, "COUNT", 100);
	local keys = scanResult[2];
	for i = 1, #keys do
		local key = keys[i];
		redis.replicate_commands()
		redis.call("DEL", key);
		count = count +1;
	end;
	cursor = scanResult[1];
until cursor == "0";
return "Total "..count.." keys Deleted" ;
