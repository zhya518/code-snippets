local rcall=redis.call

local function clean_global_rank()
    local key = ARGV[2]
    local start = ARGV[3]
    local stop = ARGV[4]
    -- no withscores
    local remarks = rcall("zrange", key, start, stop)
    local remark_key = key.."_remark"
    for i=1, #remarks do
        local delkey = remarks[i]
        rcall("hdel", remark_key, delkey)
    end
    rcall("zremrangebyrank", key, start, stop)
    return 1
end

if #ARGV==0 then
    local err={}
    err['err']='param num must large than 0'
    return err
end

if ARGV[1] == 'clean_global_rank' then
    return clean_global_rank()
else
    local err={}
    err['err']='unknown function'..func_name
    return err
end

