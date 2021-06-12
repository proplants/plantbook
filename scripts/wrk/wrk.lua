request = function()
   local url = "/health/ready"
   return wrk.format("GET", url)
end