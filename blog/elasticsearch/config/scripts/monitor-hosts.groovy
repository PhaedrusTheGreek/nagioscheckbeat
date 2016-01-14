def minutes = 5
def now = DateTime.now().getMillis()
ctx.vars.hosts = [ up: [], down: [] ]
ctx.payload.aggregations.hosts.buckets.each {
 def last_heartbeat = it.latest_heartbeat.hits.hits[0].sort[0];
 def ms_ago = now - last_heartbeat
 if (ms_ago > (minutes * 1000) ){
   ctx.vars.hosts.down.add( [ hostname: it.key, last_heartbeat: new Date(last_heartbeat) ])
 } 
}
return ctx.vars.hosts.down.size() > 0
