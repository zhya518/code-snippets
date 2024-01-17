db.uinfo.find({"vid":{"$gt":0}, "register_time":{$exists:false}},{_id:1}).sort({_id:1}).forEach(
    function(doc) {
        print( doc._id);
    }
);
