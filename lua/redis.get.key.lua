local cursor="0";
local count = 0;
repeat
-- redis-cli -h 10.26.226.144 -p 4022 --eval delete-keys.lua
-- redis-cli -h 10.26.226.144 -p 4022 --raw keys "key:uinfo:*" | xargs redis-cli -h 10.70.28.27 -p 4011 --raw GET
 local pattern = "key:uinfo:*"
 local scanResult = redis.call("SCAN", cursor, "MATCH", pattern, "COUNT", 100);
	local keys = scanResult[2];
	for i = 1, #keys do
		local key = keys[i];
		redis.replicate_commands()
		local uinfo = redis.call("GET", key);
		count = count +1;
	end;
	cursor = scanResult[1];
until cursor == "0";
return "Total "..count.." keys Deleted" ;
